# go-buffers

Package **buffers** provides tools for working with byte array, and byte slice buffers, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-buffers

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-buffers?status.svg)](https://godoc.org/sourcecode.social/reiver/go-buffers)

## Example
```go
import "sourcecode.social/reiver/go-buffers"

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

```
