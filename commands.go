package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/inancgumus/screen"
)

func add(items []string, to string, h []card, b [][]card) (bool, []card, [][]card) {
	num, err := strconv.Atoi(to)
	if err != nil {
		return false, h, b
	}
	num--
	if num < 0 || num >= len(b) {
		return false, h, b
	}
	for _, item := range items {
		c, fail := processItem(item)
		if fail == 1 {
			return false, h, b
		}
		if isIn(c, h) == -1 {
			return false, h, b
		}
		b[num] = append(b[num], c)
	}
	if isValid(sortHand(b[num])) {
		return true, h, b
	}
	return false, h, b
}

func place(items []string, h []card, b [][]card) (bool, []card, [][]card, int) {
	var cs []card
	tot := 0
	for _, item := range items {
		c, fail := processItem(item)
		if fail == 1 {
			return false, h, b, tot
		}
		if isIn(c, h) == -1 {
			return false, h, b, tot
		}
		cs = append(cs, c)
	}
	if isValid(cs) {
		for _, c := range cs {
			h = remove(h, c)
			tot += c.number
		}
		b = append(b, cs)
		return true, h, b, tot
	}
	return false, h, b, tot
}

func draw() {
	screen.Clear()
	randomIndex := rand.Intn(len(pool))
	pick := pool[randomIndex]
	hands[turn] = sortHand(append(hands[turn], pick))
	pool = removei(pool, randomIndex)
	fmt.Print("You drew: ")
	printCard(pick)
	fmt.Println("\nYour new hand is:")
	printCards(hands[turn])
	fmt.Scanln()
}
