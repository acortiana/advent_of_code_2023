package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func atoi(textnumber string) int {
	tmpvalue, _:= strconv.Atoi(textnumber)
	return tmpvalue
}

func reduce_slice(input []int) []int {
	var result []int
	for i:=0; i < len(input)-1; i++ {
		result = append(result,input[i+1]-input[i])
	}
	return result
}

func all_slice_elements_equal(input []int) bool {
	equal := true
	for i:=0; i < len(input); i++ {
		if input[i] != input[0] {
			equal = false
			break
		}
	}
	return equal
}

func get_slice_next_value(input []int) int {
	var mydata [][]int
	mydata = append(mydata,input)
	for ; !all_slice_elements_equal(mydata[len(mydata)-1]); {
		mydata = append(mydata,reduce_slice(mydata[len(mydata)-1]))
	}
	myvalue := mydata[len(mydata)-1][0]
	for i := len(mydata)-1; i > 0; i-- {
		myvalue = myvalue + mydata[i-1][len(mydata[i-1])-1]
	}
	return myvalue
}

func read_input() ([][]int) {
	var data [][]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var tmp []int
		for _,value := range strings.Fields(line) {
			tmp = append(tmp,atoi(value))
		}
		data = append(data,tmp)
	}
	return data
}

func main() {
	data := read_input()
	sum := 0
	for _, value := range data {
		sum += get_slice_next_value(value)
	}
	fmt.Println(sum)
}