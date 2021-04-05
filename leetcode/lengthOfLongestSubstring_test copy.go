package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLongestSubstring(s string) int {
	var b, e, l int
	for ; e < len(s); e++ {
		for i := e - 1; i >= b; i-- {
			if s[i] == s[e] {
				b = i + 1
				break
			}
		}
		if l < e-b+1 {
			l = e - b + 1
		}
	}
	return l
}

func TestLengthOfLongestSubstring(t *testing.T) {
	var n = 0
	s1 := "aabaab!bb"
	n = lengthOfLongestSubstring(s1)
	assert.Equal(t, 3, n)

	s2 := "bbbbb"
	n = lengthOfLongestSubstring(s2)
	assert.Equal(t, 1, n)

	s3 := "pwwkew"
	n = lengthOfLongestSubstring(s3)
	assert.Equal(t, 3, n)

	s4 := "ctnkh"
	n = lengthOfLongestSubstring(s4)
	assert.Equal(t, 5, n)
}
