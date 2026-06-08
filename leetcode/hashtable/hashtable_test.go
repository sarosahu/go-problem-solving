package hashtable

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
    got := LengthOfLongestSubstring("abcdeca")
	fmt.Println("got : ", got)
    want := 5
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
}

func TestLongestConsecutiveSeq(t *testing.T) {
    got := LongestConsecutive([]int{0,3,7,2,5,8,4,6,0,1})
	fmt.Println("got : ", got)
    want := 9
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
}