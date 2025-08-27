package middle_of_a_linked_list

import (
	"testing"

	linkedlists "github.com/mThanuj/dsa/datastructures/linked_lists"
	testutils "github.com/mThanuj/dsa/test_utils"
)

func Test1(t *testing.T) {
	ll := linkedlists.LinkedList[int]{}

	inputs := []int{1, 2, 3, 4, 5}
	for _, n := range inputs {
		ll.Add(n)
	}

	got := MiddleOfALinkedList(&ll)
	want := 3

	if err := testutils.Equals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}

func Test2(t *testing.T) {
	ll := linkedlists.LinkedList[int]{}

	inputs := []int{1, 2, 3, 4}
	for _, n := range inputs {
		ll.Add(n)
	}

	got := MiddleOfALinkedList(&ll)
	want := 3

	if err := testutils.Equals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}
