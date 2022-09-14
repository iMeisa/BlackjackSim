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
var baseBet = 5

// Game
var deckCount = 1
var shuffleAt = 0 // How many decks are played
var account = 1000

const blackjackMultiplier = 1.5
const deckSize = 52

/***********
	Game
 **********/

var (
	hardTotal, softTotal, splitting = strats.Load()
	shoe                            = models.NewShoe(deckCount)
)

var (
	runningCount = 0
	trueCount    = 0
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
		house, player := deal(shoe, strats.GetBet(baseBet))

		// Check if house has blackjack
		houseBlackjack := house.Calculate() == 21

		handCount := len(player.Hands)

		// Play through player hands
		for j := 0; j < handCount; j++ {
			hand := player.Hands[j]

			playerBlackjack := hand.Calculate() == 21

			if houseBlackjack {
				if playerBlackjack {
					continue
				}
				account -= baseBet
			}

			// Check for blackjack
			if playerBlackjack {
				account += int(float64(baseBet) * blackjackMultiplier)
				continue
			}

			// Play hand
			account += playHand(&house, &hand)
		}

		err := bar.Add(1)
		if err != nil {
			panic(err)
		}
	}

	println(account)
}
