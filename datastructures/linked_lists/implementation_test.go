package linkedlists

import (
	"testing"

	testutils "github.com/mThanuj/dsa/test_utils"
)

func TestAdd(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	got := ll.ToSlice()
	want := []int{10, 20, 30}

	if err := testutils.SliceEquals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}

func TestInsertAtBeginning(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	ll.InsertAtBeginning(5)

	got := ll.ToSlice()
	want := []int{5, 10, 20, 30}

	if err := testutils.SliceEquals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}

func TestInsertAt(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	err := ll.InsertAt(15, 1)
	if err != nil {
		t.Fatal(err.Error())
	}

	got := ll.ToSlice()
	want := []int{10, 15, 20, 30}

	if err = testutils.SliceEquals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}

func TestRemove(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	ll.Remove(10)

	got := ll.ToSlice()
	want := []int{20, 30}

	if err := testutils.SliceEquals(got, want); err != nil {
		t.Fatal(err.Error())
	}

	ll.InsertAtBeginning(10)
	ll.InsertAtBeginning(10)
	ll.Add(40)

	ll.Remove(10)

	got = ll.ToSlice()
	want = []int{10, 20, 30, 40}

	if err := testutils.SliceEquals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}

func TestRemoveAt(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	err := ll.RemoveAt(1)
	if err != nil {
		t.Fatal(err.Error())
	}

	got := ll.ToSlice()
	want := []int{10, 30}

	if err = testutils.SliceEquals(got, want); err != nil {
		t.Fatal(err.Error())
	}

	ll.InsertAt(20, 1)
	ll.Add(40)

	ll.RemoveAt(0)

	got = ll.ToSlice()
	want = []int{20, 30, 40}

	if err = testutils.SliceEquals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}

func TestRemoveFirst(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	ll.RemoveFirst()

	got := ll.ToSlice()
	want := []int{20, 30}

	if err := testutils.SliceEquals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}

func TestRemoveLast(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)

	ll.RemoveLast()

	got := ll.ToSlice()
	want := []int{10}

	if err := testutils.SliceEquals(got, want); err != nil {
		t.Fatal(err.Error())
	}

	ll.RemoveLast()

	got = ll.ToSlice()
	want = []int{}

	if err := testutils.SliceEquals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}

func TestFind(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	got := ll.Find(10)
	want := true

	if err := testutils.Equals(got, want); err != nil {
		t.Fatal(err.Error())
	}

	got = ll.Find(50)
	want = false

	if err := testutils.Equals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}

func TestIndexOf(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	got := ll.IndexOf(10)
	want := 0

	if err := testutils.Equals(got, want); err != nil {
		t.Fatal(err.Error())
	}

	got = ll.IndexOf(50)
	want = -1

	if err := testutils.Equals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}

func TestLength(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	got := ll.Length()
	want := 3

	if err := testutils.Equals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}

func TestIsEmpty(t *testing.T) {
	ll := LinkedList[int]{}

	got := ll.IsEmpty()
	want := true

	if err := testutils.Equals(got, want); err != nil {
		t.Fatal(err.Error())
	}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	got = ll.IsEmpty()
	want = false

	if err := testutils.Equals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}

func TestClear(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	ll.Clear()

	got := ll.Length()
	want := 0

	if err := testutils.Equals(got, want); err != nil {
		t.Fatal(err.Error())
	}
}
