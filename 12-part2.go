package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func atoi(textnumber string) int {
	tmpvalue, _:= strconv.Atoi(textnumber)
	return tmpvalue
}

func read_input() ([][]int, [][]int) {
	var data1 [][]int
	var data2 [][]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		a := strings.Fields(line)
		b := strings.Split(a[1],",")
		var c []int
		for _, char := range b {
			c = append(c,atoi(char))
		}
		data1 = append(data1,transform_input(a[0]))
		data2 = append(data2,c)
	}
	return data1, data2
}

func add_leading_trailing_zeroes(input [][]int) [][]int {
	var output [][]int
	for _, value := range input {
		tmp := append([]int{0}, value...)
		tmp = append(tmp, []int{0}...)
		output = append(output,tmp)
	}
	return output
}

func transform_input(input string) []int {
	var m = map[byte]int{
		'.': 0,
		'#': 1,
		'?': 2,
	}
	var output []int
	for _, value := range input {
		output = append(output,m[byte(value)])
	}
	return output
}

func compare_patterns(input1 []int, input2 []int) bool {
	if len(input1) != len(input2) { return false }
	for i:=0; i < len(input1); i++ {
		if input1[i] != input2[i] { return false }
	}
	return true
}

func is_possible(input []int) bool {
	for i:=0; i < (len(input) - 1); i++ {
		if input[i] == 0 { return false }
	}
	if input[len(input)-1] == 1 { return false }
	return true
}

func more_sharps(input []int, index int) bool {
	for i:=index; i < len(input); i++ {
		if input[i] == 1 { return true }
	}
	return false
}

func get_i_max(input []int, pattern []int, patternindex int) int {
	result := 0
	for i:=patternindex; i < len(pattern); i++ {
		result += pattern[i] + 1
	}
	result = len(input) - result
	return result
}

func analyze(input []int, pattern []int, inputindex int, patternindex int, analyze_cache map[[2]int]int) int {
	tmp, ok := analyze_cache[[2]int{inputindex,patternindex}]
	if ok { return tmp }
	var result int
	pattern_number := pattern[patternindex]
	i_max := get_i_max(input,pattern,patternindex)
	for i:=inputindex; i <= i_max ; i++ {
		if i > 0 && i < (len(input) - 1) && input[i-1] == 1 { break }
		if (i+pattern_number) >= len(input) { break }
		if is_possible(input[i:i+pattern_number+1]) {
			if patternindex == len(pattern) - 1 {
				if more_sharps(input,i+pattern_number+1) { continue }
				result++
			} else {
				result += analyze(input,pattern,i+pattern_number+1,patternindex+1,analyze_cache)
			}
		}
	}
	analyze_cache[[2]int{inputindex,patternindex}] = result
	return result
}

func slice_x_5 (input [][]int) [][]int {
	var output [][]int
	for i:=0; i < len(input); i++ {
		output = append(output,[]int{})
		output[i] = append(output[i],input[i]...)
		output[i] = append(output[i],input[i]...)
		output[i] = append(output[i],input[i]...)
		output[i] = append(output[i],input[i]...)
		output[i] = append(output[i],input[i]...)
	}
	return output
}

func slice_x_5_and_question_mark (input [][]int) [][]int {
	var output [][]int
	for i:=0; i < len(input); i++ {
		output = append(output,[]int{})
		output[i] = append(output[i],input[i]...)
		output[i] = append(output[i],2)
		output[i] = append(output[i],input[i]...)
		output[i] = append(output[i],2)
		output[i] = append(output[i],input[i]...)
		output[i] = append(output[i],2)
		output[i] = append(output[i],input[i]...)
		output[i] = append(output[i],2)
		output[i] = append(output[i],input[i]...)
	}
	return output
}

func main() {
	data1, data2 := read_input()
	data1 = slice_x_5_and_question_mark(data1)
	data2 = slice_x_5(data2)
	data1 = add_leading_trailing_zeroes(data1)
	var sum int
	for i:=0; i < len(data1); i++ {
		analyze_cache := make(map[[2]int]int)
		sum += analyze(data1[i],data2[i],0,0,analyze_cache)
	}
	fmt.Println(sum)
}