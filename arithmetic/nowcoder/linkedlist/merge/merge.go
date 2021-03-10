package main

import "github.com/rosenlo/toolkits/structure/linkedlist"

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */
type ListNode = linkedlist.ListNode

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil && pHead2 == nil {
		return nil
	}

	if pHead1 == nil {
		return pHead2
	}

	if pHead2 == nil {
		return pHead1
	}

	var nHead *ListNode
	if pHead1.Val < pHead2.Val {
		nHead, pHead1 = pHead1, pHead1.Next
	} else {
		nHead, pHead2 = pHead2, pHead2.Next
	}
	ret := nHead

	for pHead1 != nil && pHead2 != nil {

		if pHead1.Val < pHead2.Val {
			nHead.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			nHead.Next = pHead2
			pHead2 = pHead2.Next
		}
		nHead = nHead.Next
	}

	if pHead1 != nil {
		nHead.Next = pHead1
	}
	if pHead2 != nil {
		nHead.Next = pHead2
	}
	return ret
}
