package main

import (
	"BlackjackSim/cards"
	"BlackjackSim/models"
	"github.com/schollz/progressbar/v3"
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

/***********
	Game
 **********/

func main() {

	shoe := models.NewShoe(deckCount)

	//for _, card := range house.Cards {
	//	fmt.Print(card.Name, " ")
	//}
	//fmt.Println(house.Calculate())
	//
	//for _, card := range player.Hands[0].Cards {
	//	fmt.Print(card.Name, " ")
	//}
	//fmt.Println(player.Hands[0].Calculate())

	bar := progressbar.Default(int64(cycles))
	for i := 0; i < cycles; i++ {

		// Check if shoe needs to be shuffled
		if shoe.Index >= deckSize*shuffleAt {
			shoe.Shuffle()
		}

		// Deal
		house, player := deal(shoe)

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

			if hand.Cards[0] == hand.Cards[1] {
				newHand := models.Hand{
					Cards: []cards.Card{
						hand.Cards[1],
						shoe.NextCard(),
					},
				}

				player.Hands = append(player.Hands, newHand)
				handCount++
			}
		}

		err := bar.Add(1)
		if err != nil {
			panic(err)
		}
	}

	println(account)
}

// deal returns house hand model and player model
func deal(shoe *models.Shoe) (models.Hand, models.Player) {

	house := models.Hand{
		Cards: []cards.Card{
			shoe.NextCard(),
			shoe.NextCard(),
		},
	}
	player := models.Player{
		Hands: []models.Hand{
			{
				Cards: []cards.Card{
					shoe.NextCard(),
					shoe.NextCard(),
				},
			},
		},
	}

	house.UpCard = house.Cards[1]

	return house, player
}
