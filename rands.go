// SPDX-License-Identifier: MIT

// Package rands 生成各种随机字符串的包，仅是随机字符串，不保证唯一。
//
//  // 生成一个长度介于 [6,9) 之间的随机字符串
//  str := rands.String(6, 9, "1343567")
//
//  // 生成一个带缓存功能的随机字符串生成器
//  r := rands.New(time.Now().Unix(), 100, 5, 7, "adbcdefgadf;dfe1334")
//  str1 := r.String()
//  str2 := r.String()
package rands

import (
	"math/rand"
	"time"
)

const (
	alpha                  = "abcdefghijklmnopqrstuvwzyxABCDEFGHIJKLMNOPQRSTUVWZYX"
	number                 = "1234567890"
	punctuation            = "!@#$%^&*()_+[]{};':\",./<>?"
	alphaNumber            = number + alpha
	alphaNumberPunctuation = alphaNumber + punctuation
)

// 提供几种常见的随机字符组合方式
var (
	Alpha                  = []byte(alpha)
	Number                 = []byte(number)
	Punctuation            = []byte(punctuation)
	AlphaNumber            = []byte(alphaNumber)
	AlphaNumberPunctuation = []byte(alphaNumberPunctuation)
)

// 供全局函数使用的随机函数生成。
// Bytes 和 String 依赖此项。
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// Seed 手动指定一个随机种子，默认情况下使用当前包初始化时的时间戳作为随机种子。
// Bytes 和 String 依赖此项。但是 Rands 有专门的随机函数，不受此影响。
func Seed(seed int64) {
	random = rand.New(rand.NewSource(seed))
}

// Bytes 产生随机字符数组，其长度为[min, max)。
// bs 所有的随机字符串从此处取。
func Bytes(min, max int, bs []byte) []byte {
	checkArgs(min, max, bs)

	return bytes(random, min+random.Intn(max-min), bs)
}

// String 产生一个随机字符串，其长度为[min, max)。
// bs 可用的随机字符串。
func String(min, max int, bs []byte) string {
	return string(Bytes(min, max, bs))
}

// Rands 提供了对参数的简单包装，方便用户批量产生相同的类型的随机字符串。
type Rands struct {
	random    *rand.Rand
	min, max  int
	bytes     []byte
	hasBuffer bool
	channel   chan []byte
	done      chan struct{}
}

// New 声明一个 Rands 变量。
//
// seed 随机种子，若为 0 表示使用当前时间作为随机种子。
// bufferSize 缓存的随机字符串数量，若为 0,表示不缓存。
func New(seed int64, bufferSize, min, max int, bs []byte) *Rands {
	checkArgs(min, max, bs)

	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	ret := &Rands{
		random:    rand.New(rand.NewSource(seed)),
		min:       min,
		max:       max,
		bytes:     bs,
		hasBuffer: bufferSize > 0,
		done:      make(chan struct{}),
	}
	if ret.hasBuffer {
		ret.channel = make(chan []byte, bufferSize)

		go func() {
			for {
				select {
				case <-ret.done:
					close(ret.channel)
					return
				case ret.channel <- bytes(ret.random, min+ret.random.Intn(max-min), bs):
				}
			} // end ofr
		}()
	}
	return ret
}

// Seed 重新指定随机种子。
func (r *Rands) Seed(seed int64) {
	r.random.Seed(seed)
}

// Bytes 产生随机字符数组，功能与全局函数Bytes()相同，但参数通过New()预先指定。
func (r *Rands) Bytes() []byte {
	if r.hasBuffer {
		if ret, ok := <-r.channel; ok {
			return ret
		}
		return nil
	}

	return bytes(r.random, r.min+r.random.Intn(r.max-r.min), r.bytes)
}

// String 产生一个随机字符串，功能与全局函数 String() 相同，但参数通过 New() 预先指定。
func (r *Rands) String() string {
	if r.hasBuffer {
		return string(<-r.channel)
	}

	return string(r.Bytes())
}

// Stop 停止产生随机字符串
func (r *Rands) Stop() {
	r.done <- struct{}{}
}

// 生成指定指定长度的随机字符数组
func bytes(r *rand.Rand, l int, bs []byte) []byte {
	ret := make([]byte, l, l)

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
