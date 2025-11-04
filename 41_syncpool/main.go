package main

import (
	"bytes"
	"fmt"
	"sync"
)

// Работа с syncPool

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func processData(data []byte) string {

	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf)

	buf.Reset()
	buf.Write(data)

	return buf.String()
}

func main() {
	fmt.Println(processData([]byte("hello")))
}
