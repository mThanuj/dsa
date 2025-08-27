package middle_of_a_linked_list

import (
	linkedlist "github.com/mThanuj/dsa/datastructures/linked_lists"
)

func MiddleOfALinkedList[T comparable](ll *linkedlist.LinkedList[T]) T {
	slow, fast := ll.Head, ll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow.Data
}
