package fibonacci

import "testing"

func TestFibonacci0(t *testing.T) {
	result := Fibonacci(0)
	if result != 0 {
		t.Errorf("Failed asserting Fibonacci(0) is 0; got %d", result)
	}
}

func TestFibonacci1(t *testing.T) {
	result := Fibonacci(1)
	if result != 1 {
		t.Errorf("Failed asserting Fibonacci(1) is 1; got %d", result)
	}
}

func TestFibonacci2(t *testing.T) {
	result := Fibonacci(2)
	if result != 1 {
		t.Errorf("Failed asserting Fibonacci(2) is 1; got %d", result)
	}
}

func TestFibonacci10(t *testing.T) {
	result := Fibonacci(10)
	if result != 55 {
		t.Errorf("Failed asserting Fibonacci(10) is 55; got %d", result)
	}
}
