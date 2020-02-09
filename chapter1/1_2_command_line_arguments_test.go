package main

import "testing"

//TestCla test function
func TestCla(t *testing.T) {
	got := Cla()
	want := " /var/folders/92/2gz7fm7x3n1d8c9302rh1b100000gn/T/go-build843363716/b001/chapter1.test -test.timeout=30s -test.timeout=30s -test.run=^(TestCla)$ "

	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}

//BenchmarkCla runs a benchmark on Cla() function
func BenchmarkCla(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cla()
	}
}
