package ptests

import (
	"crypto/sha512"
	"testing"
)

func BenchmarkSHA512(b *testing.B) {
	data := []byte("Mary had a little lamb")
	for i := 0; i < b.N; i++ {
		sha512.Sum512(data)
	}
}

func BenchmarkSHA512Alloc(b *testing.B) {
	data := []byte("Mary had a little lamb")
	for i := 0; i < b.N; i++ {
		h := sha512.New() // due to the use of "New()" here some memory is allocated but nothing was allocated in "BenchmarkSHA512",
		// you should see this when you run "go test -bench 512 -benchmem",
		// in the command "512" is the regex to run both "BenchmarkSHA512" and "BenchmarkSHA512Alloc" tests
		h.Sum(data)
		// to get more info you can do "go test -bench Alloc -memprofile 512Alloc_profile.out"
		// here regex "Alloc" is going to make sure that only "BenchmarkSHA512Alloc" is run
		// you cannot manually read any info from "512Alloc_profile.out", you'll need to do: "go tool pprof 512Alloc_profile.out"
		// this will open an interface where you can type commands to perform actions on data in "512Alloc_profile.out"
		// now to get a visual output enter: "svg" to generate an ".svg" file, and to see properly open it in web browser
		// you'll see call stack, memory utilization information etc
	}
}
