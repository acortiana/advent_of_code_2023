package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func get_digit_from_text(input string) uint8 {
	array := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i:=0; i < len(array); i++ {
		if strings.HasPrefix(input,array[i]) { return uint8(i+1) }
	}
	return 0
}

func main() {
	var total uint64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var firstdigit, lastdigit uint8;
		for i:=0; i < len(line); i++ {
			var mydigit uint8
			tmp := get_digit_from_text(line[i:])
			if tmp > 0 {
				mydigit = tmp
			} else if line[i] >= '0' && line[i] <= '9' {
				mydigit = uint8((line[i] - '0'))
			} else {
				continue
			}
			if firstdigit == 0 { firstdigit = mydigit }
			lastdigit = mydigit
		}
		total += uint64(firstdigit * 10 + lastdigit)
	}
	fmt.Printf("%d\n",total)
}