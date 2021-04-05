package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestPalindrome1(s string) string {
	l, r, b, e := 0, 0, 0, 0
	for i := 0; i < len(s)-1; i++ {
	loop:
		for j := i; j < len(s); j++ {
			l = i
			r = j
			for l < r {
				if s[l] != s[r] {
					continue loop
				}
				l++
				r--
			}
			if j-i > e-b {
				b = i
				e = j
			}
		}
	}

	return s[b : e+1]
}

func longestPalindrome(s string) string {
	n := len(s)
	for l := n - 1; l > 0; l-- {
	loop:
		for i := 0; i+l < n; {
			for x := 0; x < l-x; x++ {
				if s[i+x] != s[i+l-x] {
					i++
					continue loop
				}
			}
			return s[i : i+l+1]
		}
	}
	return s[0:1]
}

func TestLongestPalindrome(t *testing.T) {
	var s string
	s1 := "babad"
	s = longestPalindrome(s1)
	assert.Contains(t, []string{"bab", "aba"}, s)

	s2 := "cbba"
	s = longestPalindrome(s2)
	assert.Contains(t, []string{"bb"}, s)

	s3 := "a"
	s = longestPalindrome(s3)
	assert.Contains(t, []string{"a"}, s)

	s4 := "ac"
	s = longestPalindrome(s4)
	assert.Contains(t, []string{"a"}, s)

	s5 := "aacabdkacaa"
	s = longestPalindrome(s5)
	assert.Contains(t, []string{"aca"}, s)
}
