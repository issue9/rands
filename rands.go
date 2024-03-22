// SPDX-FileCopyrightText: 2016-2024 caixw
//
// SPDX-License-Identifier: MIT

// Package rands 生成各种随机字符串
//
//	// 生成一个长度介于 [6,9) 之间的随机字符串
//	str := rands.String(6, 9, "1343567")
//
//	// 生成一个带缓存功能的随机字符串生成器
//	r := rands.New(nil, 100, 5, 7, "adbcdefgadf;dfe1334")
//	ctx,cancel := context.WithCancel(context.Background())
//	go r.Serve(ctx)
//	defer cancel()
//	str1 := r.String()
//	str2 := r.String()
//
// NOTE: 仅是随机字符串，不保证唯一性。
package rands

import (
	"context"
	"math/rand/v2"
)

// Bytes 产生随机字符数组
//
// 其长度为[min, max)，bs 所有的随机字符串从此处取。
func Bytes(min, max int, bs []byte) []byte {
	checkArgs(min, max, bs)
	return bytes(rand.IntN, rand.Uint64, min, max, bs)
}

// String 产生一个随机字符串
//
// 其长度为[min, max)，bs 可用的随机字符。
func String(min, max int, bs []byte) string { return string(Bytes(min, max, bs)) }

// Rands 提供随机字符串的生成
type Rands struct {
	intn func(int) int
	u64  func() uint64

	min, max int
	bytes    []byte
	channel  chan []byte
}

// New 声明 [Rands]
//
// 如果 r 为 nil，将采用默认的随机函数；
// bufferSize 缓存的随机字符串数量，若为 0,表示不缓存；
func New(r *rand.Rand, bufferSize, min, max int, bs []byte) *Rands {
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

	return &Rands{
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
func (r *Rands) Bytes() []byte { return <-r.channel }

// String 产生一个随机字符串
//
// 功能与全局函数 [String] 相同，但参数通过 [New] 预先指定。
func (r *Rands) String() string { return string(<-r.channel) }

func (r *Rands) Serve(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			close(r.channel)
			return ctx.Err()
		case r.channel <- bytes(r.intn, r.u64, r.min, r.max, r.bytes):
		}
	}
}

// 生成介于 [min,max) 长度的随机字符数组
func bytes(intN func(int) int, u64 func() uint64, min, max int, bs []byte) []byte {
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

	ret := make([]byte, l)
	for i := 0; i < l; {
		// 在 index 不够大时，index>>bit 可能会让 index 变为 0，为 0 的 index 应该抛弃。
		for index := u64(); index > 0 && i < l; {
			if idx := index & mask; idx < ll {
				ret[i] = bs[idx]
				i++
			}
			index >>= bit
		}
	}

	return ret
}

// 检测各个参数是否合法
func checkArgs(min, max int, bs []byte) {
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
