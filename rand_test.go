// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package rand

import (
	"testing"

	"github.com/issue9/assert"
)

func TestCheckArgs(t *testing.T) {
	a := assert.New(t)

	// min < 0
	a.Error(checkArgs(-1, 1, []int{Lower}))

	// max <= min
	a.Error(checkArgs(5, 5, []int{Lower}))

	// cats为空
	a.Error(checkArgs(5, 6, []int{}))

	// cats的取值非法
	a.Error(checkArgs(5, 6, []int{100, 101}))
	a.Error(checkArgs(5, 6, []int{-1, -2}))

	a.NotError(checkArgs(5, 6, []int{Lower, Upper}))
}

// bytes
func TestBytes1(t *testing.T) {
	a := assert.New(t)

	a.NotEqual(bytes(random, 10, []int{Lower}), bytes(random, 10, []int{Lower}))
	a.NotEqual(bytes(random, 10, []int{Lower}), bytes(random, 10, []int{Lower}))
	a.NotEqual(bytes(random, 10, []int{Upper}), bytes(random, 10, []int{Lower}))

	a.NotEqual(bytes(random, 10, []int{Lower, Digit}), bytes(random, 10, []int{Lower, Digit}))
}

// Bytes
func TestBytes2(t *testing.T) {
	a := assert.New(t)

	// 测试固定长度
	a.Equal(len(Bytes(8, 9, Lower)), 8)

	// 非固定长度
	l := len(Bytes(8, 10, Lower))
	a.True(l >= 8 && l <= 10)
}

func TestString(t *testing.T) {
	t.Log("String(8,10,Lower):", String(8, 10, Lower))
	t.Log("String(8,10,Upper):", String(8, 10, Upper))
	t.Log("String(8,10,Digit):", String(8, 10, Digit))
	t.Log("String(8,10,Punct):", String(8, 10, Punct))
	t.Log("String(8,10,Lower, Punct):", String(8, 10, Lower, Punct))
	t.Log("String(8,10,Lower, Upper, Digit, Punct):", String(8, 10, Lower, Upper, Digit, Punct))
}

func TestRandNoBuffer(t *testing.T) {
	a := assert.New(t)

	r, err := New(0, 0, 5, 7, Lower, Digit)
	a.NotError(err).NotNil(r)
	a.Equal(cap(r.channel), 0)

	a.NotEqual(r.Bytes(), r.Bytes())
	a.NotEqual(r.Bytes(), r.Bytes())
	a.NotEqual(r.String(), r.String())
}

func TestRandBuffer(t *testing.T) {
	a := assert.New(t)

	r, err := New(10000134, 100, 5, 7, Lower, Digit)
	a.NotError(err).NotNil(r)
	a.Equal(cap(r.channel), 100)

	a.NotEqual(r.String(), r.String())
	a.NotEqual(r.String(), r.String())
	a.NotEqual(r.Bytes(), r.Bytes())
}
