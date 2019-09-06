/*
Package buffers provides tools for working with byte array, and byte slice buffers, for the Go programming language.

Example

Here is an example of buffers.Writer being used to provide a io.Writer interface to an byte array, and byte slice:

	import "github.com/reiver/go-buffers"
	
	// ...
	
	var buffer [1024]byte
	
	var p []byte = buffer[:]
	
	writer := buffers.NewWriter(p)
	
	n, err := writer.Write(data)
	
	switch casted := err.(type) {
	case buffers.TooShort:
		//@TODO
	default:
		return err
	}
*/
package buffers
