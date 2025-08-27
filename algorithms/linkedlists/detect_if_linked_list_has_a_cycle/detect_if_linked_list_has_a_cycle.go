package detect_if_linked_list_has_a_cycle

import linkedlists "github.com/mThanuj/dsa/datastructures/linked_lists"

func DetectIfLinkedListHasACycle[T comparable](ll *linkedlists.LinkedList[T]) bool {
	slow, fast := ll.Head, ll.Head

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
