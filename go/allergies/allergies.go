// Package allergies provides Allergies solution.
package allergies

// scores contains allergy score of each allergen.
var scores = map[string]int{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

// Allergies returns list allergens from a person's allergy score.
func Allergies(score uint) []string {
	var allergens []string
	for allergen := range scores {
		if AllergicTo(score, allergen) {
			allergens = append(allergens, allergen)
		}
	}
	return allergens
}

// AllergicTo determines whether a person is allergic to an allergen based on given allergy score.
func AllergicTo(score uint, allergen string) bool {
	return score&uint(scores[allergen]) != 0
}
