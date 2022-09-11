package models

import (
	"math/rand"
	"time"
)

type Shoe struct {
	Index int
	Cards []Card
}

func NewShoe(decks int) *Shoe {
	newShoe := Shoe{}

	for _, cardType := range cards {
		for i := 0; i < 4*decks; i++ {
			newShoe.Cards = append(newShoe.Cards, cardType)
		}
	}

	newShoe.Shuffle()

	return &newShoe
}

func (s *Shoe) Shuffle() {
	a := s.Cards
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}
