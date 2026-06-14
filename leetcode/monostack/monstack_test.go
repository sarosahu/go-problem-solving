package stack

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "Standard LeetCode Case 1",
			nums1: []int{4, 1, 2},
			nums2: []int{1, 3, 4, 2},
			want:  []int{-1, 3, -1},
		},
		{
			name:  "Standard LeetCode Case 2",
			nums1: []int{2, 4},
			nums2: []int{1, 2, 3, 4},
			want:  []int{3, -1},
		},
		{
			name:  "All Elements Decreasing",
			nums1: []int{3, 2, 1},
			nums2: []int{3, 2, 1},
			want:  []int{-1, -1, -1},
		},
		{
			name:  "All Elements Increasing",
			nums1: []int{1, 2, 3},
			nums2: []int{1, 2, 3, 4},
			want:  []int{2, 3, 4},
		},
		{
			name:  "Single Element Match",
			nums1: []int{5},
			nums2: []int{5},
			want:  []int{-1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NextGreaterElement(tt.nums1, tt.nums2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
