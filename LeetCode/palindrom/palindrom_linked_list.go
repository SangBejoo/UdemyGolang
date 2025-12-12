package palindrom

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindromeLinkedList(head *ListNode) bool {
	values := []int{}
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		if values[i] != values[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("This is a module for checking palindromic linked lists.")
	fmt.Println(isPalindromeLinkedList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}))
}
