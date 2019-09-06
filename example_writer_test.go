package buffers_test

import (
	"github.com/reiver/go-buffers"

	"fmt"
)

func ExampleWriter() {

	var buffer [256]byte

	var p []byte = buffer[:]

	writer := buffers.NewWriter(p)

	data := []byte("Hello world!")

	n, err := writer.Write(data)
	if nil != err {
		fmt.Printf("ERROR: Problem writing: %s\n", err)
		return
	}

	fmt.Printf("Wrote %d bytes to the buffer.\n", n)
	fmt.Printf("Those %d bytes in the buffer have a value of “%s”.\n", n, buffer)

	// Output:
	// Wrote 12 bytes to the buffer.

	// Those 12 bytes in the buffer have a value of “Hello world!”.
}
