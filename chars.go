// SPDX-FileCopyrightText: 2024 caixw
//
// SPDX-License-Identifier: MIT

package rands

var chars = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+[]{};':\",./<>?")

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
