package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// TODO Implement basic strategy

// Config
// Dev
var debug = true
var cycles = 1000
var baseBet = 10

// Game
var deckCount = 6
var shuffleAt = 5 // How many decks are played
var account = 1000

const blackjackMultiplier = 1.5
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

func playerDecide(dealerUpCard string, hand Hand, firstHit bool) string {
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
	if value >= 9 && firstHit {
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
	if !debug {
		fmt.Print("Progress: [")
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
			hands: []Hand{{
				hand: []string{
					bj.shoe[currentIndex+2], bj.shoe[currentIndex+3],
				},
				bet: baseBet,
			},
			},
		}
		player.hands[0].value = calculateHand(player.hands[0].hand)

		currentIndex += 4

		// Check if blackjack
		if dealer.value == 21 && player.hands[0].value != 21 {
			account -= player.hands[0].bet
			if debug {
				fmt.Printf("\n%v\n%v\n", dealer, player)
				fmt.Println("Dealer Blackjack!")
				fmt.Println(account)
			}
			continue
		} else if dealer.value == 21 && player.hands[0].value == 21 {
			if debug {
				fmt.Printf("\n%v\n%v\n", dealer, player)
				fmt.Println("Blackjack Push!")
			}
			continue
		} else if player.hands[0].value == 21 {
			account += int(float32(player.hands[0].bet) * blackjackMultiplier)
			wins++
			if debug {
				fmt.Printf("\n%v\n%v\n", dealer, player)
				fmt.Println("Blackjack!")
				fmt.Println(account)
			}
			continue
		}

		// Player plays
		var totalBusts int
		for index := 0; index < len(player.hands); index++ {
			hand := player.hands[index]
			firstHit := true
		PlayerDeal:
			for hand.value < 21 {

				choice := playerDecide(dealer.hand[1], hand, firstHit)
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
					hand.bet *= 2
					player.hands[index] = hand
					break PlayerDeal
				case split:
					newHand := Hand{
						hand: []string{
							hand.hand[1],
							bj.shoe[currentIndex],
						},
						bet: hand.bet,
					}
					currentIndex++
					newHand.value = calculateHand(newHand.hand)
					player.addHand(newHand)

					hand.hand = []string{hand.hand[0], bj.shoe[currentIndex]}
					currentIndex++
					hand.value = calculateHand(hand.hand)

					// If blackjack after split
					if hand.value == 21 {
						account += int(float32(hand.bet) * blackjackMultiplier)
						wins++
						player.hands[index] = hand
						break PlayerDeal
					}

					// If blackjack after split
					if newHand.value == 21 {
						account += int(float32(newHand.bet) * blackjackMultiplier)
						wins++
						player.hands[index] = hand
						break PlayerDeal
					}
					continue
				}

				if hand.value > 21 {
					totalBusts++
					account -= hand.bet
				}

				player.hands[index] = hand
				firstHit = false
			}
		}

		if totalBusts == len(player.hands) {
			if debug {
				fmt.Printf("\n%v\n%v\n", dealer, player)
				fmt.Println("Complete Bust!")
				fmt.Println(account)
			}
			player.bust = true
			continue
		}

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

		if debug {
			fmt.Printf("\n%v\n%v\n", dealer, player)
		}
		if dealer.value > 21 {
			if debug {
				fmt.Print("Dealer Bust!\n")
			}
			for _, hand := range player.hands {
				if hand.value <= 21 {
					account += hand.bet
					wins++
				}
			}
		} else {
			for _, hand := range player.hands {
				if hand.value == 21 && len(hand.hand) == 2 {
					fmt.Println("Blackjack!")
					continue
				}
				if hand.value > 21 {
					if debug {
						fmt.Println("Bust!")
					}
				} else if dealer.value < hand.value {
					if debug {
						fmt.Println("Win!")
					}
					account += hand.bet
					wins++
				} else if dealer.value == hand.value {
					if debug {
						fmt.Println("Push!")
					}
				} else {
					if debug {
						fmt.Println("Dealer Win!")
					}
					account -= hand.bet
				}
			}
		}

		if debug {
			fmt.Println(account)
		}

		if currentIndex > (shuffleAt * deckSize) {
			bj.shuffleShoe()
			currentIndex = 0
		}

		// Print progress bar
		if i%(cycles/50) == 0 && !debug {
			fmt.Print("#")
		}
	}

	if !debug {
		fmt.Println("]")
	}
	fmt.Printf("\nTotal Wins: %v\n", wins)
	fmt.Printf("Net: %v\n", account-1000)
}
