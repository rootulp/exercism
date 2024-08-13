package knapsack

type Item struct {
	Weight, Value int
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maximumWeight int, items []Item) (total int) {
	currentValue := 0
	currentWeight := 0
	return knapsack(items, currentValue, currentWeight, maximumWeight)
}

func knapsack(items []Item, value int, weight int, maxWeight int) (total int) {
	if len(items) == 0 {
		return value
	}
	item := items[0]
	totalWithItem := knapsack(items[1:], value+item.Value, weight+item.Weight, maxWeight)
	totalWithoutItem := knapsack(items[1:], value, weight, maxWeight)

	// If the current item is too heavy, return the total without the current item
	if weight+item.Weight > maxWeight {
		return totalWithoutItem
	}
	return max(totalWithItem, totalWithoutItem)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
