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
	fmt.Println(calculate_oxygen_rating(convert_numberstring_array_to_2d_int_array(inputs)))

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

func transpose(input_as_2d_array [][]int) [][]int {
	out := make([][]int, len(input_as_2d_array[0]))
	for _, input_array := range input_as_2d_array {
		for j, value := range input_array {
			out[j] = append(out[j], value)
		}
	}
	return out
}

func return_most_common_value(input []int) int {
	main_value := 0
	for _, number := range input {
		if number == 0 {
			main_value -= 1
		} else {
			main_value += 1
		}
	}
	if main_value > 0 {
		return 1
	} else if main_value < 0 {
		return 0
	}
	return -1
}

func calculate_main_number_in_array(input_as_transposed_2d_int [][]int, primary_num int) []int {
	var main_numbers []int
	for _, numbers := range input_as_transposed_2d_int {
		num := return_most_common_value(numbers)
		if num == -1 {
			num = primary_num
		}
		main_numbers = append(main_numbers, num)
	}
	return main_numbers
}

func convert_int_array_to_string(input_as_number_array []int) string {
	var number_string string
	for _, value := range input_as_number_array {
		number_string += strconv.Itoa(value)
	}
	return number_string
}

func inverse_int_array(input []int) []int {
	var output []int
	for _, value := range input {
		if value == 0 {
			output = append(output, 1)
		} else {
			output = append(output, 0)
		}
	}
	return output
}

func calculate_oxygen_rating(input_as_2d_int [][]int) []int {
	final_array := input_as_2d_int
	for i := 0; i < len(input_as_2d_int[0]); i += 1 {
		value_to_keep := calculate_main_number_in_array(transpose(final_array), 1)[i]
		final_array = exclude_unwanted_array(final_array, i, value_to_keep)
		if len(final_array) == 1 {
			return final_array[0]
		}
	}
	return final_array[0]
}

func exclude_unwanted_array(input [][]int, position_to_look_at int, value_we_want int) [][]int {
	var wanted_arrays [][]int
	for _, array := range input {
		if array[position_to_look_at] == value_we_want {
			wanted_arrays = append(wanted_arrays, array)
		}
	}
	return wanted_arrays
}