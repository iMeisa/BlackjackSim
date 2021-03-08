package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// TODO Implement basic strategy
// TODO Implement betting

// Config
// Dev
var debug = true
var cycles = 1000

// Game
var deckCount = 6
var shuffleAt = 5 // How many decks are played
var account = 1000

const deckSize = 52

/*
Simulated Player
*/

const (
	hit   = "Hit!"
	stand = "Stand"
	dd    = "Double Down!"
	split = "Split!"
)

func playerDecide(dealerUpCard string, hand Hand) string {
	var value = hand.value

	if hand.hand[0] == hand.hand[1] {
		return split
	}
	if value >= 17 {
		return stand
	}
	if value >= 12 {
		return hit
	}
	if value >= 9 {
		return dd
	}
	return hit
}

func dealerDecide(value int) string {
	if value >= 17 {
		return stand
	}
	return hit
}

/***********
	Game
 **********/

// Random functions
func wait(seconds float32) {
	waitTime := seconds * 1e9
	time.Sleep(time.Duration(waitTime))
}

// Deck config
type Stack struct {
	decks int
	cards []string
	shoe  []string
}

func buildDeck(deckCount int, cards []string) []string {
	var shoe = make([]string, 0, len(cards))
	for i := 0; i < deckCount; i++ {
		for j := 0; j < 4; j++ {
			shoe = append(shoe, cards...)
		}
	}
	return shoe
}

func (shoe Stack) shuffleShoe() {
	a := shoe.shoe
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}

func calculateHand(hand []string) int {
	var totalValue int
	var aces = 0
	for _, card := range hand {
		switch card {
		case "A":
			aces++
		case "K", "Q", "J", "10":
			totalValue += 10
		default:
			cardValue, _ := strconv.Atoi(card)
			totalValue += cardValue
		}
	}

	// Aces calculation
	totalValue += aces
	if aces > 0 && totalValue+10 <= 21 {
		totalValue += 10
	}

	return totalValue
}

type Hand struct {
	hand  []string
	value int
	bet   int
}

type Player struct {
	hands []Hand
	bust  bool
}

func (player *Player) addHand(hand Hand) []Hand {
	player.hands = append(player.hands, hand)
	return player.hands
}

func main() {
	fmt.Println("Hello World!")
	if debug {
		wait(1)
	}

	bj := Stack{
		decks: deckCount,
		cards: []string{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"},
	}

	bj.shoe = buildDeck(bj.decks, bj.cards)
	bj.shuffleShoe()

	currentIndex := 0
	wins := 0
	for i := 0; i < cycles; i++ {

		dealer := Hand{
			hand: []string{
				bj.shoe[currentIndex], bj.shoe[currentIndex+1]},
		}
		dealer.value = calculateHand(dealer.hand)

		player := Player{
			hands: []Hand{
				{hand: []string{bj.shoe[currentIndex+2], bj.shoe[currentIndex+3]}},
			},
		}
		player.hands[0].value = calculateHand(player.hands[0].hand)

		currentIndex += 4

		// Check if blackjack
		if dealer.value == 21 && player.hands[0].value != 21 {
			if debug {
				fmt.Printf("\n%v\n%v\n", dealer, player)
				fmt.Println("Dealer Blackjack!")
			}
			continue
		} else if dealer.value == 21 && player.hands[0].value == 21 {
			if debug {
				fmt.Printf("\n%v\n%v\n", dealer, player)
				fmt.Println("Blackjack Push!")
			}
			continue
		} else if player.hands[0].value == 21 {
			if debug {
				fmt.Printf("\n%v\n%v\n", dealer, player)
				fmt.Println("Blackjack!")
			}
			continue
		}

		// Player plays
		var totalBusts int
		for index := 0; index < len(player.hands); index++ {
			hand := player.hands[index]
		PlayerDeal:
			for hand.value < 21 {
				if hand.value == 21 && len(hand.hand) == 2 {
					if debug {
						fmt.Printf("\n%v\n%v\n", dealer, player)
						fmt.Println("Blackjack!")
					}
					break
				}

				choice := playerDecide(dealer.hand[1], hand)
				switch choice {
				case stand:
					player.hands[index] = hand
					break PlayerDeal
				case hit:
					hand.hand = append(hand.hand, bj.shoe[currentIndex])
					currentIndex++
					hand.value = calculateHand(hand.hand)
				case dd:
					hand.hand = append(hand.hand, bj.shoe[currentIndex])
					currentIndex++
					hand.value = calculateHand(hand.hand)
					player.hands[index] = hand
					break PlayerDeal
				case split:
					newHand := Hand{
						hand: []string{
							hand.hand[1],
							bj.shoe[currentIndex],
						},
					}
					currentIndex++
					newHand.value = calculateHand(newHand.hand)
					player.addHand(newHand)

					hand.hand = []string{hand.hand[0], bj.shoe[currentIndex]}
					currentIndex++
					hand.value = calculateHand(hand.hand)
				}

				if hand.value > 21 {
					totalBusts++
				}

				player.hands[index] = hand
			}
		}

		if totalBusts == len(player.hands) {
			player.bust = true
		}

		if !player.bust {
		DealerDeal:
			for dealer.value < 21 {
				choice := dealerDecide(dealer.value)
				switch choice {
				case stand:
					break DealerDeal
				case hit:
					dealer.hand = append(dealer.hand, bj.shoe[currentIndex])
					currentIndex++
					dealer.value = calculateHand(dealer.hand)
				}
			}
		}

		fmt.Printf("\n%v\n%v\n", dealer, player)
		if dealer.value > 21 {
			if debug {
				fmt.Print("Dealer Bust!\n")
			}
			for _, hand := range player.hands {
				if hand.value <= 21 {
					wins++
				}
			}
		} else {
			for _, hand := range player.hands {
				if hand.value == 21 && len(hand.hand) == 2 {
					continue
				} else if hand.value > 21 {
					if debug {
						fmt.Println("Bust!")
					}
				} else if dealer.value < hand.value {
					if debug {
						fmt.Println("Win!")
					}
					wins++
				} else if dealer.value == hand.value {
					if debug {
						fmt.Println("Push!")
					}
				} else {
					if debug {
						fmt.Println("Dealer Win!")
					}
				}
			}
		}

		if currentIndex > (shuffleAt * deckSize) {
			bj.shuffleShoe()
			currentIndex = 0
		}
	}

	fmt.Printf("\nTotal Wins: %v", wins)
}
