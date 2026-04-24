package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	var list1data []int = []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1} // Cant store a number this large
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
	var list2data []int = []int{5, 6, 4}
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
	result := addTwoNumbers(head1, head2)
	fmt.Print(result)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Add first list together
	sum1 := AddNumbersinLinkedList(l1)
	fmt.Println(sum1)
	sum2 := AddNumbersinLinkedList(l2)
	fmt.Println(sum2)

	var resultSum *big.Int = new(big.Int)
	resultSum.Add(sum1, sum2)

	fmt.Println(resultSum)
	resultMod := new(big.Int).Mod(resultSum, new(big.Int).SetInt64(10))
	resultSum.Div(resultSum, new(big.Int).SetInt64(10))
	var head *ListNode = &ListNode{int(resultMod.Uint64()), nil}
	var current *ListNode = head
	for {
		if resultSum.BitLen() != 0 {
			resultMod := new(big.Int).Mod(resultSum, new(big.Int).SetInt64(10))
			resultSum.Div(resultSum, new(big.Int).SetInt64(10))
			newNode := &ListNode{int(resultMod.Uint64()), nil}
			current.Next = newNode
			current = newNode // Moves current pointer to next
		} else {
			break
		}
	}

	return head
}

func AddNumbersinLinkedList(list *ListNode) *big.Int {
	result := new(big.Int)
	// If list is a single node return node val - no computation needed
	if list.Next == nil {
		//return big.NewInt(int64(list.Val))
		return result.SetInt64(int64(list.Val))
	}
	var strNum string
	for list != nil {
		strNum = strconv.Itoa(list.Val) + strNum
		list = list.Next
	}
	result.SetString(strNum, 10)
	return result
}
