// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package rands

import "testing"

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
