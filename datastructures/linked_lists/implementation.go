package linkedlists

import (
	"errors"
)

type ListNode[T comparable] struct {
	Data T
	Next *ListNode[T]
}

type LinkedList[T comparable] struct {
	Head  *ListNode[T]
	Count int
}

func (ll *LinkedList[T]) Add(data T) {
	if ll.Head == nil {
		ll.Head = &ListNode[T]{
			Data: data,
		}
		ll.Count += 1
		return
	}

	temp := ll.Head

	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Next = &ListNode[T]{
		Data: data,
		Next: nil,
	}
	ll.Count += 1
}

func (ll *LinkedList[T]) InsertAtBeginning(data T) {
	if ll.Head == nil {
		ll.Head = &ListNode[T]{
			Data: data,
		}
		ll.Count += 1
		return
	}

	node := &ListNode[T]{
		Data: data,
		Next: ll.Head,
	}

	ll.Head = node
	ll.Count += 1
}

func (ll *LinkedList[T]) InsertAt(data T, index int) error {
	if index < 0 || index >= ll.Count {
		return errors.New("invalid index")
	}

	temp := ll.Head

	for i := 0; i < index-1; i++ {
		temp = temp.Next
	}

	node := &ListNode[T]{
		Data: data,
		Next: temp.Next,
	}

	temp.Next = node

	ll.Count += 1

	return nil
}

func (ll *LinkedList[T]) Remove(data T) {
	if ll.Head != nil && ll.Head.Data == data {
		ll.Head = ll.Head.Next
		ll.Count -= 1
		return
	}

	temp := ll.Head

	for temp != nil {
		if temp.Next == nil {
			return
		}

		if temp.Next.Data == data {
			temp.Next = temp.Next.Next
			ll.Count -= 1
			return
		}
	}
}

func (ll *LinkedList[T]) RemoveAt(index int) error {
	if index < 0 || index >= ll.Count {
		return errors.New("invalid index")
	}

	if index == 0 {
		ll.RemoveFirst()
		return nil
	}

	temp := ll.Head

	for i := 0; i < index-1; i++ {
		temp = temp.Next
	}

	if temp.Next != nil {
		temp.Next = temp.Next.Next
	}

	ll.Count -= 1

	return nil
}

func (ll *LinkedList[T]) RemoveFirst() {
	if ll.Head != nil {
		ll.Head = ll.Head.Next
		ll.Count -= 1
		return
	}
}

func (ll *LinkedList[T]) RemoveLast() {
	if ll.Head == nil {
		return
	}

	temp := ll.Head

	if temp.Next == nil {
		ll.RemoveFirst()
		return
	}

	for temp.Next.Next != nil {
		temp = temp.Next
	}

	temp.Next = temp.Next.Next
	ll.Count -= 1
}

func (ll *LinkedList[T]) Find(data T) bool {
	temp := ll.Head

	for temp != nil {
		if temp.Data == data {
			return true
		}
		temp = temp.Next
	}

	return false
}

func (ll *LinkedList[T]) IndexOf(data T) int {
	temp := ll.Head
	index := 0

	for temp != nil {
		if temp.Data == data {
			return index
		}
		temp = temp.Next
	}

	return -1
}

func (ll *LinkedList[T]) ToSlice() []T {
	result := []T{}
	temp := ll.Head

	for temp != nil {
		result = append(result, temp.Data)
		temp = temp.Next
	}

	return result
}

func (ll *LinkedList[T]) Display() {
}

func (ll *LinkedList[T]) Length() int {
	return ll.Count
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.Count == 0
}

func (ll *LinkedList[T]) Clear() {
	ll.Head = nil
	ll.Count = 0
}
