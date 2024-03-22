// SPDX-FileCopyrightText: 2016-2024 caixw
//
// SPDX-License-Identifier: MIT

package rands

import (
	"crypto/rand"
	"testing"
)

func BenchmarkString(b *testing.B) {
	b.Run("bytes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			String(6, 7, []byte("abcdefghijklmnopkrstuvwxyz"))
		}
	})

	b.Run("runes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			String(6, 7, []rune("中文内容也可以正常显示"))
		}
	})
}

// 固定长度的随机字符串
func BenchmarkBytes_6_7_Lower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(6, 7, []byte("abcdefghijklmnopkrstuvwxyz"))
	}
}

// 固定长度的随机字符串
func BenchmarkBytes_6_7_All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(6, 7, AlphaNumberPunct())
	}
}

func BenchmarkBytes_4_6_Lower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(4, 6, []byte("abcdefghijklmnopkrstuvwxyz"))
	}
}

func BenchmarkBytes_4_6_All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(4, 6, AlphaNumberPunct())
	}
}

func BenchmarkBytes_10_32_Lower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(10, 32, []byte("abcdefghijklmnopkrstuvwxyz"))
	}
}

func BenchmarkBytes_10_32_All(b *testing.B) {
	b.Run("bytes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Bytes(10, 32, AlphaNumberPunct())
		}
	})

	b.Run("runes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Bytes(10, 32, []rune("中文内容也可以正常显示中文内容也可以正常显示"))
		}
	})
}

// crypto/rand包的随机读取能力
func BenchmarkCryptoRand_10(b *testing.B) {
	bs := make([]byte, 10)
	for i := 0; i < b.N; i++ {
		rand.Read(bs)
	}
}
