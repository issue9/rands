// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package rands

import (
	"crypto/rand"
	"testing"
)

// 固定长度的随机字符串
// BenchmarkBytes_6_7_Lower-4  	 5000000	       245 ns/op
func BenchmarkBytes_6_7_Lower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(6, 7, []byte("abcdefghijklmnopkrstuvwxyz"))
	}
}

// 固定长度的随机字符串
// BenchmarkBytes_6_7_All-4    	 5000000	       298 ns/op
func BenchmarkBytes_6_7_All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(6, 7, []byte("abcdefghijklmnopkrstuvwxyzABCDEFGHIJKLMNOPKRSTUVWXYZ1234567890!@#$%^&*()_+~|[]{};':\",./<>?\\|"))
	}
}

// BenchmarkBytes_4_6_Lower-4  	10000000	       211 ns/op
func BenchmarkBytes_4_6_Lower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(4, 6, []byte("abcdefghijklmnopkrstuvwxyz"))
	}
}

// BenchmarkBytes_4_6_All-4    	 5000000	       264 ns/op
func BenchmarkBytes_4_6_All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(4, 6, []byte("abcdefghijklmnopkrstuvwxyzABCDEFGHIJKLMNOPKRSTUVWXYZ1234567890!@#$%^&*()_+~|[]{};':\",./<>?\\|"))
	}
}

// BenchmarkBytes_10_32_Lower-4	 2000000	       618 ns/op
func BenchmarkBytes_10_32_Lower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(10, 32, []byte("abcdefghijklmnopkrstuvwxyz"))
	}
}

// BenchmarkBytes_10_32_All-4  	 2000000	       658 ns/op
func BenchmarkBytes_10_32_All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(10, 32, []byte("abcdefghijklmnopkrstuvwxyzABCDEFGHIJKLMNOPKRSTUVWXYZ1234567890!@#$%^&*()_+~|[]{};':\",./<>?\\|"))
	}
}

// crypto/rand包的随机读取能力
func BenchmarkCryptoRand(b *testing.B) {
	bs := make([]byte, 10, 10)
	for i := 0; i < b.N; i++ {
		rand.Read(bs)
	}
}
