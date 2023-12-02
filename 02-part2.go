package main

import (
	"fmt"
	"regexp"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var total int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		r1,_ := regexp.Compile("^Game ([0-9]+): (.*)$")
		tmp1 := r1.FindStringSubmatch(line)
		line = tmp1[2]
		data := strings.Split(line,";")
		r2, _ := regexp.Compile("([0-9]+) (red|blue|green)")
		mymax := make(map[string]int)
		mymax["red"] = 0
		mymax["green"] = 0
		mymax["blue"] = 0
		for i:=0; i < len(data); i++ {
			b := r2.FindAllStringSubmatch(data[i],-1)
			for j:=0; j < len(b); j++ {
				tmp2,_ := strconv.Atoi(b[j][1])
				if tmp2 > mymax[b[j][2]] {
					mymax[b[j][2]] = tmp2
				}
			}
		}
		total += mymax["red"] * mymax["green"] * mymax["blue"]
	}
	fmt.Println(total)
}