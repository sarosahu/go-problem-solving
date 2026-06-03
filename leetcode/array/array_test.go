package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
    got := TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println("got : ", got)
    want := []int{0, 1}
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
}

func TestLongestCommonPrefix(t *testing.T) {
	got := LongestCommonPrefix([]string{"flower","flow","flight"})
	want := "fl"
	if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
}
