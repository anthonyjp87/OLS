package ols

import (
	"testing"
)

func TestOls(t *testing.T) {

	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	y := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	b1, b0, r2 := reg(x, y)
	if b1 != 1 {
		t.Errorf("Error b1")
	}
	if b0 != 0 {
		t.Errorf("Error b0")
	}
	if r2 != 1 {
		t.Errorf("Error R2")
	}

}

func TestMean(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	if mean(x) != 3 {
		t.Error("Mean func is wrong")
	}

}

func TestSum(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	if sum(x) != 15 {
		t.Error("Sum is wrong")
	}
}

func TestDif(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	difx := dif(x)
	xexp := []float64{-2, -1, 0, 1, 2}

	for i := range difx {
		if difx[i] != xexp[i] {
			t.Error("Diff func is wrong")
		}
	}
}

func TestMultiply(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{1, 2, 3, 4, 5}
	imulti := multiply(x, y)
	expected := []float64{1, 4, 9, 16, 25}

	for i := range imulti {
		if imulti[i] != expected[i] {
			t.Error("Multi func is wrong")
		}
	}
}

func TestSquare(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	square := square(x)
	expected := []float64{1, 4, 9, 16, 25}

	for i := range square {
		if square[i] != expected[i] {
			t.Error("Square func is wrong")
		}
	}
}

//BENCHMARKS

func Benchmark(b *testing.B) {
	for n := 0; n < b.N; n++ {
		x := []float64{1, 2, 3, 4, 5}
		y := []float64{1, 4, 5, 4, 5}
		reg(x, y)

	}
}

func BenchmarkMean(b *testing.B) {

	for n := 0; n < b.N; n++ {
		x := []float64{1, 2, 3, 4, 5}
		mean(x)

	}
}
func BenchmarkSum(b *testing.B) {

	for n := 0; n < b.N; n++ {
		x := []float64{1, 2, 3, 4, 5}
		sum(x)

	}
}
func BenchmarkSquare(b *testing.B) {

	for n := 0; n < b.N; n++ {
		x := []float64{1, 2, 3, 4, 5}
		square(x)

	}
}
func BenchmarkDif(b *testing.B) {

	for n := 0; n < b.N; n++ {
		x := []float64{1, 2, 3, 4, 5}
		dif(x)

	}
}

func BenchmarkMultiply(b *testing.B) {

	for n := 0; n < b.N; n++ {
		x := []float64{1, 2, 3, 4, 5}
		y := []float64{1, 4, 5, 4, 5}
		multiply(x, y)

	}
}

//200k n 1.2
