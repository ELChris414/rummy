package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	scanner *bufio.Scanner
)

func main() {
	fmt.Println("Welcome to Rummy!")
	scanner = bufio.NewScanner(os.Stdin)
	initializeGame()
	startGame()
}
