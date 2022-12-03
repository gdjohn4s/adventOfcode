package main

// #f20e7e

import (
	"bufio"
	"fmt"
	"os"
)

const filename = "input.txt"

func main() {
	puzzle1()
}

// Day 2 Puzzle 1
func puzzle1() {
	file, err := os.Open(filename)
	if err != nil {
		println("Error opening %s -> %s", filename, err)
	}

	sc := bufio.NewScanner(file)

	var (
		totalScore int = 0
	)

	scores := map[string]int{
		"B X": 1,
		"C Y": 2,
		"A Z": 3,
		"A X": 4,
		"B Y": 5,
		"C Z": 6,
		"C X": 7,
		"A Y": 8,
		"B Z": 9,
	}

	for sc.Scan() {
		totalScore += scores[sc.Text()]
	}
	fmt.Println(totalScore)
}

