package main

import (
	"github.com/rosenlo/toolkits/structure/linkedlist"
)

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

type ListNode = linkedlist.ListNode

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}
	var preNode *ListNode
	for pHead != nil {
		pHead.Next, preNode, pHead = preNode, pHead, pHead.Next
	}

	return preNode
}
