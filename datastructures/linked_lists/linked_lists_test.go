package linkedlists

import (
	"errors"
	"fmt"
	"testing"
)

func assert[T comparable](got []T, want []T) error {
	if len(got) != len(want) {
		return errors.New("length of got and want mismatch")
	}

	for i := range got {
		if got[i] != want[i] {
			return fmt.Errorf("at index %d: got %v, but wanted %v", i, got[i], want[i])
		}
	}

	return nil
}

func assertEquals[T comparable](got T, want T) error {
	if got != want {
		return fmt.Errorf("got %v, but wanted %v", got, want)
	}

	return nil
}

func TestAdd(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	got := ll.ToSlice()
	want := []int{10, 20, 30}

	err := assert(got, want)
	if err != nil {
		t.Error(err.Error())
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

	err := assert(got, want)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestInsertAt(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	err := ll.InsertAt(15, 1)
	if err != nil {
		t.Error(err.Error())
		return
	}

	got := ll.ToSlice()
	want := []int{10, 15, 20, 30}

	err = assert(got, want)
	if err != nil {
		t.Error(err.Error())
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

	err := assert(got, want)
	if err != nil {
		t.Error(err.Error())
		return
	}

	ll.InsertAtBeginning(10)
	ll.InsertAtBeginning(10)
	ll.Add(40)

	ll.Remove(10)

	got = ll.ToSlice()
	want = []int{10, 20, 30, 40}

	err = assert(got, want)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestRemoveAt(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	err := ll.RemoveAt(1)
	if err != nil {
		t.Error(err.Error())
		return
	}

	got := ll.ToSlice()
	want := []int{10, 30}

	err = assert(got, want)
	if err != nil {
		t.Error(err.Error())
		return
	}

	ll.InsertAt(20, 1)
	ll.Add(40)

	ll.RemoveAt(0)

	got = ll.ToSlice()
	want = []int{20, 30, 40}

	err = assert(got, want)
	if err != nil {
		t.Error(err.Error())
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

	err := assert(got, want)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestRemoveLast(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)

	ll.RemoveLast()

	got := ll.ToSlice()
	want := []int{10}

	err := assert(got, want)
	if err != nil {
		t.Error(err.Error())
		return
	}

	ll.RemoveLast()

	got = ll.ToSlice()
	want = []int{}

	err = assert(got, want)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestFind(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	got := ll.Find(10)
	want := true

	err := assertEquals(got, want)
	if err != nil {
		t.Error(err.Error())
		return
	}

	got = ll.Find(50)
	want = false

	err = assertEquals(got, want)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestIndexOf(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	got := ll.IndexOf(10)
	want := 0

	err := assertEquals(got, want)
	if err != nil {
		t.Error(err.Error())
		return
	}

	got = ll.IndexOf(50)
	want = -1

	err = assertEquals(got, want)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestLength(t *testing.T) {
	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	got := ll.Length()
	want := 3

	err := assertEquals(got, want)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestIsEmpty(t *testing.T) {
	ll := LinkedList[int]{}

	got := ll.IsEmpty()
	want := true

	err := assertEquals(got, want)
	if err != nil {
		t.Error(err.Error())
		return
	}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	got = ll.IsEmpty()
	want = false

	err = assertEquals(got, want)
	if err != nil {
		t.Error(err.Error())
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

	err := assertEquals(got, want)
	if err != nil {
		t.Error(err.Error())
	}
}
