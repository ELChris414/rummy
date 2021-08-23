package main

import (
	"fmt"
	"math/rand"
	"time"
)

type card struct {
	color  rune
	number int
}

var (
	board      [][]card
	hands      [][]card
	pool       []card
	turn       int
	rounds     int
	overThirty bool
)

func remove(s []card, i int) []card {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func setupCards() {
	colors := []rune{'b', 'y', 'u', 'r'}
	// b Black
	// y Yellow
	// u Blue
	// r Red
	for i := 1; i < 14; i++ {
		for _, c := range colors {
			pool = append(pool, card{c, i})
			pool = append(pool, card{c, i})
		}
	}
	pool = append(pool, card{'j', 0})
	pool = append(pool, card{'j', 1})
}

func initializeGame() {
	setupColors()
	setupCards()
	turn = 0
	rounds = 0
	overThirty = false

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 4; i++ {
		var hand []card
		for y := 0; y < 14; y++ {
			randomIndex := rand.Intn(len(pool))
			pick := pool[randomIndex]
			hand = append(hand, pick)
			pool = remove(pool, randomIndex)
		}
		hands = append(hands, hand)
	}
}

func startGame() {
	renderBoard()
	//checkLegalMoves(hands[turn])
}

func isOverThirty() {
	sum := 0
	for i := 0; i < len(board); i++ {
		sum += len(board[i])
	}
	if sum >= 30 {
		fmt.Println("Sum of cards is over 30! Group breaking is now permitted")
		overThirty = true
	}
}

/*func checkLegalMoves(cards []card) bool {
	// Check for addition ability on board

}*/
