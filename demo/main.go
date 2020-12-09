package main

import (
	"log"
	"math"
)

type Header map[string]string

type Request struct {
	HeaderLine string
	Body       []byte
	Header     Header
}

func main() {
	res := mergeKLists([]*ListNode{
		buildLink([]int{1, 4, 5}),
		buildLink([]int{1, 3, 4}),
		buildLink([]int{2, 6}),
	})
	for res != nil {
		log.Println(res.Val)
		res = res.Next
	}
}

func buildLink(values []int) *ListNode {
	var node *ListNode
	var head *ListNode
	for _, v := range values {
		if node == nil {
			node = &ListNode{Val: v}
			head = node
		} else {
			node.Next = &ListNode{Val: v}
			node = node.Next
		}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode
	var resHead *ListNode
	for {
		index := -1
		min := math.MaxInt32
		for i, node := range lists {
			if node == nil {
				continue
			}
			if min > node.Val {
				min = node.Val
				index = i
			}
		}
		if index != -1 {
			if res == nil {
				res = lists[index]
				resHead = res
			} else {
				res.Next = lists[index]
				res = res.Next
			}
			lists[index] = lists[index].Next
			//如果为空，直接删除，减少循环次数
			if lists[index] == nil {
				lists = append(lists[:index], lists[index+1:]...)
			}
		} else {
			break
		}
	}
	return resHead
}
