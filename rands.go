// SPDX-FileCopyrightText: 2016-2024 caixw
//
// SPDX-License-Identifier: MIT

// Package rands 生成各种随机字符串
//
//	// 生成一个长度介于 [6,9) 之间的随机字符串
//	str := rands.String(6, 9, []byte("1343567"))
//
//	// 生成一个带缓存功能的随机字符串生成器
//	r := rands.New(nil, 100, 5, 7, []byte("adbcdefgadf;dfe1334"))
//	ctx,cancel := context.WithCancel(context.Background())
//	go r.Serve(ctx)
//	defer cancel()
//	str1 := r.String()
//	str2 := r.String()
//
// NOTE: 仅是随机字符串，不保证唯一性。
package rands

import (
	"bytes"
	"context"
	"math/rand/v2"
	"unsafe"
)

// Char 约束随机字符串字符的类型
//
// byte 的性能为好于 rune。
type Char interface{ ~byte | ~rune }

// Bytes 从 bs 中随机抓取 [min,max) 个字符并返回
//
// NOTE: bs 的类型可以是 rune，但是返回类型始终是 []byte，所以用 len 判断返回值可能其值会很大。
func Bytes[T Char](min, max int, bs []T) []byte {
	checkArgs(min, max, bs)
	return gen(rand.IntN, rand.Uint64, min, max, bs)
}

// String 产生一个随机字符串
//
// 其长度为[min, max)，bs 可用的随机字符。
func String[T Char](min, max int, bs []T) string {
	cs := Bytes(min, max, bs)
	return unsafe.String(unsafe.SliceData(cs), len(cs))
}

// Rands 提供随机字符串的生成
type Rands[T Char] struct {
	intn func(int) int
	u64  func() uint64

	min, max int
	bytes    []T
	channel  chan []byte
}

// New 声明 [Rands]
//
// 如果 r 为 nil，将采用默认的随机函数；
// bufferSize 缓存的随机字符串数量，若为 0,表示不缓存；
func New[T Char](r *rand.Rand, bufferSize, min, max int, bs []T) *Rands[T] {
	checkArgs(min, max, bs)

	if bufferSize <= 0 {
		panic("bufferSize 必须大于零")
	}

	var intn func(int) int
	var u64 func() uint64
	if r == nil {
		intn = rand.IntN
		u64 = rand.Uint64
	} else {
		intn = r.IntN
		u64 = rand.Uint64
	}

	return &Rands[T]{
		intn: intn,
		u64:  u64,

		min:     min,
		max:     max,
		bytes:   bs,
		channel: make(chan []byte, bufferSize),
	}
}

// Bytes 产生随机字符数组
//
// 功能与全局函数 [Bytes] 相同，但参数通过 [New] 预先指定。
func (r *Rands[T]) Bytes() []byte { return <-r.channel }

// String 产生一个随机字符串
//
// 功能与全局函数 [String] 相同，但参数通过 [New] 预先指定。
func (r *Rands[T]) String() string {
	cs := r.Bytes()
	return unsafe.String(unsafe.SliceData(cs), len(cs))
}

func (r *Rands[T]) Serve(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			close(r.channel)
			return ctx.Err()
		case r.channel <- gen(r.intn, r.u64, r.min, r.max, r.bytes):
		}
	}
}

// 生成介于 [min,max) 长度的随机字符数组
//
// intN 用于生成一个介于 [min, max) 之间的数据；
// u64 用生成一个 uint64 的随机数；
func gen[T Char](intN func(int) int, u64 func() uint64, min, max int, bs []T) []byte {
	var l int
	if max-1 == min {
		l = min
	} else {
		l = min + intN(max-min)
	}

	ll := uint64(len(bs))

	// 将一个 uint64 的随机数据，按 bs 的长度拆分为多个数组下标。
	// 比如 bs 的长度为 10，那么一个 uint64 的随机数，
	// 理论上可以表示 64 / 10 = 6 个数组下标从 bs 中拿值。

	var bit int     // 表示 len(bs) 下标的最大数值需要的位数
	var mask uint64 // bit 的掩码
	for bit = 0; bit < 64; bit++ {
		if mask = 1<<bit - 1; mask >= ll {
			break
		}
	}

	if _, isRune := any(bs).([]rune); !isRune {
		ret := make([]byte, l)
		for i := 0; i < l; { // 循环内会增加 i 的值。比直接使用 for i:= range l 要快。
			// 在 index 不够大时，index>>bit 可能会让 index 提早变为 0，为 0 的 index 应该抛弃。
			for index := u64(); index > 0 && i < l; {
				if idx := index & mask; idx < ll {
					ret[i] = byte(bs[idx])
					i++
				}
				index >>= bit
			}
		}
		return ret
	} else {
		s := bytes.Buffer{}
		s.Grow(2 * l)

		for i := 0; i < l; { // 循环内会增加 i 的值。比直接使用 for i:= range l 要快。
			// 在 index 不够大时，index>>bit 可能会让 index 提早变为 0，为 0 的 index 应该抛弃。
			for index := u64(); index > 0 && i < l; {
				if idx := index & mask; idx < ll {
					s.WriteRune(rune(bs[idx]))
					i++
				}
				index >>= bit
			}
		}

		return s.Bytes()
	}
}

// 检测各个参数是否合法
func checkArgs[T Char](min, max int, bs []T) {
	if min <= 0 {
		panic("min 值必须大于 0")
	}
	if max <= min {
		panic("max 必须大于 min")
	}

	if len(bs) == 0 {
		panic("无效的 bs 参数")
	}
}
