package models

import (
	"BlackjackSim/cards"
	"math/rand"
	"time"
)

type Shoe struct {
	Index int
	Cards []cards.Card
}

func NewShoe(decks int) *Shoe {
	newShoe := Shoe{}

	for _, cardType := range cards.Cards {
		for i := 0; i < 4*decks; i++ {
			newShoe.Cards = append(newShoe.Cards, cardType)
		}
	}

	newShoe.Shuffle()

	return &newShoe
}

func (s *Shoe) NextCard() cards.Card {
	card := s.Cards[s.Index]
	s.Index++
	return card
}

func (s *Shoe) Shuffle() {
	a := s.Cards
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	s.Index = 0
}
