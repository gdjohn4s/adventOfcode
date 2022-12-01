package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const filename string = "input.txt"

func main() {
	elves1()
	elves2()
}

/*
PUZZLE 1 -> Elves max calories
*/
func elves1() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: ", err)
	}
	defer file.Close()

	maxCalories := 0
	currentCalories := 0

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		snack, err := strconv.Atoi(sc.Text())
		currentCalories += snack

		if err != nil {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
		}
	}
	fmt.Println(maxCalories)
}

/*
PUZZLE 2 -> Backup calories
*/
func elves2() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file")
	}
	defer file.Close()

	maxCal1, maxCal2, maxCal3, curCal := 0, 0, 0, 0

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		max, err := strconv.Atoi(sc.Text())
		curCal += max

		if err != nil {
			if curCal > maxCal3 {
				maxCal3 = curCal
			} else if maxCal3 > maxCal2 {
				maxCal3, maxCal2 = maxCal2, maxCal3
			} else if maxCal2 > maxCal1 {
				maxCal2, maxCal1 = maxCal1, maxCal2
			}
			curCal = 0
		}
	}
	fmt.Println(maxCal1 + maxCal2 + maxCal3)
}

