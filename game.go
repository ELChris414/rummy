package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/inancgumus/screen"
)

type card struct {
	color  int // 0 black, 1 yellow, 2 blue, 3 red, 4 emptyJokerB, 5 emptyJokerR
	number int // 0 emptyJoker, 1-13 typical numbers
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
		won = hasWon()
	}
	fmt.Printf("Player %s has won the game!\n", playerLet[won])
	fmt.Printf("Game finished in %v rounds\n", rounds)
	fmt.Println("Final board:")
	renderBoard(board)
	fmt.Println("\nScores:")
	scores := calculateScores(won)
	renderScores(scores)
}

func playTurn() bool {
	var cpHand [][]card
	var cpBoard [][][]card
	var hold [][]card = [][]card{{}}
	var actions []string
	var total []int = []int{0}
	var tot int

	cpBoard = append(cpBoard, board)
	cpHand = append(cpHand, hands[turn])

	for {
		//screen.Clear()
		fmt.Printf("Currently playing: Player %s\n", playerLet[turn])
		fmt.Printf("Cards on board: %v\n", len(pool))
		fmt.Printf("Round: %v\n\n", rounds)
		renderBoard(cpBoard[len(cpBoard)-1])
		renderActions(actions, total[len(total)-1])
		renderHand(cpHand[len(cpHand)-1])
		renderHold(hold[len(hold)-1])
		fmt.Print("Let's play! Give commands. Possible commands:  place")
		if total[len(total)-1] >= 30 || laid[turn] {
			fmt.Print("  add  pick  exchange  pickall  break")
		}
		if len(cpHand) > 1 {
			fmt.Print("  restart  undo")
		}
		if len(hold[len(hold)-1]) == 0 && len(cpHand) != 1 && !(total[len(total)-1] < 30 && !laid[turn]) {
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

		h := make([]card, len(cpHand[len(cpHand)-1]))
		b := make([][]card, len(cpBoard[len(cpBoard)-1]))
		ho := make([]card, len(hold[len(hold)-1]))

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
			if len(command) < 3 {
				fmt.Println("Insufficient amount of arguments.")
				break
			}
			items := command[2:]
			success, h, b, ho = add(items, command[1], h, b, ho)
			if success {
				actions = append(actions, action)
				cpHand = append(cpHand, sortHand(h))
				cpBoard = append(cpBoard, b)
				total = append(total, total[len(total)-1])
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
			success, h, b, ho, tot = place(items, h, b, ho)
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
		case "exchange":
			if total[len(total)-1] < 30 && !laid[turn] {
				fmt.Println("Initial placements not done yet.")
				break
			}
			if len(command) != 3 {
				fmt.Println("Insufficient amount of arguments.")
				break
			}
			success, h, b, ho = exchange(command[1], command[2], h, b, ho)
			if success {
				actions = append(actions, action)
				cpHand = append(cpHand, sortHand(h))
				cpBoard = append(cpBoard, b)
				total = append(total, total[len(total)-1])
				hold = append(hold, ho)
				fmt.Println()
			} else {
				fmt.Println("Incorrect arguments.")
			}
		case "pick":
			if total[len(total)-1] < 30 && !laid[turn] {
				fmt.Println("Initial placements not done yet.")
				break
			}
			if len(command) < 3 {
				fmt.Println("Insufficient amount of arguments.")
				break
			}
			items := command[2:]
			success, b, ho = pick(items, command[1], b, ho)
			if success {
				actions = append(actions, action)
				cpHand = append(cpHand, h)
				cpBoard = append(cpBoard, b)
				total = append(total, total[len(total)-1])
				hold = append(hold, ho)
				fmt.Println()
			} else {
				fmt.Println("Incorrect arguments.")
			}
		default:
			fmt.Println("Incorrect command, try again.")
		}
	}
}
