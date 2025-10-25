package main

import (
	"sync/atomic"
	"testing"
)

// run go test --race
// return error
// cause Data Race
func TestBad(t *testing.T) {
	var counter int32

	for i := 0; i < 10; i++ {
		go func(i int) {
			counter += int32(i)
		}(i)
	}
}

// Goes fine
func TestGood(t *testing.T) {
	var counter int32

	for i := 0; i < 10; i++ {
		go func(i int) {
			atomic.AddInt32(&counter, int32(i))
		}(i)
	}
}
