package reverse_a_linked_list

import (
	linkedlist "github.com/mThanuj/dsa/datastructures/linked_lists"
)

func ReverseLinkedList[T comparable](ll *linkedlist.LinkedList[T]) {
	var prev *linkedlist.ListNode[T]
	cur := ll.Head

	if ll.Head == nil {
		return
	}

	for cur != nil {
		nxt := cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}

	ll.Head = prev
}
