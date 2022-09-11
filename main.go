package main

import (
	"BlackjackSim/models"
	"fmt"
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

	for _, card := range shoe.Cards {
		fmt.Print(card.Name + " ")
	}
	fmt.Println("\n", len(shoe.Cards))

	//bar := progressbar.Default(int64(cycles))

	//for i := 0; i < cycles; i++ {
	//	err := bar.Add(1)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//
	//	time.Sleep(100 * time.Millisecond)
	//}
}
