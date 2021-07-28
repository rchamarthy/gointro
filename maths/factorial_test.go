package maths_test

import (
	"testing"

	"github.com/rchamarthy/gointro/maths"
	"github.com/stretchr/testify/assert"
)

func factorialTable() map[uint8]uint64 {
	return map[uint8]uint64{
		0:  1, // Boundary value
		1:  1,
		2:  2,
		3:  6,
		4:  24,
		5:  120,
		6:  720,
		7:  5040,
		8:  40320,
		9:  362880,
		10: 3628800,
		11: 39916800,
		12: 479001600,
		13: 6227020800,
		14: 87178291200,
		15: 1307674368000,
		16: 20922789888000,
		17: 355687428096000,
		18: 6402373705728000,
		19: 121645100408832000,
		20: 2432902008176640000,
	}
}

func TestFactorial(t *testing.T) {
	assert := assert.New(t)
	for n, v := range factorialTable() {
		assert.Equal(maths.Factorial(n), v)
	}
}

func TestLoopFactorial(t *testing.T) {
	assert := assert.New(t)
	for n, v := range factorialTable() {
		assert.Equal(maths.LoopFactorial(n), v)
	}
}

func BenchmarkFactorial(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < 21; i++ {
			maths.Factorial(uint8(i))
		}
	}
}

func BenchmarkLoopFactorial(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < 21; i++ {
			maths.LoopFactorial(uint8(i))
		}
	}
}
