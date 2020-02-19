package max_parwise_product

import "testing"

func TestMaxParwiseProduct6(t *testing.T) {
	result := MaxParwiseProduct([]int{1, 2, 3})
	if result != 6 {
		t.Errorf("Failed asserting MaxParwiseProduct([]int{1, 2, 3}) is 6; got %d", result)
	}
}

func TestMaxParwiseProduct56(t *testing.T) {
	result := MaxParwiseProduct([]int{7, 2, 3, 8})
	if result != 56 {
		t.Errorf("Failed asserting MaxParwiseProduct([]int{7, 2, 3, 8}) is 56; got %d", result)
	}
}

func TestMaxParwiseProduct90(t *testing.T) {
	result := MaxParwiseProduct([]int{7, 9, 10, 2, 5, 3, 8})
	if result != 90 {
		t.Errorf("Failed asserting MaxParwiseProduct([]int{7, 9, 10, 2, 5, 3, 8}) is 90; got %d", result)
	}
}
