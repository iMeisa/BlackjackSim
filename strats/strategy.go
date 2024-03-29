package strats

import (
	"BlackjackSim/cards"
	"BlackjackSim/models"
)

// GetBet returns the bet amount for each hand
// Modify this for bet strategies
func GetBet(minBet, maxBet, baseBet, trueCount int) int {

	if trueCount < 1 {
		return minBet
	}

	if trueCount > 1 {
		return min(baseBet*(trueCount-1), maxBet)
	}
	return baseBet
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Decide(upCard cards.Card, playerHand *models.Hand, hardTotal, softTotal, splitting map[int]map[int]string) Decision {

	// Check for blackjack already handled by main()

	// Return split decision if applicable
	if len(playerHand.Cards) == 2 && playerHand.Cards[0] == playerHand.Cards[1] {
		if sToD(splitting[playerHand.Cards[0].Value][upCard.Value]) == Split {
			return Split
		}
	}

	// Return hit/stand decision if soft
	if playerHand.Soft {
		return sToD(softTotal[playerHand.Calculate()][upCard.Value])
	}
	// Return hit/stand decision if hard
	return sToD(hardTotal[playerHand.Calculate()][upCard.Value])
}

func HouseDecision(houseHand *models.Hand) Decision {
	houseValue := houseHand.Calculate()
	if houseValue < 17 {
		return Hit
	}

	return Stand
}

// sToD converts a string to a Decision
func sToD(s string) Decision {
	switch s {
	case "H":
		return Hit
	case "S":
		return Stand
	case "D":
		return Double
	case "Y":
		return Split
	case "N":
		return noSplit
	default:
		return Stand
	}
}
