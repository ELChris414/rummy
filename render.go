package main

import (
	"fmt"
	"strings"

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
		counter.Printf("%v: ", i+1)
		printCards(b[i])
		fmt.Println()
	}
	fmt.Println()
}

func renderHand(cards []card) {
	fmt.Println("\nPlayer's hand:")
	printCards(cards)
	fmt.Println()
}

func renderHold(cards []card) {
	if len(cards) != 0 {
		fmt.Println("\nPlayer's hold:")
		printCards(cards)
	}
	fmt.Println()
}

func printCards(cards []card) {
	bg.Print(" ")
	for i := 0; i < len(cards); i++ {
		printCard(cards[i])
		bg.Print(" ")
	}
}

func printCard(c card) {
	switch c.joker {
	case 0:
		switch c.color {
		case 0:
			black.Print(c.number)
		case 1:
			yellow.Print(c.number)
		case 2:
			blue.Print(c.number)
		case 3:
			red.Print(c.number)
		default:
			return
		}
	case 1:
		c.joker = 0
		black.Print("J")
		printCard(c)
		return
	case 2:
		c.joker = 0
		red.Print("J")
		printCard(c)
		return
	}
}

func renderActions(actions []string, total int) {
	fmt.Println()
	if len(actions) != 0 {
		fmt.Println("Current Actions:")
		for i, action := range actions {
			fmt.Printf("%v. ", i+1)
			printAction(action)
		}
		if !laid[turn] {
			fmt.Printf("Total: %v", total)
			if total >= 30 {
				fmt.Print(", valid to submit")
			}
			fmt.Println()
		}
	}
}

func printAction(action string) {
	action = strings.ToLower(action)
	command := strings.Split(action, " ")
	switch command[0] {
	case "add":
		fmt.Print("Player added ")
		var cs []card
		for _, item := range command[2:] {
			c, _ := processItem(item)
			cs = append(cs, c)
		}
		printCards(cs)
		fmt.Printf(" at level %s\n", command[1])
	case "place":
		fmt.Print("Player placed ")
		var cs []card
		for _, c := range command[1:] {
			c, _ := processItem(c)
			cs = append(cs, c)
		}
		printCards(cs)
		fmt.Println()
	case "exchange":
		fmt.Print("Player exchanged ")
		c, _ := processItem(command[1])
		printCardF(c)
		fmt.Print(" with the joker at level ")
		fmt.Println(command[2])
	case "pick":
		fmt.Print("Player picked ")
		var cs []card
		for _, item := range command[2:] {
			c, _ := processItem(item)
			cs = append(cs, c)
		}
		printCards(cs)
		fmt.Printf(" from level %s\n", command[1])
	case "pickall":
		fmt.Printf("Player picked all cards from level %s\n", command[1])
	case "break":
		fmt.Printf("Player broke level %s at point %s\n", command[1], command[2])
	}
	// For pick it must show the whole sequence of num card card card
	// So example: Player grabbed [card] from [num] [card] [card] [card]
}

func printCardF(c card) {
	bg.Print(" ")
	printCard(c)
	bg.Print(" ")
}

func renderScores(scores []int) {
	for i := 0; i < players; i++ {
		fmt.Printf("Player %s: %v\n", playerLet[i], scores[i])
	}
}
