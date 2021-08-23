package main

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	black  *color.Color
	yellow *color.Color
	blue   *color.Color
	red    *color.Color
	bg     *color.Color
)

func setupColors() {
	black = color.New(color.BgHiBlack).Add(color.FgBlack)
	yellow = color.New(color.BgHiBlack).Add(color.FgYellow)
	blue = color.New(color.BgHiBlack).Add(color.FgBlue)
	red = color.New(color.BgHiBlack).Add(color.FgRed)
	bg = color.New(color.BgHiBlack)
}

func renderBoard() {
	fmt.Printf("Currently playing: Player %v\n", turn)
	fmt.Printf("Cards on board: %v\n", len(pool))
	fmt.Printf("Round: %v\n\n", rounds)
	for i := 0; i < len(board); i++ {
		printCards(board[i])
	}
	fmt.Println("\nPlayer's hand:")
	printCards(hands[turn])
}

func printCards(cards []card) {
	bg.Print(" ")
	for i := 0; i < len(cards); i++ {
		switch cards[i].color {
		case 'r':
			red.Print(cards[i].number)
		case 'b':
			black.Print(cards[i].number)
		case 'y':
			yellow.Print(cards[i].number)
		case 'u':
			blue.Print(cards[i].number)
		case 'j':
			if cards[i].number == 0 {
				black.Print("J")
			} else {
				red.Print("J")
			}
		}
		bg.Print(" ")
	}
	fmt.Println()
}
