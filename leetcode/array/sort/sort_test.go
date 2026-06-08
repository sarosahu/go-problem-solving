package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
    got := MergeSort([]int{20, 10, 11, 5, 2, 1, 3, 6, 3, 9, 12, 4})
	fmt.Println("got : ", got)
    want := []int{1, 2, 3, 3, 4, 5, 6, 9, 10, 11, 12, 20}
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
}

func TestHeapSort(t *testing.T) {
    got := HeapSort([]int{20, 10, 11, 5, 2, 1, 3, 6, 3, 9, 12, 4})
	fmt.Println("got : ", got)
    want := []int{1, 2, 3, 3, 4, 5, 6, 9, 10, 11, 12, 20}
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
}