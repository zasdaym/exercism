package knapsack

type Item struct {
	Weight, Value int
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maxWeight
func Knapsack(maxWeight int, items []Item) int {
	cache := make([]int, maxWeight+1)
	for _, item := range items {
		for w := maxWeight; w >= item.Weight; w-- {
			cache[w] = max(cache[w], item.Value+cache[w-item.Weight])
		}
	}
	return cache[maxWeight]
}
