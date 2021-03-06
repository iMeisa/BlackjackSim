package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// TODO Implement splitting
// TODO Implement basic strategy
// TODO Implement betting

// Config
// Dev
var debug = false
var cycles = 1000

// Game
var deckCount = 6
var shuffleAt = 1
var account = 1000

/*
Simulated Player
 */

const (
	hit = "Hit!"
	stand = "Stand"
	dd = "Double Down!"
)

func decide(dealer, player Hand, isDealer bool) string{
	var value = player.value
	if isDealer{value = dealer.value}
	
	if value >= 17{return stand}
	if value >= 12{return hit}
	if !isDealer && value >= 9{return dd}
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
	decks     int
	cards     []string
	shoe      []string
}

func buildDeck(deckCount int, cards []string) []string{
	var shoe []string
	for i := 0; i < deckCount; i++ {
		for j := 0; j < 4; j++ {
			shoe = append(shoe, cards...)
		}
	}
	return shoe
}

func (shoe Stack) shuffleShoe(){
	a := shoe.shoe
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}


func calculateHand(hand []string) int{
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
	if aces > 0 && totalValue + 10 <= 21 {
		totalValue += 10
	}

	return totalValue
}


type Hand struct {
	hand	[]string
	value   int
}


func main() {
	fmt.Println("Hello World!")
	if debug {wait(1)}

	bj := Stack {
		decks: deckCount,
		cards: []string {"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"},
	}

	bj.shoe = buildDeck(bj.decks, bj.cards)
	bj.shuffleShoe()

	currentIndex := 0
	wins := 0
	for i := 0; i < cycles; i++ {

		dealer := Hand{hand: []string{bj.shoe[currentIndex], bj.shoe[currentIndex + 1]}}
		dealer.value = calculateHand(dealer.hand)
		player := Hand{hand: []string{bj.shoe[currentIndex + 2], bj.shoe[currentIndex + 3]}}
		player.value = calculateHand(player.hand)

		currentIndex += 4

		// Player plays
		PlayerDeal:
			for player.value < 21 {
				choice := decide(dealer, player, false)
				switch choice{
				case stand:
					break PlayerDeal
				case hit:
					player.hand = append(player.hand, bj.shoe[currentIndex])
					currentIndex++
					player.value = calculateHand(player.hand)
				case dd:
					player.hand = append(player.hand, bj.shoe[currentIndex])
					currentIndex++
					player.value = calculateHand(player.hand)
					break PlayerDeal
				}
			}

		if player.value > 21 {
			fmt.Printf("\n%v\n%v\nBust!\n", dealer, player)
			continue
		}

		var dealerBust = false
		DealerDeal:
			for dealer.value < 21 {
				choice := decide(dealer, player, true)
				switch choice{
				case stand:
					break DealerDeal
				case hit:
					dealer.hand = append(dealer.hand, bj.shoe[currentIndex])
					currentIndex++
					dealer.value = calculateHand(dealer.hand)
				}
			}

		if dealer.value > 21 {
			dealerBust = true
			fmt.Print("\nDealer Bust!")
		}
		fmt.Printf("\n%v\n%v\n", dealer, player)

		if dealer.value < player.value || dealerBust {
			fmt.Println("Win!")
			wins++
		}

		if currentIndex < (shuffleAt * 52){
			bj.shuffleShoe()
			currentIndex = 0
		}
	}

	fmt.Printf("\nTotal Wins: %v", wins)
}
