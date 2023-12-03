package main

import (
	"fmt"
	"bufio"
	"os"
	"slices"
	"strconv"
)

func adjacent_star(stars_table [][2]int, coords [2]int) [2]int {
	x := coords[0]
	y := coords[1]
	for i:=-1; i <= 1; i++ {
		cur_x := x+i
		if cur_x < 0 { continue }
		for j:=-1; j <= 1; j++ {
			cur_y := y+j
			if cur_y < 0 { continue }
			if slices.Contains(stars_table, [2]int{cur_x,cur_y}) { return [2]int{cur_x,cur_y}}
		}
	}
	return [2]int{-1,-1}
}

func is_digit(mychar byte) bool {
	if mychar >= '0' && mychar <= '9' { return true }
	return false
}

func is_star(mychar byte) bool {
	if mychar == '*' { return true }
	return false
}

func build_stars_table(textslice [][]byte) [][2]int {
	var char_coords [][2]int
	for lineid, myline := range textslice {
		for charid, mychar := range myline {
			if is_star(mychar) {
				tmp := [2]int{lineid, charid}
				char_coords = append(char_coords,tmp)
			}
		}
	}
	return char_coords
}

func build_gears_datastructure(textslice [][]byte,stars_table [][2]int) map[[2]int][]int {
	data := make(map[[2]int][]int)
	found_adjacent_star := false
	last_star := [2]int{-1,-1}
	text_number := ""
	for lineid, myline := range textslice {
		for charid, mychar := range myline {
			if is_digit(mychar) {
				text_number = text_number + string(mychar)
				mystar := adjacent_star(stars_table,[2]int{lineid,charid})
				if  mystar != [2]int{-1,-1} {
					found_adjacent_star = true
					last_star = mystar
				}
			} else {
				if found_adjacent_star && len(text_number) > 0 {
					tmp, _ := strconv.Atoi(text_number)
					data[last_star] = append(data[last_star],tmp)
				}
				text_number = ""
				found_adjacent_star = false
				last_star = [2]int{-1,-1}
			}
		}
		if found_adjacent_star && len(text_number) > 0 {
			tmp, _ := strconv.Atoi(text_number)
			data[last_star] = append(data[last_star],tmp)
		}
		text_number = ""
		found_adjacent_star = false
		last_star = [2]int{-1,-1}
	}
	return data
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var textslice [][]byte
	for scanner.Scan() {
		textslice = append(textslice,[]byte(scanner.Text()))
	}
	stars_table := build_stars_table(textslice)
	gears_datastructure := build_gears_datastructure(textslice,stars_table)
	sum := 0
	for _, value := range gears_datastructure {
		if len(value) == 2 {
			sum += value[0] * value[1]
		}
	}
	fmt.Println(sum)
}