package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var total uint64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var firstdigit, lastdigit uint8;
		for i:=0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				mydigit:= (line[i] - '0')
				if firstdigit == 0 { firstdigit = mydigit }
				lastdigit = mydigit
			}
		}
		total += uint64(firstdigit * 10 + lastdigit)
	}
	fmt.Printf("%d\n",total)
}