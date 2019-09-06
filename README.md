# go-buffers

Package **buffers** provides tools for working with byte array, and byte slice buffers.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-buffers

[![GoDoc](https://godoc.org/github.com/reiver/go-buffers?status.svg)](https://godoc.org/github.com/reiver/go-buffers)

## Example
```go
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

```
