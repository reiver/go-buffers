package buffers

import (
	"github.com/reiver/go-fck"
)

var (
	BufferOverflow = fck.Error("buffer overflow")
)

var (
	errNilDestination = fck.Error("nil destination")
	errNilReader      = fck.Error("nil reader")
	errNilReceiver    = fck.Error("nil receiver")
	errShortWrite     = fck.Error("short write")
)
