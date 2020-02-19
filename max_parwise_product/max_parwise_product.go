package max_parwise_product

import "sort"

// MaxParwiseProduct returns the highest product of 2 items in the given list (n).
func MaxParwiseProduct(n []int) int {
	sort.Sort(sort.IntSlice(n))
	return n[len(n)-1] * n[len(n)-2]
}
