package buffers_test

import (
	"testing"

	"sourcecode.social/reiver/go-buffers"
)

func TestLimitedBuffer_oneByteAtATime(t *testing.T) {

	tests := []struct{
		Data string
		Expected string
		Limit int
	}{
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "",
			Limit: 0,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0",
			Limit: 1,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "01",
			Limit: 2,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "012",
			Limit: 3,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123",
			Limit: 4,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "01234",
			Limit: 5,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "012345",
			Limit: 6,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456",
			Limit: 7,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "01234567",
			Limit: 8,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "012345678",
			Limit: 9,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789",
			Limit: 10,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"A",
			Limit: 11,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"AB",
			Limit: 12,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABC",
			Limit: 13,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCD",
			Limit: 14,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDE",
			Limit: 15,
		},



		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXY",
			Limit: 35,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Limit: 36,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"a",
			Limit: 37,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"ab",
			Limit: 38,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abc",
			Limit: 39,
		},



		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvw",
			Limit: 59,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwx",
			Limit: 60,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxy",
			Limit: 61,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Limit: 62,
		},
	}

	loop: for testNumber, test := range tests {

		{
			var buffer buffers.Buffer = buffers.LimitedBuffer(test.Limit)

			var i int
			for i=0; i<test.Limit; i++ {
				var b [1]byte
				var p  []byte = b[:]

				b[0] = test.Data[i]

				n, err := buffer.Write(p)
				if nil != err {
					t.Errorf("For test #%d, did not expect an error but actually got one,", testNumber)
					t.Logf("ERROR: (%T) %q", err, err)
					t.Logf("DATA:          %q", test.Data)
					t.Logf("EXPECTED-DATA: %q", test.Expected)
					t.Logf("ACTUAL-DATA:   %q", buffer.String())
					t.Logf("LIMIT: %d", test.Limit)
					continue loop
				}
				if expected, actual := 1, n; expected != actual {
					t.Errorf("For test #%d, the actual number of bytes written is not what was expected.", testNumber)
					t.Logf("EXPECTED: %d", expected)
					t.Logf("ACTUAL:   %d", actual)
					t.Logf("DATA:          %q", test.Data)
					t.Logf("EXPECTED-DATA: %q", test.Expected)
					t.Logf("ACTUAL-DATA:   %q", buffer.String())
					continue loop
				}
			}
			{
				var b [1]byte
				var p  []byte = b[:]

				if len(test.Data) < i {
					b[0] = test.Data[i]
				} else {
					b[0] = '?'
				}
				n, err := buffer.Write(p)
				if nil == err {
					t.Errorf("For test #%d, expected an error but did not actually get one,", testNumber)
					t.Logf("ERROR: (%T) %q", err, err)
					t.Logf("DATA:          %q", test.Data)
					t.Logf("EXPECTED-DATA: %q", test.Expected)
					t.Logf("ACTUAL-DATA:   %q", buffer.String())
					t.Logf("LIMIT: %d", test.Limit)
					continue loop
				}
				if buffers.BufferOverflow != err {
					t.Errorf("For test #%d, expected a buffer-overflow error but actually got a different error.", testNumber)
					t.Logf("ERROR: (%T) %q", err, err)
					t.Logf("DATA:          %q", test.Data)
					t.Logf("EXPECTED-DATA: %q", test.Expected)
					t.Logf("ACTUAL-DATA:   %q", buffer.String())
					t.Logf("LIMIT: %d", test.Limit)
					continue loop
				}
				if expected, actual := 0, n; expected != actual {
					t.Errorf("For test #%d, the actual number of bytes written is not what was expected.", testNumber)
					t.Logf("EXPECTED: %d", expected)
					t.Logf("ACTUAL:   %d", actual)
					t.Logf("DATA:     %q", test.Data)
					t.Logf("EXPECTED-DATA: %q", test.Expected)
					t.Logf("ACTUAL-DATA:   %q", buffer.String())
					continue loop
				}
			}
			{
				var expected string = test.Expected
				var actual   string = buffer.String()

				if expected != actual {
					t.Errorf("For test #%d, the actual value in the buffer is not what was expected.", testNumber)
					t.Logf("EXPECTED: %q", expected)
					t.Logf("ACTUAL:   %q", actual)
					t.Logf("DATA: %q", test.Data)
					t.Logf("LIMIT: %d", test.Limit)
					continue loop
				}
			}
		}
	}
}

func TestLimitedBuffer_all(t *testing.T) {

	tests := []struct{
		Data string
		Expected string
		Limit int
	}{
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "",
			Limit: 0,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0",
			Limit: 1,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "01",
			Limit: 2,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "012",
			Limit: 3,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123",
			Limit: 4,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "01234",
			Limit: 5,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "012345",
			Limit: 6,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456",
			Limit: 7,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "01234567",
			Limit: 8,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "012345678",
			Limit: 9,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789",
			Limit: 10,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"A",
			Limit: 11,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"AB",
			Limit: 12,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABC",
			Limit: 13,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCD",
			Limit: 14,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDE",
			Limit: 15,
		},



		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXY",
			Limit: 35,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Limit: 36,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"a",
			Limit: 37,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"ab",
			Limit: 38,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abc",
			Limit: 39,
		},



		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvw",
			Limit: 59,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwx",
			Limit: 60,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxy",
			Limit: 61,
		},
		{
			Data:     "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Expected: "0123456789"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"abcdefghijklmnopqrstuvwxyz",
			Limit: 62,
		},
	}

	loop: for testNumber, test := range tests {

		var limit int = len(test.Expected)

		var buffer buffers.Buffer = buffers.LimitedBuffer(limit)

		n, err := buffer.Write([]byte(test.Expected))
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one,", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
			t.Logf("DATA:          %q", test.Data)
			t.Logf("EXPECTED-DATA: %q", test.Expected)
			t.Logf("ACTUAL-DATA:   %q", buffer.String())
			t.Logf("LIMIT: %d", test.Limit)
			continue loop
		}
		if expected, actual := limit, n; expected != actual {
			t.Errorf("For test #%d, the actual number of bytes written is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("DATA:          %q", test.Data)
			t.Logf("EXPECTED-DATA: %q", test.Expected)
			continue loop
		}
		{
			var expected string = test.Expected
			var actual   string = buffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual value in the buffer is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("DATA: %q", test.Data)
				t.Logf("LIMIT: %d", test.Limit)
				continue loop
			}
		}
	}
}
