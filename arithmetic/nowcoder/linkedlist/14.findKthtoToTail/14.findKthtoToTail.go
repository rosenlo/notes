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
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param pHead ListNode类
 * @param k int整型
 * @return ListNode类
 */
func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	if pHead == nil || k == 0 {
		return nil
	}
	arr := make([]*ListNode, 0)
	for ; pHead != nil; pHead = pHead.Next {
		arr = append(arr, pHead)
	}
	if k > len(arr) {
		return nil
	}
	return arr[len(arr)-k]
}
