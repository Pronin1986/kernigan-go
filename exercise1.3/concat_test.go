package main

import "testing"

func BenchmarkSimpleConcat(b *testing.B) {
	strSlice := randStrings(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SimpleConcat(strSlice)
	}
}

func BenchmarkStringJoinConcat(b *testing.B) {
	strSlice := randStrings(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = StringJoinConcat(strSlice)
	}
}

func BenchmarkStringBuilderConcat(b *testing.B) {
	strSlice := randStrings(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = StringBuilderConcat(strSlice)
	}
}

func BenchmarkStringBuilderWithAllocConcat(b *testing.B) {
	strSlice := randStrings(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = StringBuilderWithAllocConcat(strSlice)
	}
}
