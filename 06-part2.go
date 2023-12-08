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

func read_input() ([][]int) {
	var result_a [][]int
	var result_b [][]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var tmp_a []int
		tmp_a = append(tmp_a,atoi(strings.Join(strings.Fields(line)[1:],"")))
		result_a = append(result_a,tmp_a)
	}
	for i := 0; i < len(result_a[0]); i++ {
		result_b = append(result_b,[]int{result_a[0][i],result_a[1][i]})
	}
	return result_b
}

func better_races_count(race *[]int) int {
	ms := (*race)[0]
	best := (*race)[1]
	mybestcount := 0
	for i:=1; i < ms; i++ {
		result := (ms - i) * i
		if result > best {
			mybestcount++
		}
	}
	return mybestcount
}

func main() {
	races := read_input()
	result := 1
	for _, race := range races {
		result *= better_races_count(&race)
	}
	fmt.Println(result)
}