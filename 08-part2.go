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
	var mydata [][]int
	for key,_ := range mymap2 {
		if key[2] != 'A' { continue }
		var result []int
		movecount := 0
		for ; len(result) <2 ; {
			movecount++
			key = mymap2[key]
			if key[2] == 'Z' { result = append(result,movecount) }
		}
		result = append(result,result[1]-result[0])
		mydata = append(mydata, result)
	}

	for i:=0; ; i++ {
		var success bool
		for k:=0; k < len(mydata); k++ {
			success = true
			value := mydata[k][0] + mydata[k][2]*i
			for j:=0; j < len(mydata); j++ {
				if !(value % mydata[j][2] == (mydata[j][0] - (mydata[j][2]))) {
					success = false
					break
				}
			}
			if success == true {
				fmt.Println(value*len(mypath))
				break
			}
		}
		if success == true { break }
	}
}