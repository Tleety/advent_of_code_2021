package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 3 - Task 1")
	fmt.Println("https://adventofcode.com/2021/day/3")
	fmt.Println("Instructions: The submarine has been making some creaking noices. The iagnostics are written in binary. Try to decode them to find out what has happened.")
	inputs := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	fmt.Println(inputs)

	fmt.Println("\nAnswer: ")
}

func convert_numberstring_to_int_array(input string) []int {
	input_as_array := strings.Split(input, "")
	int_array := []int{}
	for _, digit_as_string := range input_as_array {
		digit_as_int, err := strconv.Atoi(digit_as_string)
		if err != nil {
			panic(err)
		}
		int_array = append(int_array, digit_as_int)
	}
	return int_array
}

func convert_numberstring_array_to_2d_int_array(array_of_stringnumbers []string) [][]int {
	var converted_input_to_int_2d_array [][]int
	for _, number_as_string := range array_of_stringnumbers {
		converted_input_to_int_2d_array = append(converted_input_to_int_2d_array, convert_numberstring_to_int_array(number_as_string))
	}
	return converted_input_to_int_2d_array
}
