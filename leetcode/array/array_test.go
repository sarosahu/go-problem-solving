package array

import (
	"fmt"
	"reflect"
	"testing"
)

// helper function to convert a slice of ints into a linked list
func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := &ListNode{}
	curr := dummy
	for _, val := range nums {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return dummy.Next
}

// helper function to convert a linked list back into a slice for easy comparison
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// helper function to compare two integer slices
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}



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

	got = LongestCommonPrefix2([]string{"flower","flow","flight"})
	want = "fl"
	if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
}

func TestIsValidSudoku(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

    got := IsValidSudoku(board)
	fmt.Println("got : ", got)
    want := true
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
	board = [][]byte{
		{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
		{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
		{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
		{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	got = IsValidSudoku(board)
	fmt.Println("got : ", got)
    want = false
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
}

func TestMergeSort(t *testing.T) {
    got := ([]int{2, 7, 11, 15})
	fmt.Println("got : ", got)
    want := []int{0, 1}
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
}


func TestMergeKLists(t *testing.T) {
	// Define the test cases using a table-driven approach
	tests := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name: "Standard case with multiple lists",
			input: [][]int{
				{1, 4, 5},
				{1, 3, 4},
				{2, 6},
			},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:     "Empty outer slice",
			input:    [][]int{},
			expected: []int{},
		},
		{
			name: "Slice containing empty lists",
			input: [][]int{
				{},
				{},
			},
			expected: []int{},
		},
		{
			name: "Mix of empty and populated lists",
			input: [][]int{
				{},
				{2, 3},
				{},
				{1},
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "Lists with identical elements",
			input: [][]int{
				{2, 2},
				{2},
			},
			expected: []int{2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 1. Convert input slices to list nodes
			var lists []*ListNode
			for _, subSlice := range tt.input {
				lists = append(lists, sliceToList(subSlice))
			}

			// 2. Run the function under test
			actualList := MergeKLists(lists)

			// 3. Convert result back to slice
			actualSlice := listToSlice(actualList)

			// 4. Assert the result matches expected output
			if !equalSlices(actualSlice, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, actualSlice)
			}
		})
	}
}