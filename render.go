package main

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	black   *color.Color
	yellow  *color.Color
	blue    *color.Color
	red     *color.Color
	bg      *color.Color
	counter *color.Color

	playerLet []string = []string{"A", "B", "C", "D"}
)

func setupColors() {
	black = color.New(color.BgHiBlack).Add(color.FgBlack)
	yellow = color.New(color.BgHiBlack).Add(color.FgYellow)
	blue = color.New(color.BgHiBlack).Add(color.FgBlue)
	red = color.New(color.BgHiBlack).Add(color.FgRed)
	bg = color.New(color.BgHiBlack)
	counter = color.New(color.BgBlack).Add(color.FgWhite)
}

func renderBoard() {
	for i := 0; i < len(board); i++ {
		counter.Print(i)
		printCards(board[i])
	}
}

func renderHand() {
	fmt.Println("\nPlayer's hand:")
	printCards(hands[turn])
}

func printCards(cards []card) {
	bg.Print(" ")
	for i := 0; i < len(cards); i++ {
		printCard(cards[i])
		bg.Print(" ")
	}
	fmt.Println()
}

func printCard(c card) {
	switch c.color {
	case 'r':
		red.Print(c.number)
	case 'b':
		black.Print(c.number)
	case 'y':
		yellow.Print(c.number)
	case 'u':
		blue.Print(c.number)
	case 'j':
		if c.number == 0 {
			black.Print("J")
		} else {
			red.Print("J")
		}
	}
}
