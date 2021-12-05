package main

import (
	"fmt"
)

type Board struct {
	numbers       [][]int
	active_number [][]bool
	score         int
}

func main() {
	fmt.Println("Day 4 - Task 1")
	fmt.Println("https://adventofcode.com/2021/day/4")
	fmt.Println("Instructions: Find out where the submarine ends up by following its movement log.")

	input_numbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	input_board := [][]int{{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
		{3, 15, 0, 2, 22},
		{9, 18, 13, 17, 5},
		{19, 8, 7, 25, 23},
		{20, 11, 10, 24, 4},
		{14, 21, 16, 12, 6},
		{14, 21, 17, 24, 4},
		{10, 16, 15, 9, 19},
		{18, 8, 23, 26, 20},
		{22, 11, 13, 6, 5},
		{2, 0, 12, 3, 7},
	}
	fmt.Println("input", input_numbers)
	fmt.Println("Boards:", input_board)
	split_input_into_boards(input_board)

	fmt.Println("\nAnswer: ")
}

func initialize_array(size_x int, size_y int, value bool) [][]bool {
	new_array := make([][]bool, size_x)
	for i := 0; i < size_x; i++ {
		new_array[i] = make([]bool, size_y)
	}
	return new_array
}

func split_input_into_boards(input [][]int) []Board {
	var boards []Board

	var board_length = len(input[0])
	for i := 0; i < len(input); {
		var new_board Board
		new_board.numbers = input[i : i+board_length]
		new_board.active_number = initialize_array(board_length, board_length, false)
		i = i + board_length
		boards = append(boards, new_board)
	}
	return boards
}
