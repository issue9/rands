// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package rands

import (
	"testing"

	"github.com/issue9/assert"
)

func TestCheckArgs(t *testing.T) {
	a := assert.New(t)

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
	a := assert.New(t)

	a.NotEqual(bytes(random, 10, []byte("1234123lks;df")), bytes(random, 10, []byte("1234123lks;df")))
	a.NotEqual(bytes(random, 10, []byte("1234123lks;df")), bytes(random, 10, []byte("1234123lks;df")))
	a.NotEqual(bytes(random, 10, []byte("1234123lks;df")), bytes(random, 10, []byte("1234123lks;df")))
}

// Bytes
func TestBytes2(t *testing.T) {
	a := assert.New(t)

	// 测试固定长度
	a.Equal(len(Bytes(8, 9, []byte("1ks;dfp123;4j;ladj;fpoqwe"))), 8)

	// 非固定长度
	l := len(Bytes(8, 10, []byte("adf;wieqpwekwjerpq")))
	a.True(l >= 8 && l <= 10)
}

func TestRandsNoBuffer(t *testing.T) {
	a := assert.New(t)

	r, err := New(0, 0, 5, 7, []byte("ad;fqeqwejqw;ejnweqwer"))
	a.NotError(err).NotNil(r)
	a.Equal(cap(r.channel), 0)

	a.NotEqual(r.Bytes(), r.Bytes())
	a.NotEqual(r.Bytes(), r.Bytes())
	a.NotEqual(r.String(), r.String())
}

func TestRandsBuffer(t *testing.T) {
	a := assert.New(t)

	r, err := New(10000134, 2, 5, 7, []byte(";adkfjpqwei12124nbnb"))
	a.NotError(err).NotNil(r)
	a.Equal(cap(r.channel), 2)

	a.NotEqual(r.String(), r.String())
	a.NotEqual(r.String(), r.String())
	a.NotEqual(r.Bytes(), r.Bytes())

	r.Stop()
	a.NotEqual(r.String(), r.String()) // 读取 channel 中的数据
	a.Equal(r.String(), r.String())    // 没有数据了，都是空值
}
