package models

type Hand struct {
	Cards []Card
}

func (h *Hand) Calculate() int {
	var totalValue int
	var aces = 0
	for _, card := range h.Cards {
		switch card.Name {
		case "A":
			aces++
		default:
			totalValue += card.Value
		}
	}

	// Aces calculation
	totalValue += aces
	if aces > 0 && totalValue+10 <= 21 {
		totalValue += 10
	}

	return totalValue
}

type Player struct {
	Hands []Hand
}
