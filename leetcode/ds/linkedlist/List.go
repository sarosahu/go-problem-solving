package linkedlist

import (
	"errors"
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
	Size int
}

// NewLinkedList creates a new initialized empty list
func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Size: 0,
	}
}

// Prepend adds a new node at the very beginning (O(1) time)
func (l *LinkedList) Prepend(val int) {
	newNode := &ListNode{Val: val, Next: l.Head}
	l.Head = newNode
	l.Size++
}

// Append adds a new node at the very end (O(N) time)
func (l *LinkedList) Append(val int) {
	newNode := &ListNode{Val: val}

	// Handle the edge case of an empty list
	if l.Head == nil {
		l.Head = newNode
		l.Size++
		return
	}
	
	// Traverse to the last element
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
	l.Size++
}

// Delete removes the first occurrence of a value from the list
func (l *LinkedList) Delete(val int) error {
	if l.Head == nil {
		return errors.New("list is empty")
	}

	// Edge case: target is at the head node
	if l.Head.Val == val {
		l.Head = l.Head.Next
		l.Size--
		return nil
	}

	// Traverse and find the parent node of the target value
	curr := l.Head
	for curr.Next != nil {
		if curr.Next.Val ==  val {
			curr.Next = curr.Next.Next
			l.Size--
			return nil
		}
		curr = curr.Next
	}
	return errors.New("value not found")
}

// Display prints out the visual path of the list
func (l *LinkedList) Display() {
	curr := l.Head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}

