package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"slices"
)

func atoi(textnumber string) int {
	tmpvalue, _:= strconv.Atoi(textnumber)
	return tmpvalue
}

func read_input() ([][5]byte, map[string]int) {
	var hands [][5]byte
	mymap := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tmp_1 := strings.Fields(line)
		var tmp_2 [5]byte
		copy(tmp_2[:],tmp_1[0])
		hands = append(hands,tmp_2)
		mymap[tmp_1[0]] = atoi(tmp_1[1])
	}
	return hands, mymap
}

func get_card_number(x byte) byte {
	switch x {
		case 'A':
			return 12
		case 'K':
			return 11
		case 'Q':
			return 10
		case 'J':
			return 0
		case 'T':
			return 9
		default:
			return x - '1'
	}
}

func compare_cards(x byte, y byte) int {
	if x == y {
		return 0
	}
	tmp := [2]byte{x,y}
	for i := 0; i < 2; i++ {
		tmp[i] = get_card_number(tmp[i])
	}
	if tmp[0] > tmp[1] { return 1 }
	return -1
}

func get_hand_type(hand [5]byte) byte {
	var cardcount [13]byte
	for i := 0; i < len(hand); i++ {
		cardcount[get_card_number(hand[i])]++
	}

	for i := 12; i > 0; i-- {
		if cardcount[i] + cardcount[0] == 5 { return 6 }
	}

	for i := 12; i > 0; i-- {
		if cardcount[i] + cardcount[0] == 4 { return 5 }
	}

	for i := 12; i > 0; i-- {
		if cardcount[i] + cardcount[0] == 3 {
			for j := 12; j > 0; j-- {
				if cardcount[j] == 2 && j != i { return 4 }
			}
			return 3
		}
	}

	for i := 12; i > 0; i-- {
		if cardcount[i] + cardcount[0] == 2 {
			for j := 12; j > 0; j-- {
				if cardcount[j] == 2 && j != i { return 2 }
			}
			return 1
		}
	}
	return 0
}

func are_hands_equal(x [5]byte, y [5]byte) bool {
	for i:= 0; i < len(x); i++ {
		if x[i] != y[i] { return false }
	}
	return true
}

func compare_hands(x [5]byte, y [5]byte) int {
	if are_hands_equal(x,y) {
		return 0
	}
	hands := [2]*[5]byte{&x, &y}
	var handstype [2]byte
	for i, hand := range hands {
		handstype[i] = get_hand_type(*hand)
	}
	if handstype[0] > handstype[1] { return 1 }
	if handstype[0] < handstype[1] { return -1 }
	for i := 0; i < len(x); i++ {
		card_compare_result := compare_cards(x[i],y[i])
		if card_compare_result != 0 { return card_compare_result }
	}
	return 0
}

func main() {
	hands, mymap := read_input()
	slices.SortFunc(hands,compare_hands)
	totalwinning := 0
	for i, hand := range hands {
		totalwinning += mymap[string(hand[:])] * (i+1)
	}
	fmt.Println(totalwinning)
}