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

func renderBoard(b [][]card) {
	for i := 0; i < len(b); i++ {
		counter.Print(i)
		printCards(b[i])
	}
}

func renderHand(cards []card) {
	fmt.Println("\nPlayer's hand:")
	printCards(cards)
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
	switch c.joker {
	case 1:
		black.Print("J")
		return
	case 2:
		red.Print("J")
		return
	}
	switch c.color {
	case 0:
		black.Print(c.number)
	case 1:
		yellow.Print(c.number)
	case 2:
		blue.Print(c.number)
	case 3:
		red.Print(c.number)
	}
}
