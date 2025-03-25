package binarysearch

func SearchInts(list []int, key int) int {
	n := len(list)
	l, r := 0, n
	for l < r {
		m := (l + r) >> 1
		if list[m] == key {
			return m
		}
		if list[m] < key {
			l = m + 1
		} else {
			r = m
		}
	}
	return -1
}
