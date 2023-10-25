package buffers

import (
	"sourcecode.social/reiver/go-erorr"
)

var (
	BufferOverflow = erorr.Error("buffer overflow")
)

var (
	errNilDestination = erorr.Error("nil destination")
	errNilReader      = erorr.Error("nil reader")
	errNilReceiver    = erorr.Error("nil receiver")
	errShortWrite     = erorr.Error("short write")
)
