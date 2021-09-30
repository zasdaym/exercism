// Package lasagna provides Lasagna Master solution.
package lasagna

// PreparationTime returns the preparation time for the given layers.
func PreparationTime(layers []string, duration int) int {
	if duration == 0 {
		duration = 2
	}
	return len(layers) * duration
}

// Quantities returns the quantities of noodles and and sauce needed to cook the given layers.
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, l := range layers {
		switch l {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		default:
		}
	}
	return noodles, sauce
}

// AddSecretIngredient adds the last ingredient in secret into original.
func AddSecretIngredient(secret, original []string) []string {
	return append(original, secret[len(secret)-1])
}

// ScaleRecipe scales the given quantities for given portions.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaled []float64
	for _, q := range quantities {
		scaled = append(scaled, q/2*float64(portions))
	}
	return scaled
}
