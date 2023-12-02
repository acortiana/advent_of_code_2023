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
	var possible int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		r1,_ := regexp.Compile("^Game ([0-9]+): (.*)$")
		tmp1 := r1.FindStringSubmatch(line)
		gameid,_ := strconv.Atoi(tmp1[1])
		line = tmp1[2]
		data := strings.Split(line,";")
		r2, _ := regexp.Compile("([0-9]+) (red|blue|green)")
		var fail bool = false
		count := make(map[string]int)
		for i:=0; i < len(data); i++ {
			count["red"] = 0
			count["green"] = 0
			count["blue"] = 0
			b := r2.FindAllStringSubmatch(data[i],-1)
			for j:=0; j < len(b); j++ {
				tmp2,_ := strconv.Atoi(b[j][1])
				count[b[j][2]] += tmp2
			}
			if count["red"] > 12 || count["green"] > 13 || count["blue"] > 14 {
				fail = true
				break
			}
		}
		if fail == false { possible += gameid }
	}
	fmt.Println(possible)
}