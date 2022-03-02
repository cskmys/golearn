// Benchmark tests are a subset/particular type of test
// hence you will use standard library's testing suite, and all the naming conventions apply here too.
// file naming, blackbox, white-box testing conventions etc are all same but there are two main differences
// We prefix function name with "Benchmark" instead of "Test" and the benchmark test functions accept "*testing.B" instead of "*testing.T"

// here we benchmark few hashing algorithms from standard library

package bmark

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func BenchmarkSHA1(b *testing.B) { // function name prefixed with "Benchmark" instead of "Test" and the benchmark test functions accept "*testing.B" instead of "*testing.T"
	// normally by default the benchmark timer starts at the beginning of the benchmark test execution and stops at the end of the benchmark test execution
	// additionally there are APIs to control the start and stop of the timer
	data := []byte("Mary had a little lamb") // setup code
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
	b.StopTimer()
	// clean-up code
}

func BenchmarkSHA256(b *testing.B) {
	data := []byte("Mary had a little lamb")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
	// no need to explicitly write code to stop timer coz at the end of the test timer stops automatically
}

func BenchmarkSHA512(b *testing.B) {
	data := []byte("Mary had a little lamb")
	// no need to explicitly write code to start timer coz at the beginning of the test timer starts automatically
	// however, here the timer starts before "data" is allocated, however the time taken for that is trivial, and we need not worry about its impact
	for i := 0; i < b.N; i++ {
		sha512.Sum512(data)
	}
}
