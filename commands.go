package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/inancgumus/screen"
)

/*
	TODO:
	I've noticed that when you have a card both in your hold and in your hand, you may want to choose from where
	you want to utilize the card. In reality, hold is always preffered so you can close out a move, but
	having the option is better.
*/

func add(items []string, to string, h []card, b [][]card, ho []card) (bool, []card, [][]card, []card) {
	var cs []card
	num, err := strconv.Atoi(to)
	if err != nil {
		return false, h, b, ho
	}
	num--
	if num < 0 || num >= len(b) {
		return false, h, b, ho
	}

	var bn = make([]card, len(b[num]))
	copy(bn, b[num])

	for _, item := range items {
		c, fail := processItem(item)
		if fail == 1 {
			return false, h, b, ho
		}
		cs = append(cs, c)
		bn = append(bn, c)
	}
	bn = sortHand(bn)
	if !isValid(bn) {
		return false, h, b, ho
	}
	for _, c := range cs {
		if isIn(c, h) != -1 {
			h = remove(h, c)
		} else if isIn(c, ho) != -1 {
			ho = remove(ho, c)
		} else {
			return false, h, b, ho
		}
	}
	b[num] = bn
	return true, h, b, ho
}

func place(items []string, h []card, b [][]card, ho []card) (bool, []card, [][]card, []card, int) {
	var cs []card
	tot := 0
	for _, item := range items {
		c, fail := processItem(item)
		if fail == 1 {
			fmt.Println("Process failed")
			return false, h, b, ho, tot
		}
		cs = append(cs, c)
	}
	cs = sortHand(cs)
	if !isValid(cs) {
		fmt.Println("Validity failed")
		return false, h, b, ho, tot
	}
	for _, c := range cs {
		if isIn(c, h) != -1 {
			h = remove(h, c)
			tot += c.number
		} else if isIn(c, ho) != -1 {
			ho = remove(ho, c)
		} else {
			fmt.Println("In failed")
			return false, h, b, ho, tot
		}
	}
	b = append(b, cs)
	return true, h, b, ho, tot
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

func exchange(item string, to string, h []card, b [][]card, ho []card) (bool, []card, [][]card, []card) {
	num, err := strconv.Atoi(to)
	if err != nil {
		return false, h, b, ho
	}
	num--
	if num < 0 || num >= len(b) {
		return false, h, b, ho
	}
	c, fail := processItem(item)
	if fail == 1 {
		return false, h, b, ho
	}
	cjok := c
	cjok.joker = 1
	i := isIn(cjok, b[num])
	if i == -1 {
		cjok.joker = 2
	}
	i = isIn(cjok, b[num])
	if i == -1 {
		return false, h, b, ho
	}
	h = remove(h, c)
	b[num][i].joker = 0
	ho = append(ho, cleanJoker(cjok))
	return true, h, b, ho
}

func pick(items []string, from string, b [][]card, ho []card) (bool, [][]card, []card) {
	num, err := strconv.Atoi(from)
	if err != nil {
		return false, b, ho
	}
	num--
	if num < 0 || num >= len(b) {
		return false, b, ho
	}
	var bn = make([]card, len(b[num]))
	copy(bn, b[num])
	for _, item := range items {
		c, fail := processItem(item)
		if fail == 1 {
			return false, b, ho
		}
		i := isIn(c, bn)
		if i == -1 {
			return false, b, ho
		}
		bn = removei(bn, i)
		ho = append(ho, c)
	}
	if !isValid(bn) {
		return false, b, ho
	}
	b[num] = bn
	return true, b, ho
}

func pickall(from string, b [][]card, ho []card) (bool, [][]card, []card) {
	num, err := strconv.Atoi(from)
	if err != nil {
		return false, b, ho
	}
	num--
	if num < 0 || num >= len(b) {
		return false, b, ho
	}
	ho = append(ho, b[num]...)
	b = removeBi(b, num)
	return true, b, ho
}

func breakLevel(level string, on string, b [][]card) (bool, [][]card) {
	num, err := strconv.Atoi(level)
	if err != nil {
		return false, b
	}
	num--
	if num < 0 || num >= len(b) {
		return false, b
	}
	point, err := strconv.Atoi(on)
	if err != nil {
		return false, b
	}
	if point < 3 || point >= len(b[num])-2 {
		return false, b
	}
	left := b[num][point:]
	right := b[num][:point]
	b[num] = left
	b = append(b, right)
	return true, b
}
