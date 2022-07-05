package search

func LinearSearch(a []int, key int) int {
	for i, item := range a {
		if item == key {
			return i
		}
	}
	return -1
}
