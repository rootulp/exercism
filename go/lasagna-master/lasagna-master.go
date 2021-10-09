package lasagna

func PreparationTime(layers []string, avgPrepTimePerLayer int) int {
	const DEFAULT_PREP_TIME_PER_LAYER = 2

	if avgPrepTimePerLayer == 0 {
		return len(layers) * DEFAULT_PREP_TIME_PER_LAYER
	}
	return len(layers) * avgPrepTimePerLayer
}

func Quantities(layers []string) (quantityNoodles int, quantitySauce float64) {
	numNoodles := count(layers, "noodles")
	numSauce := count(layers, "sauce")
	return numNoodles * 50, float64(numSauce) * 0.2
}

// count returns the number of times item occurs in s
func count(s []string, item string) (result int) {
	for _, v := range s {
		if v == item {
			result += 1
		}
	}
	return result
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	lastItem := friendsList[len(friendsList)-1]
	result := append(myList, lastItem)
	return result
}

func ScaleRecipe(quantities []float64, desiredNumPortions int) []float64 {
	result := make([]float64, len(quantities))
	for i, quantity := range quantities {
		quantityForOnePortion := quantity / 2.0
		result[i] = quantityForOnePortion * float64(desiredNumPortions)
	}
	return result
}
