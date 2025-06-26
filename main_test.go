package main

import "testing"

func escape() *int {
	x := 42
	return &x
}

func noEscape() int {
	x := 42
	return x
}

func BenchmarkEscape(b *testing.B) {
	for b.Loop() {
		_ = escape()
	}
}

func BenchmarkNoEscape(b *testing.B) {
	for b.Loop() {
		_ = noEscape()
	}
}
