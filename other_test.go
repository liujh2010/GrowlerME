package growler

import (
	"fmt"
	"math/big"
	"testing"
)

func BenchmarkFloat(b *testing.B) {
	a, c, d := 0.001, 0.23452143242142, 1314.1231313
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = c * d
	}
	a = a / c
}

func BenchmarkBigFloat(b *testing.B) {
	// a, c, d := 0.001, 0.23452143242142, 1314.1231313
	c := big.NewFloat(0.23452143242142)
	d := big.NewFloat(1314.1231313)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Mul(c, d)
	}
}

func BenchmarkInt(b *testing.B) {
	a, c, d := 0, 34, 89
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = c * d
	}
	a = a * c
	f := 99999999994345.3
	g := 32424442424353.2
	f = f - g
	fmt.Println(f)
}

func BenchmarkBigInt(b *testing.B) {
	a, ok1 := new(big.Int).SetString("184467440737095516", 10)
	c, ok2 := new(big.Int).SetString("184467440737095516", 10)

	if !ok1 && !ok2 {
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Mul(a, c)
	}
}

func BenchmarkJudgement(b *testing.B) {
	a, c := 221312313, 34354353123
	var d bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d = a == c
	}
	if d {

	}
}

func TestUint64Div(t *testing.T) {
	// a := 1798467921
	a := 21
	b := 967434
	c := a / b
	d := a % b
	fmt.Println(c, d)
}
