package maths

// Factorial returns the factorial of n using recursion
func Factorial(n uint8) uint64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 1
	}

	return uint64(n) * Factorial(n-1)
}

// LoopFactorial returns the factorial of n using a for loop.
func LoopFactorial(n uint8) uint64 {
	if n == 0 {
		return 1
	}

	v := uint64(1)
	for i := uint8(1); i <= n; i++ {
		v *= uint64(i)
	}

	return v
}
