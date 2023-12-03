package main

import (
	"fmt"
	"bufio"
	"os"
	"slices"
	"strconv"
)

func adjacent_symbol(symbols_table [][2]int, coords [2]int) bool {
	x := coords[0]
	y := coords[1]
	for i:=-1; i <= 1; i++ {
		cur_x := x+i
		if cur_x < 0 { continue }
		for j:=-1; j <= 1; j++ {
			cur_y := y+j
			if cur_y < 0 { continue }
			if slices.Contains(symbols_table, [2]int{cur_x,cur_y}) { return true }
		}
	}
	return false
}

func is_digit(mychar byte) bool {
	if mychar >= '0' && mychar <= '9' { return true }
	return false
}

func is_symbol(mychar byte) bool {
	if is_digit(mychar) { return false }
	if mychar == '.' { return false }
	return true
}

func build_symbols_table(textslice [][]byte) [][2]int {
	var char_coords [][2]int
	for lineid, myline := range textslice {
		for charid, mychar := range myline {
			if is_symbol(mychar) {
				tmp := [2]int{lineid, charid}
				char_coords = append(char_coords,tmp)
			}
		}
	}
	return char_coords
}

func find_part_numbers(textslice [][]byte,symbols_table [][2]int) []int {
	var resultslice []int
	found_adjacent_symbol := false
	text_number := ""
	for lineid, myline := range textslice {
		for charid, mychar := range myline {
			if is_digit(mychar) {
				text_number = text_number + string(mychar)
				if adjacent_symbol(symbols_table,[2]int{lineid,charid}) {
					found_adjacent_symbol = true
				}
			} else {
				if found_adjacent_symbol && len(text_number) > 0 {
					tmp, _ := strconv.Atoi(text_number)
					resultslice = append(resultslice,tmp)
				}
				text_number = ""
				found_adjacent_symbol = false
			}
		}
		if found_adjacent_symbol && len(text_number) > 0 {
			tmp, _ := strconv.Atoi(text_number)
			resultslice = append(resultslice,tmp)
		}
		text_number = ""
		found_adjacent_symbol = false
	}
	return resultslice
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var textslice [][]byte
	for scanner.Scan() {
		textslice = append(textslice,[]byte(scanner.Text()))
	}
	symbols_table := build_symbols_table(textslice)
	part_numbers := find_part_numbers(textslice,symbols_table)
	sum := 0
	for _, i := range part_numbers {
		sum += i
	}
	fmt.Println(sum)
}