package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/inancgumus/screen"
)

type card struct {
	color  int // 0 black, 1 yellow, 2 blue, 3 red, 4 emptyJokerB, 5 emptyJokerR
	number int
	joker  int // 0 no joker, 1 blackJoker, 2 redJoker
}

const (
	initialCards = 40
)

var (
	players int
	board   [][]card
	hands   [][]card
	pool    []card
	turn    int    = 0
	rounds  int    = 0
	laid    []bool = []bool{false, false, false, false}
)

func removei(s []card, i int) []card {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func remove(s []card, c card) []card {
	i := isIn(c, s)
	return removei(s, i)
}

func setupCards() {
	colors := []int{0, 1, 2, 3}
	for i := 1; i < 14; i++ {
		for _, c := range colors {
			pool = append(pool, card{c, i, 0})
			pool = append(pool, card{c, i, 0})
		}
	}
	pool = append(pool, card{4, 0, 1})
	pool = append(pool, card{5, 0, 2})
}

func initializeGame() {
	var err error
	setupColors()
	setupCards()

	fmt.Println("How many players are going to play? (2-4)")
	scanner.Scan()
	num := scanner.Text()
	players, err = strconv.Atoi(num)
	//fmt.Println(err)
	for err != nil || players > 4 || players < 2 {
		fmt.Println("Repeat answer")
		scanner.Scan()
		num := scanner.Text()
		players, err = strconv.Atoi(num)
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < players; i++ {
		var hand []card
		for y := 0; y < initialCards; y++ {
			randomIndex := rand.Intn(len(pool))
			pick := pool[randomIndex]
			hand = append(hand, pick)
			pool = removei(pool, randomIndex)
		}
		hands = append(hands, sortHand(hand))
	}
}

func hasWon() int {
	for i, h := range hands {
		if len(h) == 0 {
			return i
		}
	}
	return -1
}

func startGame() {
	won := hasWon()
	for won == -1 {
		screen.Clear()
		fmt.Printf("Player %s's turn, press enter to play.\n", playerLet[turn])
		scanner.Scan()
		playTurn()
		turn++
		if turn == players {
			turn = 0
			rounds++
		}
	}
	fmt.Printf("Player %s has won the game!\n", playerLet[won])
	fmt.Printf("Game finished in %v rounds\n", rounds)
	fmt.Println("Final board:")

}

func isPoolEmpty() bool {
	return len(pool) == 0
}

func playTurn() bool {
	var cpHand [][]card
	var cpBoard [][][]card
	var hold [][]card = [][]card{{}}
	var actions []string
	var total []int = []int{0}
	var tot int
	var h []card
	var ho []card
	var b [][]card

	cpBoard = append(cpBoard, board)
	cpHand = append(cpHand, hands[turn])

	for {
		screen.Clear()
		fmt.Printf("Currently playing: Player %s\n", playerLet[turn])
		fmt.Printf("Cards on board: %v\n", len(pool))
		fmt.Printf("Round: %v\n\n", rounds)
		renderBoard(cpBoard[len(cpBoard)-1])
		renderActions(actions, total[len(total)-1])
		renderHand(cpHand[len(cpHand)-1])
		renderHold(hold[len(hold)-1])
		fmt.Print("Let's play! Give commands. Possible commands:  place")
		if total[len(total)-1] >= 30 || laid[turn] {
			fmt.Print("  add  pick")
		}
		if len(cpHand) > 1 {
			fmt.Print("  restart  undo")
		}
		if len(hold[len(hold)-1]) == 0 && len(cpHand) != 1 {
			fmt.Print("  done")
		}
		if !isPoolEmpty() {
			fmt.Print("  draw")
		}
		fmt.Println()
		scanner.Scan()
		action := scanner.Text()
		action = strings.ToLower(action)
		command := strings.Split(action, " ")

		h = make([]card, len(cpHand[len(cpHand)-1]))
		b = make([][]card, len(cpBoard[len(cpBoard)-1]))
		ho = make([]card, len(hold[len(hold)-1]))

		copy(h, cpHand[len(cpHand)-1])
		copy(b, cpBoard[len(cpBoard)-1])
		copy(ho, hold[len(hold)-1])

		success := false
		switch command[0] {
		case "add":
			if total[len(total)-1] < 30 && !laid[turn] {
				fmt.Println("Initial placements not done yet.")
				break
			}
			if len(command) != 3 {
				fmt.Println("Insufficient amount of arguments.")
				break
			}
			success, h, b, tot = add(command[1], command[2], h, b)
			if success {
				actions = append(actions, action)
				cpHand = append(cpHand, sortHand(h))
				cpBoard = append(cpBoard, b)
				total = append(total, total[len(total)-1]+tot)
				hold = append(hold, ho)
				fmt.Println()
			} else {
				fmt.Println("Incorrect arguments.")
			}
		case "place":
			if len(command) < 3 {
				fmt.Println("Insufficient amount of arguments.")
				break
			}
			items := command[1:]
			success, h, b, tot = place(items, h, b)
			if success {
				actions = append(actions, action)
				cpHand = append(cpHand, sortHand(h))
				cpBoard = append(cpBoard, b)
				total = append(total, total[len(total)-1]+tot)
				hold = append(hold, ho)
				fmt.Println()
			} else {
				fmt.Println("Incorrect arguments.")
			}
		case "undo":
			actions = actions[:len(actions)-1]
			cpHand = cpHand[:len(cpHand)-1]
			cpBoard = cpBoard[:len(cpBoard)-1]
			total = total[:len(total)-1]
			hold = hold[:len(hold)-1]
			fmt.Println("Undid last step.")
		case "restart":
			actions = actions[:0]
			cpHand = cpHand[:1]
			cpBoard = cpBoard[:1]
			total = total[:1]
			hold = hold[:1]
			fmt.Println("Restarted.")
		case "draw":
			if !isPoolEmpty() {
				draw()
				return true
			} else {
				fmt.Println("No more cards in the pool!")
			}
		case "done":
			if len(hold[len(hold)-1]) == 0 && len(cpHand) != 1 {
				if total[len(total)-1] < 30 && !laid[turn] {
					fmt.Println("Initial placements not done yet.")
					break
				} else {
					hands[turn] = cpHand[len(cpHand)-1]
					board = cpBoard[len(cpBoard)-1]
					laid[turn] = true
					return true
				}
			} else {
				fmt.Println("You're not permitted to finish yet.")
			}
		default:
			fmt.Println("Incorrect command, try again.")
		}
	}
}

func add(item string, to string, h []card, b [][]card) (bool, []card, [][]card, int) {
	var tot int
	c, fail := processItem(item)
	if fail == 1 {
		return false, h, b, tot
	}
	num, err := strconv.Atoi(to)
	if err != nil {
		return false, h, b, tot
	}
	num--
	if num < 0 || num >= len(b) {
		return false, h, b, tot
	}
	if isIn(c, h) == -1 {
		return false, h, b, tot
	}
	if isValid(append(b[num], c)) {
		h = remove(h, c)
		b[num] = append(b[num], c)
		tot = c.number
		return true, h, b, tot
	} else if isValid(append([]card{c}, b[num]...)) {
		h = remove(h, c)
		b[num] = append([]card{c}, b[num]...)
		tot = c.number
		return true, h, b, tot
	}
	return false, h, b, tot
}

func place(items []string, h []card, b [][]card) (bool, []card, [][]card, int) {
	var cs []card
	tot := 0
	printCards(h)
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

func isValid(b []card) bool {
	validRun := true
	validGroup := true
	if len(b) < 3 {
		return false
	}
	color := b[0].color
	count := b[0].number
	colors := []bool{false, false, false, false} // black, yellow, blue, red
	for i := 1; i < len(b); i++ {
		if b[i].color == color && b[i].number == count+1 {
			count++
		} else {
			validRun = false
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

func isIn(c card, h []card) int {
	for i, a := range h {
		if c == a || (a.color == -1 && c.joker != 0) {
			return i
		}
	}
	return -1
}

func processItem(item string) (card, int) {
	// Normal items:
	// b_11
	// Empty jokers:
	// jr
	// Signed jokers:
	// jr_y_8
	var c card
	c.joker = 0
	items := strings.Split(item, "_")
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
	case 4:
		c.joker = 1
		if l == 1 {
			c.color = -1
			c.number = 0
		} else if l == 3 {
			switch whatColor(items[1]) {
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
	case "jb", "jokerblack":
		return 4
	case "jr", "jokerred":
		return 5
	}
	return -1
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

func sortHand(h []card) []card {
	// Sorting algorithm
	// First by color, then by ascending number
	// Last cards are jokers
	sort.Slice(h, func(i, j int) bool {
		return h[i].joker != 0
	})
	sort.Slice(h, func(i, j int) bool {
		/*if h[i].joker != 0 || h[j].joker != 0 {
			return false
		}*/
		if h[i].color < h[j].color {
			return true
		} else if h[i].color > h[j].color {
			return false
		}
		return h[i].number < h[j].number
	})

	return h
}
