package main

import (
	"fmt"
)

type Board struct {
	numbers       [][]int
	active_number [][]bool
	bingo         bool
	final_number  int
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
	boards := split_input_into_boards(input_board)
	_, bingo_board := roll_numbers_and_check_boards(boards, input_numbers)
	fmt.Println("Final board: ")
	print_active_numbers(bingo_board)

	fmt.Println("\nAnswer: ")
}

func initialize_array(size_x int, size_y int, value bool) [][]bool {
	new_array := make([][]bool, size_x)
	for i := 0; i < size_x; i++ {
		new_array[i] = make([]bool, size_y)
	}
	return new_array
}

func print_2dArray(array [][]int) {
	for _, row := range array {
		fmt.Println(row)
	}
}

func print_active_numbers(board Board) {
	for i, row := range board.numbers {
		for j, num := range row {
			if board.active_number[i][j] == true {
				fmt.Print("X\t")
			} else {
				fmt.Print(num, "\t")
			}
		}
		fmt.Print("\n")
	}
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

func roll_numbers_and_check_boards(boards []Board, numbers_rolled []int) (bool, Board) {
	bingo := false
	var bingo_board Board
	for _, number := range numbers_rolled {
		bingo, bingo_board, boards = activate_number_on_all_boards(boards, number)
		if bingo {
			return true, bingo_board
		}
	}
	return false, bingo_board
}

func activate_number_on_all_boards(boards []Board, number_rolled int) (bool, Board, []Board) {
	bingo := false
	var bingo_board Board
	for _, board := range boards {
		bingo, board = activate_number_on_board(board, number_rolled)
		if bingo {
			return bingo, board, boards
		}
	}
	return bingo, bingo_board, boards
}

func activate_number_on_board(board Board, number_rolled int) (bool, Board) {
	for i, row := range board.numbers {
		for j, number := range row {
			if number == number_rolled {
				board.active_number[i][j] = true
				board.bingo = check_for_bingo(i, j, board)
				break
			}
		}
	}
	return board.bingo, board
}

func check_column_for_bingo(column int, board Board) bool {
	for _, row := range board.active_number {
		if !row[column] {
			return false
		}
	}
	return true
}

func check_row_for_bingo(row []bool) bool {
	for _, value := range row {
		if !value {
			return false
		}
	}
	return true
}

func check_for_bingo(row_number int, col_number int, board Board) bool {
	if check_column_for_bingo(col_number, board) {
		return true
	}
	if check_row_for_bingo(board.active_number[row_number]) {
		return true
	}
	return false
}
