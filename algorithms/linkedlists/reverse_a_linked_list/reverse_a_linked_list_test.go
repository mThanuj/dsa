package reverse_a_linked_list

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

	ReverseLinkedList(&ll)

	got := ll.ToSlice()
	want := []int{5, 4, 3, 2, 1}

	if err := testutils.SliceEquals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}

func Test2(t *testing.T) {
	ll := linkedlists.LinkedList[int]{}

	ReverseLinkedList(&ll)

	got := ll.ToSlice()
	want := []int{}

	if err := testutils.SliceEquals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}
