package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/inancgumus/screen"
)

type card struct {
	color  rune
	number int
}

var (
	players    int
	board      [][]card
	hands      [][]card
	pool       []card
	hold       []card
	turn       int = 0
	rounds     int = 0
	jokerRed   card
	jokerBlack card
	laid       []bool
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

	fmt.Println("How many players are going to play? (2-4)")
	_, err := fmt.Scanln(&players)
	for err != nil || players > 4 || players < 2 {
		fmt.Println("Repeat answer")
		_, err = fmt.Scanln(&players)
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < players; i++ {
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

func hasWon() int {
	for i := 0; i < players; i++ {
		if len(hands[i]) == 0 {
			return i
		}
	}
	return -1
}

func startGame() {
	won := hasWon()
	for won == -1 {
		playTurn()
		turn++
		if turn == players {
			turn = 0
		}
	}
	fmt.Printf("Player %s has won the game!\n", playerLet[won])
	fmt.Printf("Game finished in %v rounds\n", rounds)
	fmt.Println("Final board:")

}

func isPoolEmpty() bool {
	return len(pool) == 0
}

func playTurn() {
	screen.Clear()
	fmt.Printf("Currently playing: Player %s\n", playerLet[turn])
	fmt.Printf("Cards on board: %v\n", len(pool))
	fmt.Printf("Round: %v\n\n", rounds)
	renderBoard()
	renderHand()
	fmt.Print("Possible actions: ")
	if !isPoolEmpty() {
		fmt.Print(" draw ")
	}
	fmt.Println(" move ")
	var action string
	done := false
	for !done {
		fmt.Scanln(&action)
		switch action {
		case "draw":
			draw()
			done = true
		default:
			fmt.Println("Incorrect command, try again.")
		}
	}
}

func draw() {
	screen.Clear()
	randomIndex := rand.Intn(len(pool))
	pick := pool[randomIndex]
	hands[turn] = append(hands[turn], pick)
	pool = remove(pool, randomIndex)
	fmt.Print("You drew: ")
	printCard(pick)
	fmt.Println("\nYour new hand is:")
	printCards(hands[turn])
	fmt.Scanln()
}
