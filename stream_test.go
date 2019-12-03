// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go test -bench . -benchmem -memprofile p.out -gcflags "-newescape=false -m=2"

// Tests to see how each algorithm compare.
package main

import (
	"bytes"
	"testing"
)

// Capture the time it takes to execute algorithm one.
func BenchmarkMethod1(b *testing.B) {
	var output bytes.Buffer
	find := []byte("sunny")
	repl := []byte("Sunny")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		method1(rawData, find, repl, &output)
	}
}

// Capture the time it takes to execute algorithm two.
func BenchmarkMethod2(b *testing.B) {
	var output bytes.Buffer
	find := []byte("sunny")
	repl := []byte("Sunny")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		method2(rawData, find, repl, &output)
	}
}

// Capture the time it takes to execute algorithm two.
func BenchmarkMethod3(b *testing.B) {
	var output bytes.Buffer
	find := []byte("sunny")
	repl := []byte("Sunny")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		method3(rawData, find, repl, &output)
	}
}