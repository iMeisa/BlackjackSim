package main

import (
	"BlackjackSim/cards"
	"BlackjackSim/models"
	"BlackjackSim/strats"
)

// deal returns house hand model and player model
func deal(shoe *models.Shoe, bet int) (models.Hand, models.Player) {

	house := models.Hand{
		Cards: []cards.Card{
			shoe.NextCard(),
			shoe.NextCard(),
		},
	}
	player := models.Player{
		Hands: []models.Hand{
			{
				Bet: bet,
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

func playHand(house, playerHand *models.Hand, player *models.Player) int {

	// Check for blackjack already handled by main()

	// Decide on player hand
PlayerHand:
	for {
		decision := strats.Decide(house.UpCard, playerHand, hardTotal, softTotal, splitting)

		//var playerCards string
		//for _, card := range playerHand.Cards {
		//	if card.Face {
		//		faceLetter := string(card.Name[0])
		//		playerCards += " " + faceLetter
		//	} else {
		//		playerCards += " " + strconv.Itoa(card.Value)
		//	}
		//}
		//
		//fmt.Printf("%s,%s: %s\n", house.UpCard.Name, playerCards, decision)

		// Check if player busted
		if playerHand.Calculate() > 21 {
			//updateRunningCount(house.Cards[0]) // Update running count for hidden card
			return -playerHand.Bet
		}

		switch decision {
		case strats.Hit:
			hit(playerHand, shoe)
		case strats.Double:
			playerHand.Bet *= 2
			hit(playerHand, shoe)
		case strats.Split:
			player.AddHand(split(playerHand, shoe))
		case strats.Stand:
			break PlayerHand
		}
	}

	// Decide on house hand
	//updateRunningCount(house.Cards[0]) // Update running count for hidden card
HouseHand:
	for {
		switch strats.HouseDecision(house) {
		case strats.Hit:
			hit(house, shoe)
		case strats.Stand:
			break HouseHand
		}
	}

	// Check if house busted
	if house.Calculate() > 21 {
		return playerHand.Bet
	}

	// Compare hands
	if house.Calculate() > playerHand.Calculate() {
		return -playerHand.Bet
	} else if house.Calculate() < playerHand.Calculate() {
		return playerHand.Bet
	}

	return 0
}

func hit(hand *models.Hand, shoe *models.Shoe) {
	hand.Cards = append(hand.Cards, shoe.NextCard())
}

func split(hand *models.Hand, shoe *models.Shoe) models.Hand {
	newHand := models.Hand{
		Bet: hand.Bet,
		Cards: []cards.Card{
			hand.Cards[1],
		},
	}

	hand.Cards = []cards.Card{
		hand.Cards[0],
	}

	hit(hand, shoe)
	hit(&newHand, shoe)

	return newHand
}
