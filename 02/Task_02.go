package main

import (
	"fmt"
)

type Position struct {
	Horizontal, Depth, Aim int
}
type Movement struct {
	move  string
	value int
}

func main() {
	fmt.Println("Day 2 - Task 2")
	fmt.Println("https://adventofcode.com/2021/day/2#part2")
	fmt.Println("Instructions: Find out where the submarine ends up by following its movement log. Instead of depth changing by value, down and up moves an aim, and the depth changes by forward*aim when moving forward.")
	inputs := []Movement{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}
	var submarine_position Position
	submarine_position = execute_submarine_movement_log(submarine_position, inputs)
	fmt.Println("\nAnswer: ")
	fmt.Println("Final position is a horizontal position of", submarine_position.Horizontal, "and depth of", submarine_position.Depth)
	fmt.Println("To get the final answer multiply these two values together.")
	fmt.Println("Final answer:", submarine_position.Depth, "*", submarine_position.Horizontal, "=", submarine_position.Depth*submarine_position.Horizontal)
}

func move_submarine(position Position, input Movement) Position {
	switch input.move {
	case "forward":
		position.Horizontal += input.value
		position.Depth += input.value * position.Aim
	case "down":
		position.Aim += input.value
	case "up":
		position.Aim -= input.value
	}
	return position
}

func execute_submarine_movement_log(position Position, movement_log []Movement) Position {
	for _, movement := range movement_log {
		position = move_submarine(position, movement)
	}
	return position
}
