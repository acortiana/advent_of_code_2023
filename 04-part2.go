package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strings"
)

func main() {
	var card_matches []int
	scanner := bufio.NewScanner(os.Stdin)
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
		card_matches = append(card_matches,tmp)
	}
	var card_count []int
	for i:=0; i < len(card_matches); i++ {
		card_count = append(card_count,1)
	}
	for i:=0; i < len(card_matches); i++ {
		for j:=1; j <= card_matches[i]; j++ {
			if (i+j) < len(card_matches) {
				card_count[i+j] += card_count[i]
			}
		}
	}
	var total int
	for _,value := range card_count {
		total += value
	}
	fmt.Println(total)
}