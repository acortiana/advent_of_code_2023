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

func get_galaxies_distance(galaxy1 [2]int, galaxy2 [2]int, galaxies [][2]int) int {
	result := abs(galaxy1[0] - galaxy2[0]) + abs(galaxy1[1] - galaxy2[1])
	coords := [2]int{galaxy2[0],galaxy2[1]}
	for i:=0; i < 2; i++ {
		inc := 1
		if galaxy2[i] - galaxy1[i] > 0 { inc = -1 }
		if galaxy2[i] == galaxy1[i] { continue }
		for true {
			coords[i] += inc
			if coords[i] == galaxy1[i] { break }
			expand_x, expand_y := is_expanded(coords,galaxies)
			if i == 0 && expand_x { result += 999999 }
			if i == 1 && expand_y { result += 999999 }
		}
	}
	return result
}

func is_expanded(point [2]int, galaxies [][2]int) (bool, bool) {
	expanded_x := true
	expanded_y := true
	for _, galaxy := range galaxies {
		if galaxy[0] == point[0] { expanded_x = false}
		if galaxy[1] == point[1] { expanded_y = false}
	}
	return expanded_x,expanded_y
}

func main() {
	data := read_input()
	values := make(map[[4]int]int)
	galaxies := get_galaxies(data)
	for _, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies {
			if galaxy1 == galaxy2 { continue }
			couple_identifier := get_couple_identifier(galaxy1,galaxy2)
			if _, ok := values[couple_identifier]; !ok {
				values[couple_identifier] = get_galaxies_distance(galaxy1,galaxy2,galaxies)
			}
		}
	}
	sum := 0
	for _, value := range values {
		sum += value
	}	
	fmt.Println(sum)

}