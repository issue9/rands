// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// 生成各种随机字符串的包
//
//  // 生成一个长度介于[6,9)之间由小写字母与数字组成的的随机字符串
//  str := rands.String(6, 9, rands.Lower, rands.Digit)
//
//  // 生成一个带缓存功能的随机字符串生成器
//  r := rands.New(time.Now().Unix(), 100, 5, 7, rands.Lower, rands.Digit, rands.Punct)
//  str1 := r.String()
//  str2 := r.String()
package rands

import (
	"errors"
	"math/rand"
	"time"
)

const (
	Upper = iota // 大写字母
	Lower        // 小写字母
	Digit        // 数字
	Punct        // 标点符号
	size
)

// 供全局函数使用的随机函数生成。
// Bytes和String依赖此项。
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// 随机字符串取值的表
var table = [][]byte{
	Upper: []byte("ABCDEFGHIJKLMNPQRSTUVWXYZ"),
	Lower: []byte("abcdefghijkmnpqrstuvwxyz"),
	Digit: []byte("0123456789"),
	Punct: []byte("~!@#$%^&*()_+-={}|[]\\:\";'<>,.?/"),
}

// 手动指定一个随机种子，默认情况下使用当前包初始化时的时间戳作为随机种子。
// Bytes和String依赖此项。但是Rands有专门的随机函数，不受此影响。
func Seed(seed int64) {
	random = rand.New(rand.NewSource(seed))
}

// 产生随机字符数组，其长度为[min, max)。
// cats 随机字符串的类型，指定非法值是触发panic
func Bytes(min, max int, cats ...int) []byte {
	if err := checkArgs(min, max, cats); err != nil {
		panic(err)
	}

	return bytes(random, min+random.Intn(max-min), cats)
}

// 产生一个随机字符串，其长度为[min, max)。
// cats 随机字符串的类型，指定非法值是触发panic
func String(min, max int, cats ...int) string {
	return string(Bytes(min, max, cats...))
}

// Rands 提供了对参数的简单包装，方便用户批量产生相同的类型的随机字符串。
type Rands struct {
	random    *rand.Rand
	min, max  int
	cats      []int
	hasBuffer bool
	channel   chan []byte
}

// 声明一个Rands变量。
// seed 随机种子，若为0表示使用当前时间作为随机种子。
// bufferSize 缓存的随机字符串数量，若为0,表示不缓存。
func New(seed int64, bufferSize, min, max int, cats ...int) (*Rands, error) {
	if err := checkArgs(min, max, cats); err != nil {
		return nil, err
	}
	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	ret := &Rands{
		random:    rand.New(rand.NewSource(seed)),
		min:       min,
		max:       max,
		cats:      cats,
		hasBuffer: bufferSize > 0,
	}
	if ret.hasBuffer {
		ret.channel = make(chan []byte, bufferSize)
		go func() {
			for {
				ret.channel <- bytes(ret.random, min+ret.random.Intn(max-min), cats)
			}
		}()
	}
	return ret, nil
}

// 重新指定随机种子。
func (r *Rands) Seed(seed int64) {
	r.random.Seed(seed)
}

// 产生随机字符数组，功能与全局函数Bytes()相同，但参数通过New()预先指定。
func (r *Rands) Bytes() []byte {
	if r.hasBuffer {
		return <-r.channel
	}

	return bytes(r.random, r.min+r.random.Intn(r.max-r.min), r.cats)
}

// 产生一个随机字符串，功能与全局函数String()相同，但参数通过New()预先指定。
func (r *Rands) String() string {
	if r.hasBuffer {
		return string(<-r.channel)
	}

	return string(r.Bytes())
}

// 生成指定指定长度的随机字符数组
func bytes(r *rand.Rand, l int, cats []int) []byte {
	bs := make([]byte, 0, l)
	i := 0
	for {
		for _, cat := range cats {
			s := table[cat]
			index := r.Intn(len(s))
			bs = append(bs, s[index])

			i++
			if i >= l {
				return bs
			}
		} // end for cats
	} // end for true
}

// 检测各个参数是否合法
func checkArgs(min, max int, cats []int) error {
	if min <= 0 {
		return errors.New("rands.checkArgs:min值必须大于0")
	}
	if max <= min {
		return errors.New("rands.checkArgs:max必须大于min")
	}

	if len(cats) == 0 {
		return errors.New("rands.checkArgs:无效的cat参数")
	}

	for _, cat := range cats {
		if cat < 0 || cat >= size {
			return errors.New("rands.Bytes:无效的cat参数")
		}
	}

	return nil
}
