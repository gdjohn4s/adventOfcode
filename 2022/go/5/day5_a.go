package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename string = "input.prod"

func main() {
	puzzle1()
	// puzzle2()
}

type stack struct {
	crates []rune
}

func (s *stack) Push(c rune) {
	s.crates = append(s.crates, c)
}

func (s *stack) Pop() (c rune) {
	c = s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]
	return
}

func (s *stack) bottom(c rune) {
	s.crates = append([]rune{c}, s.crates...)
}

func (s *stack) parse() string {
	var str string
	for _, c := range s.crates {
		str += string(c) + " "
	}
	return str
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

/* Day 4 Puzzle 1 */
func puzzle1() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	// slice of stacks
	stacks := make([]stack, 9)
	fmt.Println(stacks)

	stacks = parseInput(sc, stacks)

	sc.Scan()
	for sc.Scan() {
		var move, from, to int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &move, &from, &to)
		fmt.Println(move, from, to)

		for i := 0; i < move; i++ {
			stacks[to-1].Push(stacks[from-1].Pop())
		}
		fmt.Println(stacks)
	}

	for _, s := range stacks {
		fmt.Print(string(s.Pop()))
	}

}

