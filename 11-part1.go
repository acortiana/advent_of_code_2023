package main

import (
	"fmt"
	"bufio"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func read_input() ([][]byte) {
	var data [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data,[]byte(line))
	}
	return data
}

func expand_vertical(data [][]byte) [][]byte {
	var result [][]byte
	for _, line := range data {
		expand := true
		for _, char := range line {
			if char != '.' {
				expand = false
				break
			}
		}
		if expand {
			result = append(result, line)
		}
		result = append(result, line)
	}
	return result
}

func expand_horizontal(data [][]byte) [][]byte {
	var result [][]byte
	for i:=0; i < len(data); i++ {
		empty := []byte{}
		result = append(result,empty)
	}

	for i:=0; i < len(data[0]); i++ {
		expand := true
		for j:=0; j < len(data); j++ {
			if data[j][i] != '.' {
				expand = false
				break
			}
		}
		for j:=0; j < len(data); j++ {
			if expand {
				result[j] = append(result[j],data[j][i])
			}
			result[j] = append(result[j],data[j][i])
		}
	}
	return result
}

func get_couple_identifier(point1 [2]int, point2 [2]int) [4]int {
	var result [4]int
	if point1[0] > point2[0] || (point1[0] == point2[0] && point1[1] > point2[1]) {
		result[0] = point1[0]
		result[1] = point1[1]
		result[2] = point2[0]
		result[3] = point2[1]
	} else {
		result[0] = point2[0]
		result[1] = point2[1]
		result[2] = point1[0]
		result[3] = point1[1]	
	}
	return result
}

func get_galaxies(data [][]byte) [][2]int {
	var result [][2]int
	for i:=0; i < len(data); i++{
		for j:=0; j < len(data[0]); j++ {
			if data[i][j] == '#' {
				result = append(result,[2]int{j,i})
			}
		}
	}
	return result
}

func get_galaxies_distance(galaxy1 [2]int, galaxy2 [2]int) int {
	return abs(galaxy1[0] - galaxy2[0]) + abs(galaxy1[1] - galaxy2[1])
}

func main() {
	data := read_input()
	data = expand_horizontal(expand_vertical(data))
	values := make(map[[4]int]int)
	for _, galaxy1 := range get_galaxies(data) {
		for _, galaxy2 := range get_galaxies(data) {
			if galaxy1 == galaxy2 { continue }
			couple_identifier := get_couple_identifier(galaxy1,galaxy2)
			if _, ok := values[couple_identifier]; !ok {
				values[couple_identifier] = get_galaxies_distance(galaxy1,galaxy2)
			}
		}
	}
	sum := 0
	for _, value := range values {
		sum += value
	}	
	fmt.Println(sum)

}