package buffers

type Writer struct {
	dst []byte
	index int
}

func NewWriter(dst []byte) Writer {
	return Writer{
		dst:dst,
	}
}

func (receiver *Writer) Write(src []byte) (n int, err error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	dst := receiver.dst
	if nil == dst {
		return 0, errNilDestination
	}
	dst = dst[receiver.index:]

	if expectedAtLeast, actual := len(src), len(dst); actual < expectedAtLeast {
		return 0, errTooShort(uint64(expectedAtLeast), uint64(actual))
	}

	n = copy(dst, src)
	receiver.index += n

	return n, err
}
