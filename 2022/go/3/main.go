package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

const filename = "input.txt"

func main() {
	puzzle1()
}

func puzzle1() {
	file, err := os.Open(filename)
	if err != nil {
		println("Error opening %s -> %s", filename, err)
	}

	var (
		totalPriority int = 0
	)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		// currPriority := 0
		// currBag := sc.Text()
		// halfBag := len(sc.Text())

		items := make(map[rune]bool)
		for _, left := range sc.Text()[:len(sc.Text())/2] {
			items[left] = true
		}

		for _, right := range sc.Text()[len(sc.Text())/2:] {
			if items[right] {
				totalPriority += int(unicode.ToLower(right) - 96)
				if unicode.IsUpper(right) {
					totalPriority += 26
				}
				break
			}
		}
	}
	fmt.Println(totalPriority)
}

// Day 3 Puzzle 2
func puzzle2() {
	input, _ := os.Open(filename)
	defer input.Close()
	sc := bufio.NewScanner(input)

	var totalPriorities int = 0

	for sc.Scan() {
		first := createSetOfItems(sc.Text())
		sc.Scan()
		second := createSetOfItems(sc.Text())
		sc.Scan()
		third := createSetOfItems(sc.Text())

		for item := range first {
			if second[item] && third[item] {
				totalPriorities += int(unicode.ToLower(item) - 96)
				if unicode.IsUpper(item) {
					totalPriorities += 26
				}
				break
			}
		}
	}
	fmt.Println(totalPriorities)
}

func createSetOfItems(items string) (set map[rune]bool) {
	set = make(map[rune]bool)
	for _, item := range items {
		set[item] = true
	}
	return
}

