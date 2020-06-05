package main

import (
	"fmt"
	"reflect"

	"golang.org/x/tour/reader"
)

type MyReader struct {
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (m MyReader) Read(b []byte) (n int, err error) {
	n = copy(b, "A")
	return

}

func main() {
	fmt.Println(reflect.TypeOf(MyReader{}))
	reader.Validate(MyReader{})
}
