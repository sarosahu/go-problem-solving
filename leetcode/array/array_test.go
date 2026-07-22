package array

import (
	"fmt"
	"reflect"
	"sort"
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

func TestCanBeValid(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		locked string
		want   bool
	}{
		{
			name:   "LeetCode Example 1",
			s:      "))()))",
			locked: "010100",
			want:   true,
		},
		{
			name:   "LeetCode Example 2",
			s:      "()()",
			locked: "0000",
			want:   true,
		},
		{
			name:   "LeetCode Example 3",
			s:      ")",
			locked: "0",
			want:   false, // Odd length
		},
		{
			name:   "Already Valid and Locked",
			s:      "(())",
			locked: "1111",
			want:   true,
		},
		{
			name:   "Invalid and Locked",
			s:      ")(((",
			locked: "1111",
			want:   false,
		},
		{
			name:   "Too Many Closing Brackets Early",
			s:      "))))((",
			locked: "110000",
			want:   false,
		},
		{
			name:   "Empty String",
			s:      "",
			locked: "",
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanBeValid(tt.s, tt.locked)
			if got != tt.want {
				t.Errorf("canBeValid2() = %v, want %v for inputs s = %q, locked = %q", got, tt.want, tt.s, tt.locked)
			}
		})
	}
}

func BenchmarkCanBeValid(b *testing.B) {
	s := "))()))"
	locked := "010100"
	for i := 0; i < b.N; i++ {
		CanBeValid(s, locked)
	}
}

func TestCheckValidString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "LeetCode Example 1",
			s:    "()",
			want: true,
		},
		{
			name: "LeetCode Example 2",
			s:    "(*)",
			want: true, // * acts as empty string
		},
		{
			name: "LeetCode Example 3",
			s:    "(*))",
			want: true, // * acts as '('
		},
		{
			name: "Wildcard acts as closing bracket",
			s:    "((*",
			want: false, // * acts as ')'
		},
		{
			name: "Invalid sequence early on",
			s:    ")*",
			want: false,
		},
		{
			name: "Empty string",
			s:    "",
			want: true,
		},
		{
			name: "Multiple wildcards making it valid",
			s:    "***)))",
			want: true,
		},
		{
			name: "Too many closing brackets",
			s:    "(((**))))",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckValidString(tt.s)
			if got != tt.want {
				t.Errorf("checkValidString(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkCheckValidString(b *testing.B) {
	s := "(((***)))(((***)))"
	for i := 0; i < b.N; i++ {
		CheckValidString(s)
	}
}

func TestExistWordSearch(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		word     string
		expected bool
	}{
		{
			name: "Word Exists - Horizontal and Vertical",
			// [A B C E]
			// [S F C S]
			// [A D E E]
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCCED",
			expected: true,
		},
		{
			name: "Word Exists - Single Letter Match",
			board: [][]byte{
				{'A'},
			},
			word:     "A",
			expected: true,
		},
		{
			name: "Word Does Not Exist",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCB",
			expected: false, // Cannot reuse the same 'B'
		},
		{
			name: "Word Longer Than Grid Elements",
			board: [][]byte{
				{'A', 'B'},
			},
			word:     "ABC",
			expected: false,
		},
		{
			name:     "Empty Grid",
			board:    [][]byte{},
			word:     "A",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := exist(tt.board, tt.word)
			if actual != tt.expected {
				t.Errorf("expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Standard Example",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49, // Formed by 8 at index 1 and 7 at index 8. Width = 7, Height = 7. Area = 49.
		},
		{
			name:     "Minimum Length Array",
			height:   []int{1, 1},
			expected: 1, // Width = 1, Height = 1.
		},
		{
			name:     "Uniform Heights",
			height:   []int{5, 5, 5, 5, 5},
			expected: 20, // Max width is best: index 0 to index 4. Width = 4, Height = 5.
		},
		{
			name:     "Strictly Increasing Heights",
			height:   []int{1, 2, 3, 4, 5},
			expected: 6, // Formed by 2 at index 1 and 5 at index 4. Width = 3, Height = 2.
		},
		{
			name:     "Strictly Decreasing Heights",
			height:   []int{5, 4, 3, 2, 1},
			expected: 6, // Formed by 5 at index 0 and 2 at index 3. Width = 3, Height = 2.
		},
		{
			name:     "Deep Valley with High Edges",
			height:   []int{10, 1, 1, 1, 10},
			expected: 40, // Far edges match perfectly. Width = 4, Height = 10.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxArea(tt.height)
			if actual != tt.expected {
				t.Errorf("maxArea(%v) = %d; want %d", tt.height, actual, tt.expected)
			}
		})
	}
}

func TestCountSmaller(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Visual Walkthrough Example",
			nums:     []int{5, 2, 6, 1},
			expected: []int{2, 1, 1, 0},
		},
		{
			name:     "LeetCode Example with Negative Numbers",
			nums:     []int{-1, -1},
			expected: []int{0, 0},
		},
		{
			name:     "Handling Duplicates",
			nums:     []int{2, 5, 2, 6, 1},
			expected: []int{1, 2, 1, 1, 0}, // The first 2 only counts the '1' to its right, not the other '2'
		},
		{
			name:     "Strictly Decreasing",
			nums:     []int{5, 4, 3, 2, 1},
			expected: []int{4, 3, 2, 1, 0},
		},
		{
			name:     "Strictly Increasing",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			name:     "Empty Slice",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "Single Element",
			nums:     []int{10},
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countSmaller(tt.nums)
			
			// reflect.DeepEqual correctly compares the content and order of two slices
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("countSmaller(%v) = %v; want %v", tt.nums, actual, tt.expected)
			}
		})
	}
}

// helper to check if two slices contain identical elements ignoring their index sequence
func matchWithoutOrder(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	// Sort copies of the slices to safely compare them
	cA := append([]int(nil), a...)
	cB := append([]int(nil), b...)
	sort.Ints(cA)
	sort.Ints(cB)
	
	for i := range cA {
		if cA[i] != cB[i] {
			return false
		}
	}
	return true
}

func TestTopKFrequentAlgorithms(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Standard Example Matrix",
			//nums:     []int{1, 1, 1, 2, 2, 3},
			nums:     []int{5, 1, 4, 5, 2, 4, 5},
			k:        2,
			//expected: []int{1, 2},
			expected: []int{4, 5},
		},
		{
			name:     "Single Element Array",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "All Elements Have Uniform Frequencies",
			nums:     []int{4, 5, 6},
			k:        2,
			expected: []int{4, 5}, // {4, 6} or {5, 6} are also logically valid
		},
		{
			name:     "Negative Numbers Included",
			nums:     []int{-1, -1, -2, 2, 2, 2},
			k:        2,
			expected: []int{2, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 1. Validate the Bucket Sort Strategy
			bucketResult := topKFrequentBucketSort(tt.nums, tt.k)
			if !matchWithoutOrder(bucketResult, tt.expected) {
				// Special exception check for uniform frequency flexibility
				if tt.name != "All Elements Have Uniform Frequencies" || len(bucketResult) != tt.k {
					t.Errorf("topKFrequent() = %v; want elements matching %v", bucketResult, tt.expected)
				}
			}

			// 2. Validate the Min-Heap Strategy
			heapResult := topKFrequentMinHeap(tt.nums, tt.k)
			if !matchWithoutOrder(heapResult, tt.expected) {
				if tt.name != "All Elements Have Uniform Frequencies" || len(heapResult) != tt.k {
					t.Errorf("topKFrequentMinHeap() = %v; want elements matching %v", heapResult, tt.expected)
				}
			}
		})
	}
}

