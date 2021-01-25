package main

import (
	"testing"
)

func BenchmarkStandardGetRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardGetRequest()
	}
}

func BenchmarkSocketGetRequest(b *testing.B) {
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		SocketGetRequest()
	}

	b.StopTimer()
}
