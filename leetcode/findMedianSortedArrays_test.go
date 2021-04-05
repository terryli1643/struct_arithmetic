package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	s := make([]int, len(nums1)+len(nums2))
	n := len(nums1) - 1
	m := len(nums2) - 1
	t := m + n + 2
	for i := range s {
		if n < 0 {
			s[i] = nums2[m]
			m--
			continue
		}
		if m < 0 {
			s[i] = nums1[n]
			n--
			continue
		}

		if nums1[n] > nums2[m] {
			s[i] = nums1[n]
			n--
		} else {
			s[i] = nums2[m]
			m--
		}
	}
	if t%2 == 0 {
		return float64(s[t/2-1]+s[t/2]) / 2
	} else {
		return float64(s[t/2])
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	var n float64
	n1 := []int{1, 3}
	n2 := []int{2}
	n = findMedianSortedArrays(n1, n2)
	assert.Equal(t, float64(2), n)

	n3 := []int{1, 3}
	n4 := []int{2, 4}
	n = findMedianSortedArrays(n3, n4)
	assert.Equal(t, float64(2.5), n)
}
