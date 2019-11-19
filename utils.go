package tree

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func longestCommonPrefix(a, b string) (i int) {
	max := min(len(a), len(b))

	for i < max && a[i] == b[i] {
		i++
	}
	return
}
