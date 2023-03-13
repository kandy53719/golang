package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(cur *ListNode) *ListNode {
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

func main() {

}
