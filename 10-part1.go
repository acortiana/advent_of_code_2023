package main

import (
	"fmt"
	"bufio"
	"os"
)

func read_input() ([]string) {
	var data []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data,line)
	}
	return data
}

func find_s(input []string) [2]int {
	for i:=0; i < len(input); i++ {
		for j:=0; j < len(input[i]); j++ {
			if input[i][j] == 'S' {
				return [2]int{j,i}
			}
		}
	}
	return [2]int{-1,-1}
}

func get_direction(src [2]int, dst[2]int) byte {
	if src[0] > dst[0] {
		return 'W'
	} else if src[0] < dst[0] {
		return 'E'
	}
	if src[1] > dst[1] {
		return 'N'
	} else if src[1] < dst[1]{
		return 'S'
	}
	return 0
}

func get_char(point [2]int, data[]string) byte {
	return data[point[1]][point[0]]
}

func is_possible(src [2]int, dst [2]int,data []string) bool {
	src_char := get_char(src,data)
	dst_char := get_char(dst,data)
	direction := get_direction(src,dst)
	if direction == 'N' {
		if src_char == '-' || src_char == '.' || src_char == '7' || src_char == 'F' { return false }
		if dst_char == '-' || dst_char == '.' || dst_char == 'J' || dst_char == 'L' { return false }
	}
	if direction == 'S' {
		if src_char == '-' || src_char == '.' || src_char == 'J' || src_char == 'L' { return false }
		if dst_char == '-' || dst_char == '.' || dst_char == '7' || dst_char == 'F' { return false }
	}
	if direction == 'E' {
		if src_char == '|' || src_char == '.' || src_char == 'J' || src_char == '7' { return false }
		if dst_char == '|' || dst_char == '.' || dst_char == 'L' || dst_char == 'F' { return false }
	}
	if direction == 'W' {
		if src_char == '|' || src_char == '.' || src_char == 'L' || src_char == 'F' { return false }
		if dst_char == '|' || dst_char == '.' || dst_char == 'J' || dst_char == '7' { return false }
	}
	return true
}

func find_point_destinations(point [2]int, data[]string) [][2]int {
	var results [][2]int
	for i :=-1; i < 2; i++ {
		for j :=-1; j < 2; j++ {
			if i == 0 && j == 0 { continue }
			if i != 0 && j != 0 { continue }
			newpoint := [2]int{point[0]+i,point[1]+j}
			if newpoint[0] < 0 || newpoint[1] < 0 { continue }
			if newpoint[0] >= len(data[0]) || newpoint[1] >= len(data) { continue }
			results = append(results,newpoint)
		}
	}
	return results
}

func find_next_dst(src [2]int, dst[2]int, data []string) [][2]int {
	var results [][2]int
	for _, newpoint := range find_point_destinations(dst,data) {
		if newpoint == src { continue }
		if is_possible(dst,newpoint,data) { 
			results = append(results,newpoint)
		}
	}
	return results
}

func main() {
	data := read_input()
	start := find_s(data)
	src := find_s(data)
	dst := find_s(data)
	steps := 0
	for ; ; steps++ {
		newdst := find_next_dst(src,dst,data)[0]
		src = dst
		dst = newdst
		if dst == start { break }
	}
	fmt.Println((steps+1)/2)
}