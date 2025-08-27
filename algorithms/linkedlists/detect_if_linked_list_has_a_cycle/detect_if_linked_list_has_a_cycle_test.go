package detect_if_linked_list_has_a_cycle

import (
	"testing"

	linkedlists "github.com/mThanuj/dsa/datastructures/linked_lists"
	testutils "github.com/mThanuj/dsa/test_utils"
)

func Test1(t *testing.T) {
	ll := linkedlists.LinkedList[int]{}

	n5 := linkedlists.ListNode[int]{
		Data: 50,
	}
	n4 := linkedlists.ListNode[int]{
		Data: 40,
		Next: &n5,
	}
	n3 := linkedlists.ListNode[int]{
		Data: 30,
		Next: &n4,
	}
	n2 := linkedlists.ListNode[int]{
		Data: 20,
		Next: &n3,
	}
	n1 := linkedlists.ListNode[int]{
		Data: 10,
		Next: &n2,
	}
	n4.Next = &n2
	ll.Head = &n1

	got := DetectIfLinkedListHasACycle(&ll)
	want := true

	if err := testutils.Equals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}

func Test2(t *testing.T) {
	ll := linkedlists.LinkedList[int]{}

	n1 := linkedlists.ListNode[int]{
		Data: 10,
	}
	ll.Head = &n1

	got := DetectIfLinkedListHasACycle(&ll)
	want := false

	if err := testutils.Equals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}
