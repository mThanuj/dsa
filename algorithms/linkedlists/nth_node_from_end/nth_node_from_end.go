package nth_node_from_end

import linkedlists "github.com/mThanuj/dsa/datastructures/linked_lists"

func NthNodeFromEnd[T comparable](ll *linkedlists.LinkedList[T], n int) T {
	slow, fast := ll.Head, ll.Head

	for ; n > 0 && fast != nil; n-- {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow.Data
}
