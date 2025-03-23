package pascal

func Triangle(n int) [][]int {
	result := [][]int{{1}}
	for i := 1; i < n; i++ {
		prev := result[i-1]
		curr := make([]int, 0, i)
		curr = append(curr, 1)
		for j := 1; j < i; j++ {
			curr = append(curr, prev[j-1]+prev[j])
		}
		curr = append(curr, 1)
		result = append(result, curr)
	}
	return result
}
