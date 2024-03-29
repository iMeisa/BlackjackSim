package main

import (
	"BlackjackSim/models"
	"BlackjackSim/strats"
	"github.com/schollz/progressbar/v3"
)

// TODO Implement basic strategy

// Config
// Dev
var debug = true
var cycles = 1000

const baseBet = 25
const minBet = 5
const maxBet = 500

const leaveCount = -3

// Game
var deckCount = 6
var shuffleAt = 5 // How many decks are played
var account = 100000

const blackjackMultiplier = 1.5
const deckSize = 52

/***********
	Game
 **********/

var (
	hardTotal, softTotal, splitting = strats.Load()
	shoe                            = models.NewShoe(deckCount)
)

func main() {

	// Games
	bar := progressbar.Default(int64(cycles))
	for i := 0; i < cycles; i++ {

		// Check if shoe needs to be shuffled
		if shoe.Index >= deckSize*shuffleAt {
			shoe.Shuffle()
		}

		// Deal
		house, player := deal(shoe, strats.GetBet(minBet, maxBet, baseBet, shoe.TrueCount()))

		// Check if house has blackjack
		houseBlackjack := house.Calculate() == 21

		handCount := len(player.Hands)

		// Play through player hands
		for j := 0; j < handCount; j++ {
			hand := player.Hands[j]
			//fmt.Println(hand.Bet)

			playerBlackjack := hand.Calculate() == 21

			if houseBlackjack {
				if playerBlackjack {
					continue
				}
				account -= hand.Bet
			}

			// Check for blackjack
			if playerBlackjack {
				account += int(float64(hand.Bet) * blackjackMultiplier)
				continue
			}

			// Play hand
			account += playHand(&house, &hand, &player)

			// Update hand count
			handCount = len(player.Hands)

			//for _, card := range house.Cards {
			//	fmt.Print(" " + card.Name)
			//}
			//fmt.Print(",")
			//for _, card := range hand.Cards {
			//	fmt.Print(" " + card.Name)
			//}
			//fmt.Printf(": %d %d\n", shoe.RunningCount, shoe.TrueCount())
		}

		if shoe.TrueCount() <= leaveCount {
			// Leave
			shoe.Shuffle()
		}

		err := bar.Add(1)
		if err != nil {
			panic(err)
		}
	}

	println(account)
}
