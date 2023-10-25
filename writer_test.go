package buffers_test

import (
	"testing"

	"reflect"

	"sourcecode.social/reiver/go-buffers"
)

func TestWriterSingleWrite(t *testing.T) {

	tests := []struct {
		BufferLength int
		Src []byte
		Expected []byte
	}{
		{
			BufferLength: 0,
			Src: []byte("apple"),
			Expected: []byte{},
		},
		{
			BufferLength: 1,
			Src: []byte("apple"),
			Expected: []byte{},
		},
		{
			BufferLength: 2,
			Src: []byte("apple"),
			Expected: []byte{},
		},
		{
			BufferLength: 3,
			Src: []byte("apple"),
			Expected: []byte{},
		},
		{
			BufferLength: 4,
			Src: []byte("apple"),
			Expected: []byte{},
		},
		{
			BufferLength: 5,
			Src: []byte("apple"),
			Expected: []byte("apple"),
		},
		{
			BufferLength: 6,
			Src: []byte("apple"),
			Expected: []byte("apple"),
		},
		{
			BufferLength: 7,
			Src: []byte("apple"),
			Expected: []byte("apple"),
		},
		{
			BufferLength: 8,
			Src: []byte("apple"),
			Expected: []byte("apple"),
		},
		{
			BufferLength: 9,
			Src: []byte("apple"),
			Expected: []byte("apple"),
		},
		{
			BufferLength: 10,
			Src: []byte("apple"),
			Expected: []byte("apple"),
		},



		{
			BufferLength: 0,
			Src: []byte("apple BANANA"),
			Expected: []byte{},
		},
		{
			BufferLength: 1,
			Src: []byte("apple BANANA"),
			Expected: []byte{},
		},
		{
			BufferLength: 2,
			Src: []byte("apple BANANA"),
			Expected: []byte{},
		},
		{
			BufferLength: 3,
			Src: []byte("apple BANANA"),
			Expected: []byte{},
		},
		{
			BufferLength: 4,
			Src: []byte("apple BANANA"),
			Expected: []byte{},
		},
		{
			BufferLength: 5,
			Src: []byte("apple BANANA"),
			Expected: []byte{},
		},
		{
			BufferLength: 6,
			Src: []byte("apple BANANA"),
			Expected: []byte{},
		},
		{
			BufferLength: 7,
			Src: []byte("apple BANANA"),
			Expected: []byte{},
		},
		{
			BufferLength: 8,
			Src: []byte("apple BANANA"),
			Expected: []byte{},
		},
		{
			BufferLength: 9,
			Src: []byte("apple BANANA"),
			Expected: []byte{},
		},
		{
			BufferLength: 10,
			Src: []byte("apple BANANA"),
			Expected: []byte{},
		},
		{
			BufferLength: 11,
			Src: []byte("apple BANANA"),
			Expected: []byte{},
		},
		{
			BufferLength: 12,
			Src: []byte("apple BANANA"),
			Expected: []byte("apple BANANA"),
		},
		{
			BufferLength: 13,
			Src: []byte("apple BANANA"),
			Expected: []byte("apple BANANA"),
		},
		{
			BufferLength: 14,
			Src: []byte("apple BANANA"),
			Expected: []byte("apple BANANA"),
		},
		{
			BufferLength: 15,
			Src: []byte("apple BANANA"),
			Expected: []byte("apple BANANA"),
		},
		{
			BufferLength: 16,
			Src: []byte("apple BANANA"),
			Expected: []byte("apple BANANA"),
		},
		{
			BufferLength: 17,
			Src: []byte("apple BANANA"),
			Expected: []byte("apple BANANA"),
		},
		{
			BufferLength: 18,
			Src: []byte("apple BANANA"),
			Expected: []byte("apple BANANA"),
		},
		{
			BufferLength: 19,
			Src: []byte("apple BANANA"),
			Expected: []byte("apple BANANA"),
		},
		{
			BufferLength: 20,
			Src: []byte("apple BANANA"),
			Expected: []byte("apple BANANA"),
		},



		{
			BufferLength: 0,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 1,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 2,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 3,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 4,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 5,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 6,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 7,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 8,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 9,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 10,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 11,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 12,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 13,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 14,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 15,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 16,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 17,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 18,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte{},
		},
		{
			BufferLength: 19,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte("apple BANANA Cherry"),
		},
		{
			BufferLength: 20,
			Src: []byte("apple BANANA Cherry"),
			Expected: []byte("apple BANANA Cherry"),
		},
	}

	for testNumber, test := range tests {

		var buffer []byte = make([]byte, test.BufferLength)

		var dst *buffers.Writer = buffers.NewWriter(buffer)

		n, err := dst.Write(test.Src)
		if _, casted := err.(buffers.TooShort); !casted {
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("SRC:          %q", test.Src)
				t.Logf("DST:          %q (full)", buffer)
				t.Logf("DST EXPECTED: %q", test.Expected)
				t.Logf("ERROR TYPE: %T", err)
				t.Logf("ERROR: %q", err)
				continue
			}
		}

		var expectedWritten int
		{
			lenBuffer := len(buffer)
			lenSrc    := len(test.Src)

			if lenBuffer < lenSrc {
				expectedWritten = lenBuffer
			} else {
				expectedWritten = lenSrc
			}

			if _, casted := err.(buffers.TooShort); casted {
				expectedWritten = 0
			}
		}

		if expected, actual := expectedWritten, n; expected != actual {
			t.Errorf("For test #%d, the actual number of bytes written was not what was expected.", testNumber)
			t.Logf("SRC:          %q", test.Src)
			t.Logf("DST:          %q (full)", buffer)
			t.Logf("DST:          %q", buffer[n:])
			t.Logf("DST EXPECTED: %q", test.Expected)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}

		if expected, actual := test.Expected, buffer[:n]; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, what was written is not what was expected.", testNumber)
			t.Logf("SRC:          %q", test.Src)
			t.Logf("DST:          %q (full)", buffer)
			t.Logf("DST:          %q", buffer[n:])
			t.Logf("DST EXPECTED: %q", test.Expected)
			t.Logf("EXPECTED: %q", expected,)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}

func TestWriterMultipleWrites(t *testing.T) {

	tests := []struct {
		BufferLength int
		Srces [][]byte
		Expected []byte
	}{
		{
			BufferLength: 0,
			Srces: [][]byte{
				[]byte("a"),
				[]byte("p"),
				[]byte("p"),
				[]byte("l"),
				[]byte("e"),
			},
			Expected: []byte{},
		},
		{
			BufferLength: 1,
			Srces: [][]byte{
				[]byte("a"),
				[]byte("p"),
				[]byte("p"),
				[]byte("l"),
				[]byte("e"),
			},
			Expected: []byte("a"),
		},
		{
			BufferLength: 2,
			Srces: [][]byte{
				[]byte("a"),
				[]byte("p"),
				[]byte("p"),
				[]byte("l"),
				[]byte("e"),
			},
			Expected: []byte("ap"),
		},
		{
			BufferLength: 3,
			Srces: [][]byte{
				[]byte("a"),
				[]byte("p"),
				[]byte("p"),
				[]byte("l"),
				[]byte("e"),
			},
			Expected: []byte("app"),
		},
		{
			BufferLength: 4,
			Srces: [][]byte{
				[]byte("a"),
				[]byte("p"),
				[]byte("p"),
				[]byte("l"),
				[]byte("e"),
			},
			Expected: []byte("appl"),
		},
		{
			BufferLength: 5,
			Srces: [][]byte{
				[]byte("a"),
				[]byte("p"),
				[]byte("p"),
				[]byte("l"),
				[]byte("e"),
			},
			Expected: []byte("apple"),
		},
		{
			BufferLength: 6,
			Srces: [][]byte{
				[]byte("a"),
				[]byte("p"),
				[]byte("p"),
				[]byte("l"),
				[]byte("e"),
			},
			Expected: []byte("apple"),
		},
		{
			BufferLength: 7,
			Srces: [][]byte{
				[]byte("a"),
				[]byte("p"),
				[]byte("p"),
				[]byte("l"),
				[]byte("e"),
			},
			Expected: []byte("apple"),
		},
		{
			BufferLength: 8,
			Srces: [][]byte{
				[]byte("a"),
				[]byte("p"),
				[]byte("p"),
				[]byte("l"),
				[]byte("e"),
			},
			Expected: []byte("apple"),
		},
		{
			BufferLength: 9,
			Srces: [][]byte{
				[]byte("a"),
				[]byte("p"),
				[]byte("p"),
				[]byte("l"),
				[]byte("e"),
			},
			Expected: []byte("apple"),
		},
		{
			BufferLength: 10,
			Srces: [][]byte{
				[]byte("a"),
				[]byte("p"),
				[]byte("p"),
				[]byte("l"),
				[]byte("e"),
			},
			Expected: []byte("apple"),
		},



		{
			BufferLength: 0,
			Srces: [][]byte{
				[]byte("BA"),
				[]byte("NA"),
				[]byte("N"),
				[]byte("A"),
			},
			Expected: []byte{},
		},
		{
			BufferLength: 1,
			Srces: [][]byte{
				[]byte("BA"),
				[]byte("NA"),
				[]byte("N"),
				[]byte("A"),
			},
			Expected: []byte{},
		},
		{
			BufferLength: 2,
			Srces: [][]byte{
				[]byte("BA"),
				[]byte("NA"),
				[]byte("N"),
				[]byte("A"),
			},
			Expected: []byte("BA"),
		},
		{
			BufferLength: 3,
			Srces: [][]byte{
				[]byte("BA"),
				[]byte("NA"),
				[]byte("N"),
				[]byte("A"),
			},
			Expected: []byte("BA"),
		},
		{
			BufferLength: 4,
			Srces: [][]byte{
				[]byte("BA"),
				[]byte("NA"),
				[]byte("N"),
				[]byte("A"),
			},
			Expected: []byte("BANA"),
		},
		{
			BufferLength: 5,
			Srces: [][]byte{
				[]byte("BA"),
				[]byte("NA"),
				[]byte("N"),
				[]byte("A"),
			},
			Expected: []byte("BANAN"),
		},
		{
			BufferLength: 6,
			Srces: [][]byte{
				[]byte("BA"),
				[]byte("NA"),
				[]byte("N"),
				[]byte("A"),
			},
			Expected: []byte("BANANA"),
		},
		{
			BufferLength: 7,
			Srces: [][]byte{
				[]byte("BA"),
				[]byte("NA"),
				[]byte("N"),
				[]byte("A"),
			},
			Expected: []byte("BANANA"),
		},
		{
			BufferLength: 8,
			Srces: [][]byte{
				[]byte("BA"),
				[]byte("NA"),
				[]byte("N"),
				[]byte("A"),
			},
			Expected: []byte("BANANA"),
		},
		{
			BufferLength: 9,
			Srces: [][]byte{
				[]byte("BA"),
				[]byte("NA"),
				[]byte("N"),
				[]byte("A"),
			},
			Expected: []byte("BANANA"),
		},
		{
			BufferLength: 10,
			Srces: [][]byte{
				[]byte("BA"),
				[]byte("NA"),
				[]byte("N"),
				[]byte("A"),
			},
			Expected: []byte("BANANA"),
		},



		{
			BufferLength: 0,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte{},
		},
		{
			BufferLength: 1,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte{},
		},
		{
			BufferLength: 2,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte{},
		},
		{
			BufferLength: 3,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte{},
		},
		{
			BufferLength: 4,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte{},
		},
		{
			BufferLength: 5,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte{},
		},
		{
			BufferLength: 6,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple "),
		},
		{
			BufferLength: 7,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple "),
		},
		{
			BufferLength: 8,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple "),
		},
		{
			BufferLength: 9,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple "),
		},
		{
			BufferLength: 10,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple "),
		},
		{
			BufferLength: 11,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple "),
		},
		{
			BufferLength: 12,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple BANANA"),
		},
		{
			BufferLength: 13,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple BANANA "),
		},
		{
			BufferLength: 14,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple BANANA "),
		},
		{
			BufferLength: 15,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple BANANA "),
		},
		{
			BufferLength: 16,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple BANANA "),
		},
		{
			BufferLength: 17,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple BANANA "),
		},
		{
			BufferLength: 18,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple BANANA "),
		},
		{
			BufferLength: 19,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple BANANA Cherry"),
		},
		{
			BufferLength: 20,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple BANANA Cherry"),
		},
		{
			BufferLength: 21,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple BANANA Cherry"),
		},
		{
			BufferLength: 22,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple BANANA Cherry"),
		},
		{
			BufferLength: 23,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple BANANA Cherry"),
		},
		{
			BufferLength: 24,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple BANANA Cherry"),
		},
		{
			BufferLength: 25,
			Srces: [][]byte{
				[]byte("apple "),
				[]byte("BANANA"),
				[]byte(" "),
				[]byte("Cherry"),
			},
			Expected: []byte("apple BANANA Cherry"),
		},
	}

	TestLoop: for testNumber, test := range tests {

		var buffer []byte = make([]byte, test.BufferLength)

		var dst *buffers.Writer = buffers.NewWriter(buffer)

		var sum int

		InnerLoop: for srcNumber, src := range test.Srces {

			n, err := dst.Write(src)
			if _, casted := err.(buffers.TooShort); casted {
				break InnerLoop
			}
			if nil != err {
				t.Errorf("For test #%d and src #%d, did not expect an error, but actually got one.", testNumber, srcNumber)
				t.Logf("BUFFER LENGTH: %d", test.BufferLength)
				for i, s := range test.Srces {
					t.Logf("SRC[%d]: %q", i, s)
				}
				t.Logf("SRC:          %q", src)
				t.Logf("DST:          %q (full)", buffer)
				t.Logf("DST EXPECTED: %q", test.Expected)
				t.Logf("ERROR TYPE: %T", err)
				t.Logf("ERROR: %q", err)
				continue TestLoop
			}

			var expectedWritten int
			{
				lenBuffer := len(buffer)
				lenSrc    := len(src)

				if lenBuffer-sum < lenSrc {
					expectedWritten = lenBuffer-sum
				} else {
					expectedWritten = lenSrc
				}
			}

			if expected, actual := expectedWritten, n; expected != actual {
				t.Errorf("For test #%d and src #%d, the actual number of bytes written was not what was expected.", testNumber, srcNumber)
				t.Logf("BUFFER LENGTH: %d", test.BufferLength)
				for i, s := range test.Srces {
					t.Logf("SRC[%d]: %q", i, s)
				}
				t.Logf("SRC:          %q", src)
				t.Logf("DST:          %q (full)", buffer)
				t.Logf("DST:          %q", buffer[n:])
				t.Logf("DST EXPECTED: %q", test.Expected)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue TestLoop
			}

			sum += n
		}

		if expected, actual := test.Expected, buffer[:sum]; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, what was written is not what was expected.", testNumber)
			t.Logf("BUFFER LENGTH: %d", test.BufferLength)
			for i, s := range test.Srces {
				t.Logf("SRC[%d]: %q", i, s)
			}
			t.Logf("DST:          %q (full)", buffer)
			t.Logf("DST:          %q", buffer[sum:])
			t.Logf("DST EXPECTED: %q", test.Expected)
			t.Logf("EXPECTED: %q", expected,)
			t.Logf("ACTUAL:   %q", actual)
			continue TestLoop
		}
	}
}
