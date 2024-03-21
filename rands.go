// SPDX-FileCopyrightText: 2016-2024 caixw
//
// SPDX-License-Identifier: MIT

// Package rands 生成各种随机字符串
//
//	// 生成一个长度介于 [6,9) 之间的随机字符串
//	str := rands.String(6, 9, "1343567")
//
//	// 生成一个带缓存功能的随机字符串生成器
//	r := rands.New(time.Now().Unix(), 100, 5, 7, "adbcdefgadf;dfe1334")
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
	"math/rand"
	"time"
)

var (
	chars = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+[]{};':\",./<>?")

	random = rand.New(rand.NewSource(time.Now().UnixNano())) // 供全局函数使用的随机函数
)

// Alpha 返回所有的字母
func Alpha() []byte { return chars[0:52] }

// Number 返回所有的数字
func Number() []byte { return chars[52:62] }

// Punct 返回所有的标点符号
func Punct() []byte { return chars[62:] }

// AlphaNumber [Alpha] + [Number]
func AlphaNumber() []byte { return chars[0:62] }

// AlphaNumberPunct [Alpha] + [Number] + [Punct]
func AlphaNumberPunct() []byte { return chars }

// Seed 手动指定一个随机种子
//
// 默认情况下使用当前包初始化时的时间戳作为随机种子。
// [Bytes] 和 [String] 依赖此项。但是 [Rands] 有专门的随机函数，不受此影响。
func Seed(seed int64) { random = rand.New(rand.NewSource(seed)) } // TODO(go1.22) 改为 rand/v2

// Bytes 产生随机字符数组
//
// 其长度为[min, max)，bs 所有的随机字符串从此处取。
func Bytes(min, max int, bs []byte) []byte {
	checkArgs(min, max, bs)
	return bytes(random, min, max, bs)
}

// String 产生一个随机字符串
//
// 其长度为[min, max)，bs 可用的随机字符。
func String(min, max int, bs []byte) string { return string(Bytes(min, max, bs)) }

// Rands 提供随机字符串的生成
type Rands struct {
	random   *rand.Rand
	min, max int
	bytes    []byte
	channel  chan []byte
}

// New 声明 [Rands]
//
// seed 随机种子，若为 0 表示使用当前时间作为随机种子。
// bufferSize 缓存的随机字符串数量，若为 0,表示不缓存。
func New(seed int64, bufferSize, min, max int, bs []byte) *Rands {
	checkArgs(min, max, bs)

	if bufferSize <= 0 {
		panic("bufferSize 必须大于零")
	}

	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	return &Rands{
		random:  rand.New(rand.NewSource(seed)),
		min:     min,
		max:     max,
		bytes:   bs,
		channel: make(chan []byte, bufferSize),
	}
}

// Seed 重新指定随机种子
func (r *Rands) Seed(seed int64) { r.random.Seed(seed) }

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
		case r.channel <- bytes(r.random, r.min, r.max, r.bytes):
		}
	}
}

// 生成介于 [min,max) 长度的随机字符数组
func bytes(r *rand.Rand, min, max int, bs []byte) []byte {
	var l int
	if max-1 == min {
		l = min
	} else {
		l = min + r.Intn(max-min)
	}

	ret := make([]byte, l)

	for i := 0; i < l; i++ {
		index := r.Intn(len(bs))
		ret[i] = bs[index]
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
