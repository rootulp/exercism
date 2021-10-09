package lasagna

func PreparationTime(layers []string, avgPrepTimePerLayer int) int {
	const DEFAULT_PREP_TIME_PER_LAYER = 2

	if avgPrepTimePerLayer == 0 {
		return len(layers) * DEFAULT_PREP_TIME_PER_LAYER
	}
	return len(layers) * avgPrepTimePerLayer
}

// TODO: define the 'Quantities()' function

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function
