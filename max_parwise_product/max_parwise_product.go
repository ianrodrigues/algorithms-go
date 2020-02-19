package max_parwise_product

func MaxParwiseProduct(n []int) int {
	var result int

	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			p := n[i] * n[j]
			if p > result {
				result = p
			}
		}
	}

	return result
}
