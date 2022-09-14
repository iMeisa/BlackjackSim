package models

import (
	"BlackjackSim/cards"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Shoe struct {
	Index        int
	Cards        []cards.Card
	RunningCount int
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
	if s.Index >= len(s.Cards) {
		s.Shuffle()
	}
	s.updateRunningCount()
	return card
}

func (s *Shoe) Shuffle() {
	a := s.Cards
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	s.Index = 0
	s.updateRunningCount()
	fmt.Println("Shuffling...")
}

func (s *Shoe) TrueCount() int {
	cardsLeft := len(s.Cards) - s.Index
	decksLeft := int(math.Max(float64(cardsLeft/52), 1))
	return s.RunningCount / decksLeft
}

func (s *Shoe) updateRunningCount() {
	s.RunningCount = 0
	for _, card := range s.Cards[:s.Index] {
		switch card.Value {
		case 2, 3, 4, 5, 6:
			s.RunningCount++
		case 10, 1:
			s.RunningCount--
		}
	}
}
