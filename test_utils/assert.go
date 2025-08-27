package testutils

import (
	"errors"
	"fmt"
)

func SliceEquals[T comparable](got []T, want []T) error {
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

func Equals[T comparable](got T, want T) error {
	if got != want {
		return fmt.Errorf("got %v, but wanted %v", got, want)
	}

	return nil
}
