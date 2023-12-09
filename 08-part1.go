package main

import (
	"fmt"
	"bufio"
	"os"
)

func read_input() (string, map[string][2]string) {
	mymap := make(map[string][2]string)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	mypath := scanner.Text()
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			mymap[line[0:3]] = [2]string{line[7:10],line[12:15]}
		}
	}
	return mypath, mymap
}

func find_one_round_result(position string, mypath string, mymap map[string][2]string) string {
	for _,direction := range mypath {
		var shift byte
		if direction == 'L' {
			shift = 0
		} else {
			shift = 1
		}
		position = mymap[position][shift]
	}
	return position
}

func main() {
	mypath, mymap := read_input()
	mymap2 := make(map[string]string)
	for key,_ := range mymap {
		mymap2[key] = find_one_round_result(key,mypath,mymap)
	}
	counter := 0
	position := "AAA"
	for true {
		counter++
		position = mymap2[position]
		if position == "ZZZ" { break }
	}
	counter = counter * len(mypath)
	fmt.Println(counter)
}