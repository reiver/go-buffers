package buffers

import (
	"fmt"
)

type TooShort interface {
	error
	TooShort() (expectedAtLeast uint64, actual uint64)
}

func errTooShort(expectedAtLeast uint64, actual uint64) error {
	var e TooShort = &internalTooShort{
		expectedAtLeast:expectedAtLeast,
		actual:actual,
	}

	return e
}

type internalTooShort struct {
	expectedAtLeast uint64
	actual          uint64
}

func (receiver internalTooShort) Error() string {
	return fmt.Sprintf("buffers: Backing Buffer Is Too Short: expected-at-least=%d actual=%d", receiver.expectedAtLeast, receiver.actual)
}

func (receiver internalTooShort) TooShort() (expectedAtLeast uint64, actual uint64) {
	return receiver.expectedAtLeast, receiver.actual
}
