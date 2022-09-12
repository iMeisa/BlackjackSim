package models

import "BlackjackSim/cards"

type Hand struct {
	Cards  []cards.Card
	UpCard cards.Card // Dealer only
	Soft   bool
}

func (h *Hand) Calculate() int {
	var totalValue int
	var aces = 0
	for _, card := range h.Cards {
		switch card {
		case cards.Ace:
			aces++
		default:
			totalValue += card.Value
		}
	}

	// Aces calculation
	totalValue += aces
	h.Soft = false
	if aces > 0 && totalValue+10 <= 21 {
		totalValue += 10
		h.Soft = true
	}

	return totalValue
}

type Player struct {
	Hands []Hand
}
