package main

import (
	"fmt"
)

type Position struct {
	Horizontal, Depth int
}
type Movement struct {
	move  string
	value int
}

func main() {
	fmt.Println("Day 2 - Task 1")
	fmt.Println("https://adventofcode.com/2021/day/2")
	fmt.Println("Instructions: Find out where the submarine ends up by following its movement log.")
	inputs := []Movement{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}

	fmt.Println("inputs: ", inputs)
}

func move_submarine(position Position, input Movement) Position {
	switch input.move {
	case "forward":
		position.Horizontal += input.value
	case "down":
		position.Depth += input.value
	case "up":
		position.Depth -= input.value
	}
	return position
}

func execute_submarine_movement_log(position Position, movement_log []Movement) Position {
	for _, movement := range movement_log {
		position = move_submarine(position, movement)
	}
	return position
}
