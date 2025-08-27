package nth_node_from_end

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

	got := NthNodeFromEnd(&ll, 2)
	want := 4

	if err := testutils.Equals(got, want); err != nil {
		t.Fatal(err.Error())
	}

	got = NthNodeFromEnd(&ll, 5)
	want = 1

	if err := testutils.Equals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}

func Test2(t *testing.T) {
	ll := linkedlists.LinkedList[int]{}

	inputs := []int{1, 2, 3, 4, 5}
	for _, n := range inputs {
		ll.Add(n)
	}

	got := NthNodeFromEnd(&ll, 5)
	want := 1

	if err := testutils.Equals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}
