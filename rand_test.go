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

func TestBytes1(t *testing.T) {
	a := assert.New(t)

	a.NotEqual(bytes(10, []int{Lower}), bytes(10, []int{Lower}))
	a.NotEqual(bytes(10, []int{Lower}), bytes(10, []int{Lower}))
	a.NotEqual(bytes(10, []int{Upper}), bytes(10, []int{Lower}))

	a.NotEqual(bytes(10, []int{Lower, Digit}), bytes(10, []int{Lower, Digit}))
}

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
}

// 固定长度的随机字符串
// BenchmarkBytes_6_7_Lower-4  	 5000000	       261 ns/op
func BenchmarkBytes_6_7_Lower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(6, 7, Lower)
	}
}

// 固定长度的随机字符串
// BenchmarkBytes_6_7_All-4    	 5000000	       253 ns/op
func BenchmarkBytes_6_7_All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(6, 7, Lower, Upper, Digit, Punct)
	}
}

// BenchmarkBytes_4_6_Lower-4  	10000000	       223 ns/op
func BenchmarkBytes_4_6_Lower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(4, 6, Lower)
	}
}

// BenchmarkBytes_4_6_All-4    	10000000	       221 ns/op
func BenchmarkBytes_4_6_All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(4, 6, Lower, Upper, Digit, Punct)
	}
}

// BenchmarkBytes_10_32_Lower-4	 2000000	       667 ns/op
func BenchmarkBytes_10_32_Lower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(10, 32, Lower)
	}
}

// BenchmarkBytes_10_32_All-4  	 2000000	       664 ns/op
func BenchmarkBytes_10_32_All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(10, 32, Lower, Upper, Digit, Punct)
	}
}
