package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strings"
)

func calculate_score (a int) int {
	if a == 0 { return 0 }
	return 1 << (a-1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		r1,_ := regexp.Compile("^Card +([0-9]+): ([0-9 ]+) . ([0-9 ]+)$")
		tmp1 := r1.FindStringSubmatch(line)
		mymap := make(map[string]int)
		winner_strings := strings.Fields(tmp1[2])
		my_strings := strings.Fields(tmp1[3])
		total_strings := append(winner_strings,my_strings...)
		tmp := 0
		for _,value := range total_strings {
			mymap[value]++
		}
		for _,value := range mymap {
			if value == 2 {
				tmp++
			}
		}
		total += calculate_score(tmp)
	}
	fmt.Println(total)
}