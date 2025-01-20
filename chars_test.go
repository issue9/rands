// SPDX-FileCopyrightText: 2024 caixw
//
// SPDX-License-Identifier: MIT

package rands

import (
	"testing"

	"github.com/issue9/assert/v4"
)

func TestChars(t *testing.T) {
	a := assert.New(t, false)

	s := Alpha()
	a.Equal(s[0], 'a').
		Equal(s[len(s)-1], 'Z')

	s = LowerAlpha()
	a.Equal(s[0], 'a').
		Equal(s[len(s)-1], 'z')

	s = UpperAlpha()
	a.Equal(s[0], 'A').
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
