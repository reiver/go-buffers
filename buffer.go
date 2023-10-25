package buffers

import (
	"io"
	"unsafe"

	"sourcecode.social/reiver/go-opt"
	"sourcecode.social/reiver/go-utf8"
)

// A Buffer is used to store a series of bytes.
//
// A Buffer can be used as is, without any type of initialization (if not doesn't want the Buffer to have a limit):
//
//	var buffer buffers.Buffer
//
// But if one wants the Buffer to have a limit, then it should be initialized like the following:
//
//	var buffer buffers.Buffer = buffers.LimitedBuffer(4194304) // 4 MB
type Buffer struct {
	data []byte
	max  opt.Optional[int]
}

var _ io.ByteWriter = &Buffer{}
var _ io.ReaderFrom = &Buffer{}
var _ io.Writer     = &Buffer{}

// LimitedBuffer returns a Buffer with a maximum size.
//
//	var buffer buffers.Buffer = buffers.LimitedBuffer(4194304) // 4 MB
func LimitedBuffer(max int) Buffer {
	return Buffer {
		max:opt.Something[int](max),
	}
}

// Len returns how many bytes have been accumulated in Buffer.
//
//	receiver.Len() == len(receiver.String())
func (receiver *Buffer) Len() int {
	if nil == receiver {
		return 0
	}

	return len(receiver.data)
}

// ReadFrom reads from the io.Reader. Copying bytes until it gets an io.EOF.
func (receiver *Buffer) ReadFrom(r io.Reader) (int64, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}
	if nil == r {
		return 0, errNilReader
	}

	return io.Copy(receiver, r)
}

// String returns the bytes that have been accumulated.
func (receiver *Buffer) String() string {
	if nil == receiver {
		return ""
	}

	if nil == receiver.data {
		return ""
	}

	return *(*string)(unsafe.Pointer(&receiver.data))
}

func (receiver *Buffer) Write(p []byte) (n int, err error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	var lenP int
	{
		lenP = len(p)
		if lenP <= 0 {
			return 0, nil
		}
	}

	{
		max, something := receiver.max.Get()
		if something {
			newLen := receiver.Len() + lenP

			if max < newLen {
				return 0, BufferOverflow
			}
		}
	}

	receiver.data = append(receiver.data, p...)

	return len(p), nil
}

// WriteByte appends another byte to the Buffer.
func (receiver *Buffer) WriteByte(c byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	var b [1]byte
	var p  []byte = b[:]
	b[0] = c

	{
		n, err := receiver.Write(p)
		if nil != err {
			return err
		}
		if expected, actual := len(b), n; expected != actual {
			return errShortWrite
		}
	}

	return nil
}

// WriteRune appends the UTF-8 encoded rune to the Buffer.
func (receiver *Buffer) WriteRune(r rune) (int, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	return utf8.WriteRune(receiver, r)
}

// WriteRune appends the string to the Buffer.
func (receiver *Buffer) WriteString(s string) (int, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	return io.WriteString(receiver, s)
}
