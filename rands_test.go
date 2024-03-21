// SPDX-FileCopyrightText: 2016-2024 caixw
//
// SPDX-License-Identifier: MIT

package rands

import (
	"context"
	"testing"
	"time"

	"github.com/issue9/assert/v4"
)

func TestChars(t *testing.T) {
	a := assert.New(t, false)

	s := Alpha()
	a.Equal(s[0], 'a').
		Equal(s[len(s)-1], 'Z')

	s = Number()
	a.Equal(s[0], '1').
		Equal(s[len(s)-1], '0')

	s = Punct()
	a.Equal(s[0], '!').
		Equal(s[len(s)-1], '?')

	s = AlphaNumber()
	a.Equal(s[0], 'a').
		Equal(s[len(s)-1], '0')

	s = AlphaNumberPunct()
	a.Equal(s[0], 'a').
		Equal(s[len(s)-1], '?')
}

func TestCheckArgs(t *testing.T) {
	a := assert.New(t, false)

	// min < 0
	a.Panic(func() {
		checkArgs(-1, 1, []byte("12"))
	})

	// max <= min
	a.Panic(func() {
		checkArgs(5, 5, []byte("12"))
	})

	// cats为空
	a.Panic(func() {
		checkArgs(5, 6, []byte(""))
	})

	a.NotPanic(func() { checkArgs(5, 6, []byte("123")) })
}

// bytes
func TestBytes1(t *testing.T) {
	a := assert.New(t, false)

	a.NotEqual(bytes(random, 10, 11, []byte("1234123lks;df")), bytes(random, 10, 11, []byte("1234123lks;df")))
	a.NotEqual(bytes(random, 10, 11, []byte("1234123lks;df")), bytes(random, 10, 11, []byte("1234123lks;df")))
	a.NotEqual(bytes(random, 10, 11, []byte("1234123lks;df")), bytes(random, 10, 11, []byte("1234123lks;df")))
}

// Bytes
func TestBytes2(t *testing.T) {
	a := assert.New(t, false)

	// 测试固定长度
	a.Equal(len(Bytes(8, 9, []byte("1ks;dfp123;4j;ladj;fpoqwe"))), 8)

	// 非固定长度
	l := len(Bytes(8, 10, []byte("adf;wieqpwekwjerpq")))
	a.True(l >= 8 && l <= 10)
}

func TestRandsBuffer(t *testing.T) {
	a := assert.New(t, false)

	a.PanicString(func() {
		New(10000134, 0, 5, 7, []byte(";adkfjpqwei12124nbnb"))
	}, "bufferSize 必须大于零")

	r := New(10000134, 2, 5, 7, []byte(";adkfjpqwei12124nbnb"))
	a.NotNil(r)
	ctx, cancel := context.WithCancel(context.Background())
	go r.Serve(ctx)
	time.Sleep(time.Microsecond * 500) // 等待 go 运行完成
	a.Equal(cap(r.channel), 2)

	a.NotEqual(r.String(), r.String())
	a.NotEqual(r.String(), r.String())
	a.NotEqual(r.Bytes(), r.Bytes())

	cancel()
	time.Sleep(time.Microsecond * 500) // 等待 cancel 运行完成
	a.NotEqual(r.String(), r.String()) // 读取 channel 中剩余的数据
	a.Equal(r.String(), r.String())    // 没有数据了，都是空值
}
