// Package blackjack provides Blackjack solution.
package blackjack

var cardValueByName = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cardValueByName[card]
}

// blackjack represents blackjack score.
const blackjack = 21

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1) + ParseCard(card2) == blackjack
}

// All available turn options.
const (
	stand = "S"	
	hit = "H"	
	split = "P"	
	autoWin = "W"	
)

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if !isBlackjack {
		return split
	}
	if dealerScore < 10 {
		return autoWin
	}
	return stand
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	switch {
	case handScore >= 17:
		return stand
	case handScore <= 11:
		return hit
	default:
		if dealerScore >= 7 {
			return hit
		}
		return stand
	}
}

// FirstTurn returns the semi-optimal decision for the first turn, given the cards of the player and the dealer.
// This function is already implemented and does not need to be edited. It pulls the other functions together in a
// complete decision tree for the first turn.
func FirstTurn(card1, card2, dealerCard string) string {
	handScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	if handScore > 20 {
		return LargeHand(IsBlackjack(card1, card2), dealerScore)
	}
	return SmallHand(handScore, dealerScore)
}
