package main

import (
	"sort"
	"strconv"
	"strings"
)

func removei(s []card, i int) []card {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func remove(s []card, c card) []card {
	i := isIn(c, s)
	return removei(s, i)
}
func isIn(c card, h []card) int {
	for i, a := range h {
		if c == a {
			return i
		}
		if a.color >= 4 && c.joker == a.color-3 {
			return i
		}
	}
	return -1
}

func cleanJoker(c card) card {
	c.number = 0
	c.color = c.joker + 3
	return c
}

func calculateScores(won int) []int {
	scores := []int{0, 0, 0, 0}
	sum := 0
	for i := 0; i < players; i++ {
		for _, c := range hands[i] {
			if c.joker != 0 {
				scores[i] -= 30
			} else {
				scores[i] -= c.number
			}
		}
		sum -= scores[i]
	}
	scores[won] = sum
	return scores
}

func isPoolEmpty() bool {
	return len(pool) == 0
}

func hasWon() int {
	for i, h := range hands {
		if len(h) == 0 {
			return i
		}
	}
	return -1
}

func whatColor(c string) int {
	switch c {
	case "b", "black":
		return 0
	case "y", "yellow":
		return 1
	case "u", "blue":
		return 2
	case "r", "red":
		return 3
	case "jb", "jokerblack", "bj":
		return 4
	case "jr", "jokerred", "rj":
		return 5
	}
	return -1
}

func isValid(b []card) bool {
	validRun := true
	validGroup := true
	if len(b) < 3 {
		return false
	}
	color := b[0].color
	count := b[0].number
	colors := []bool{false, false, false, false} // black, yellow, blue, red
	num := b[0].number
	colors[color] = true
	for i := 1; i < len(b); i++ {
		if b[i].color == color && b[i].number == count+1 {
			count++
		} else {
			validRun = false
		}
		if b[i].number != num {
			validGroup = false
		}
		if !colors[b[i].color] {
			colors[b[i].color] = true
		} else {
			validGroup = false
		}
	}
	if validRun || validGroup {
		return true
	}
	return false
}

func sortHand(h []card) []card {
	// Sorting algorithm
	// First by color, then by ascending number
	// Last cards are jokers
	sort.Slice(h, func(i, j int) bool {
		if h[i].color < h[j].color {
			return true
		} else if h[i].color > h[j].color {
			return false
		}
		return h[i].number < h[j].number
	})

	return h
}

func processItem(item string) (card, int) {
	// Normal items:
	// b-11
	// Empty jokers:
	// jr
	// Signed jokers:
	// jr-y-8
	var c card
	c.joker = 0
	items := strings.Split(item, "-")
	l := len(items)
	if l > 3 {
		return c, 1
	}
	color := whatColor(items[0])
	switch color {
	case 0, 1, 2, 3:
		if l != 2 {
			return c, 1
		}
		n, err := strconv.Atoi(items[1])
		if err != nil {
			return c, 1
		}
		if n < 1 || n > 13 {
			return c, 1
		}
		c.color = color
		c.number = n
	case 4, 5:
		if color == 4 {
			c.joker = 1
		} else {
			c.joker = 2
		}
		if l == 1 {
			c.color = -1
			c.number = 0
		} else if l == 3 {
			c.color = whatColor(items[1])
			switch c.color {
			case 0, 1, 2, 3:
				n, err := strconv.Atoi(items[2])
				if err != nil {
					return c, 1
				}
				if n < 1 || n > 13 {
					return c, 1
				}
				c.number = n
			default:
				return c, 1
			}
		} else {
			return c, 1
		}
	default:
		return c, 1
	}
	return c, 0
}
