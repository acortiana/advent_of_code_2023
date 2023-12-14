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

func get_s_char(input []string) byte {
	coords := find_s(input)
	N, S, E, W := false, false, false ,false
	points := find_possible_destinations(coords,input)
	for _, point := range points {
		direction := get_direction(coords,point)
		switch direction {
			case 'N':
				N = true
			case 'S':
				S = true
			case 'E':
				E = true
			case 'W':
				W = true
		}
	}
	if N && S { return '|'}
	if N && E { return 'L'}
	if N && W { return 'J'}
	if S && E { return 'F'}
	if S && W { return '7'}
	if E && W { return '-'}
	return 0
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

func find_possible_destinations(point[2]int, data []string) [][2]int {
	var results [][2]int
	for _, newpoint := range find_point_destinations(point,data) {
		if is_possible(point,newpoint,data) { 
			results = append(results,newpoint)
		}
	}
	return results
}

func find_next_dst(src [2]int, dst[2]int, data []string) [][2]int {
	result1 := find_possible_destinations(dst,data)
	var result2 [][2]int
	for _, point := range result1 {
		if point == src { continue }
		result2 = append(result2,point)
	}
	return result2
}

func is_point_in_point_slice(point [2]int, points_slice [][2]int) bool {
	for _, tmp := range points_slice {
		if tmp == point { return true }
	}
	return false
}

func is_in_loop(point [2]int, points_slice [][2]int, data []string) bool {
	loop := false
	y := point[1]
	var prevchar byte
	for x:=0; x < len(data[0]); x++ {
		coords := [2]int{x,y}
		if is_point_in_point_slice(coords,points_slice) {
			mychar := get_char(coords,data)
			if mychar == 'S' { mychar = get_s_char(data)}
			if (mychar == '7' && prevchar == 'L') || 
			   (mychar == 'J' && prevchar == 'F') {

			} else if mychar != '-' {
				loop = !loop
			} else {
				continue
			}
			prevchar = mychar
			continue
		}
		if coords == point { return loop }
		prevchar = '0'
	}
	return false
}

func main() {
	data := read_input()
	start := find_s(data)
	var points_slice [][2]int
	src := find_s(data)
	dst := find_s(data)
	steps := 0
	for ; ; steps++ {
		points_slice = append(points_slice,dst)
		newdst := find_next_dst(src,dst,data)[0]
		src = dst
		dst = newdst
		if dst == start { break }
	}
	counter := 0
	for x:=0; x < len(data[0]); x++ {
		for y:=0; y < len(data); y++ {
			coords := [2]int{x,y}
			if is_in_loop(coords,points_slice,data) {
				counter++
			}
		}
	}
	fmt.Println(counter)
}