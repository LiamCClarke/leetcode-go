package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var list1data []int = []int{-9, 3}
	var head1 *ListNode
	var current1 *ListNode

	for _, val := range list1data {
		newNode := &ListNode{Val: val}

		if head1 == nil {
			head1 = newNode
			current1 = newNode
		} else {
			current1.Next = newNode
			current1 = newNode
		}
	}
	var list2data []int = []int{5, 7}
	var head2 *ListNode
	var current2 *ListNode

	for _, val := range list2data {
		newNode := &ListNode{Val: val}

		if head2 == nil {
			head2 = newNode
			current2 = newNode
		} else {
			current2.Next = newNode
			current2 = newNode
		}
	}

	mergedHead := mergeTwoLists(head1, head2)
	fmt.Print(mergedHead)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	var current *ListNode

	var list1pointer *ListNode = list1
	var list2pointer *ListNode = list2

	for true {
		if list1 == nil {
			return list2
		} else if list2 == nil {
			return list1
		}
		newNode := &ListNode{}
		if head == nil {
			head = newNode
			current = head
		}
		if list1pointer == nil {
			current.Next = list2pointer
			break
		}
		if list2pointer == nil {
			current.Next = list1pointer
			break
		}
		if list1pointer.Val <= list2pointer.Val {
			newNode.Val = list1pointer.Val
			current.Next = newNode           // Sets the next node as the newNode
			current = newNode                // Moves the current pointer to the newNode
			list1pointer = list1pointer.Next // Advance the list pointer
		} else {
			newNode.Val = list2pointer.Val
			current.Next = newNode           // Sets the next node as the newNode
			current = newNode                // Moves the current pointer to the newNode
			list2pointer = list2pointer.Next // Advance the list pointer
		}
	}
	return head
}
