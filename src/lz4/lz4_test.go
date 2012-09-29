package lz4

import (
	"fmt"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	out := make([]byte, CompressBound(len(alice)))
	in := []byte(alice)
	start := time.Now()
	n := Compress(in, out)
	end := time.Now()
	out = out[:n]
	elapsed := end.Sub(start).Seconds()
	fmt.Println("Compressed", len(alice), "to", len(out), ": ", float64(len(alice))/float64(len(out)),
		"in", elapsed, "s @", float64(len(alice))/elapsed, "/s")
}

func TestBasicHC(t *testing.T) {
	out := make([]byte, len(alice))
	in := []byte(alice)
	start := time.Now()
	n := CompressHC(in, out)
	end := time.Now()
	out = out[:n]
	elapsed := end.Sub(start).Seconds()
	fmt.Println("Compressed", len(alice), "to", len(out), ": ", float64(len(alice))/float64(len(out)),
		"in", elapsed, "s @", float64(len(alice))/elapsed, "/s")

	in = out
	out = make([]byte, len(alice))
	start = time.Now()
	n = DecompressUnknownOutputSize(in, out)
	end = time.Now()

	if n != len(alice) {
		t.Error("expected ", len(alice), "got", n)
	}

	if string(out) != alice {
		t.Error("Decompression did not give identical results")
	}

	out = make([]byte, len(alice))
	start = time.Now()
	// we know upfront what the output size is meant to be.
	err := Decompress(in, out)
	end = time.Now()

	if err != nil {
		t.Error(err)
	}

	if string(out) != alice {
		t.Error("Decompression did not give identical results")
	}

	elapsed = end.Sub(start).Seconds()
	fmt.Println("Decompressed", len(out), "to", len(in),
		"in", elapsed, "s @", float64(len(alice))/elapsed, "/s")
}

var compressedAlice []byte

func init() {
	out := make([]byte, CompressBound(len(alice)))
	in := []byte(alice)
	n := Compress(in, out)
	compressedAlice = out[:n]
}

func BenchmarkBasic(b *testing.B) {
	out := make([]byte, CompressBound(len(alice)))
	in := []byte(alice)

	for ii := 0; ii < b.N; ii++ {
		Compress(in, out)
	}
}

func BenchmarkBasicHC(b *testing.B) {
	out := make([]byte, CompressBound(len(alice)))
	in := []byte(alice)

	for ii := 0; ii < b.N; ii++ {
		CompressHC(in, out)
	}
}

func BenchmarkDecompressU(b *testing.B) {
	in := compressedAlice
	out := make([]byte, len(alice))
	for ii := 0; ii < b.N; ii++ {
		DecompressUnknownOutputSize(in, out)
	}
}

func BenchmarkDecompress(b *testing.B) {
	in := compressedAlice
	out := make([]byte, len(alice))
	for ii := 0; ii < b.N; ii++ {
		err := Decompress(in, out)
		if err != nil {
			b.Error(err)
		}
	}
}
