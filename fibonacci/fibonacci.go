package fibonacci

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}

	f := [2]int{0, 1}

	for i := 2; i <= n; i++ {
		f[0], f[1] = f[1], f[0]+f[1]
	}

	return f[1]
}
