package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename string = "input.prod"

func main() {
	//puzzle1()
	puzzle2()
}

type stack struct {
	crates []rune
}

func (s *stack) Push(c []rune) {
	s.crates = append(s.crates, c...)
}

func (s *stack) Pop(n int) (c []rune) {
	c = s.crates[len(s.crates)-n : len(s.crates)]
	s.crates = s.crates[:len(s.crates)-n]
	return
}

func (s *stack) bottom(c rune) {
	s.crates = append([]rune{c}, s.crates...)
}

func parseInput(sc *bufio.Scanner, stack []stack) []stack {
	sc.Scan()
	for sc.Text() != " 1   2   3   4   5   6   7   8   9 " {
		for i, r := range sc.Text() {
			if r != ' ' && r != '[' && r != ']' {
				stack[i/4].bottom(r)
			}
		}
		sc.Scan()
	}
	return stack
}

/* Day 5 puzzle 2 */
func puzzle2() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	// slice of stacks
	stacks := make([]stack, 9)
	stacks = parseInput(sc, stacks)

	sc.Scan()
	for sc.Scan() {
		var move, from, to int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &move, &from, &to)
		stacks[to-1].Push(stacks[from-1].Pop(move))
	}

	for _, l := range stacks {
		fmt.Print(string(l.Pop(1)))
	}
}

