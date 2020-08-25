// Package allergies is solution for problem Allergies.
package allergies

var (
	allergens = []string{
		"eggs",
		"peanuts",
		"shellfish",
		"strawberries",
		"tomatoes",
		"chocolate",
		"pollen",
		"cats",
	}

	allergenCodes = map[string]byte{
		"eggs":         1,
		"peanuts":      2,
		"shellfish":    4,
		"strawberries": 8,
		"tomatoes":     16,
		"chocolate":    32,
		"pollen":       64,
		"cats":         128,
	}
)

// Allergies returns all positive allergens of given allergy score.
func Allergies(n uint) []string {
	result := make([]string, 0, 8)
	code := byte(n)

	for i := 0; code > 0; i++ {
		if code&1 == 1 {
			result = append(result, allergens[i])
		}
		code >>= 1
	}
	return result
}

// AllergicTo checks if given score is positive to given allergen.
func AllergicTo(n uint, allergen string) bool {
	return byte(n)&allergenCodes[allergen] > 0
}
