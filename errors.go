package buffers

import (
	"errors"
)

var (
	errNilDestination = errors.New("buffers: Nil Destination")
	errNilReceiver    = errors.New("buffers: Nil Receiver")
)
