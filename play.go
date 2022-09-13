package main

import (
	"BlackjackSim/cards"
	"BlackjackSim/models"
)

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

func playHand(house *models.Hand, playerHand *models.Hand, total, total2, splitting *map[int]map[int]string) {

}
