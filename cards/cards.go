package cards

type Card struct {
	Face  bool
	Name  string
	Value int
}

var (
	Ace   = Card{Face: true, Name: "Ace", Value: 1}
	Two   = Card{Face: false, Name: "Two", Value: 2}
	Three = Card{Face: false, Name: "Three", Value: 3}
	Four  = Card{Face: false, Name: "Four", Value: 4}
	Five  = Card{Face: false, Name: "Five", Value: 5}
	Six   = Card{Face: false, Name: "Six", Value: 6}
	Seven = Card{Face: false, Name: "Seven", Value: 7}
	Eight = Card{Face: false, Name: "Eight", Value: 8}
	Nine  = Card{Face: false, Name: "Nine", Value: 9}
	Ten   = Card{Face: true, Name: "Ten", Value: 10}
	Jack  = Card{Face: true, Name: "Jack", Value: 10}
	Queen = Card{Face: true, Name: "Queen", Value: 10}
	King  = Card{Face: true, Name: "King", Value: 10}
)

var Cards = []Card{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
