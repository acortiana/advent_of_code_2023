package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func text_spaced_numbers_to_uint64_slice(text string) []uint64 {
	var result []uint64
	for _, textnumber := range strings.Fields(text) {
		i, err := strconv.ParseUint(textnumber, 10, 64)
		if err != nil { return nil }
		result = append(result,i)
	}
	return result
}

func find_destination_number(source uint64, table *[][]uint64) uint64 {
	for i := 0; i < len(*table); i++ {
		if source >= (*table)[i][1] && source <= (*table)[i][1] + ((*table)[i][2] - 1) {
			return (*table)[i][0] + (source - (*table)[i][1])
		}
	}
	return source
}

func read_input() ([][][]uint64, []uint64) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	seeds := text_spaced_numbers_to_uint64_slice(scanner.Text()[7:])
	counter := 0
	lastdigit := false
	var data [][][]uint64
	for scanner.Scan() {
		line := scanner.Text()
		myslice := text_spaced_numbers_to_uint64_slice(line)
		if myslice == nil {
			if lastdigit == true {
				counter++
			}
			lastdigit = false
			continue
		}
		lastdigit = true
		if len(data) <= counter {
			data = append(data,[][]uint64{})
		}
		data[counter] = append(data[counter],myslice)
	}
	return data, seeds
}

func worker(threadcount int, threadnumber int, data [][][]uint64, messages chan uint64, seeds []uint64) {
	var lowest_location uint64 = ^uint64(0)
	for i:= 0; (i*2+1) < len(seeds); i++ {
		start := seeds[i*2]
		lenght := seeds[i*2+1]
		end := start + lenght - 1
		for start = start + uint64(threadnumber); start <= end; start += uint64(threadcount) {
			seed := start
			for i:=0; i < len(data); i++ {
				seed = find_destination_number(seed,&data[i])
			}
			if seed < lowest_location {
				lowest_location = seed
			}
		}
	}
	messages <- lowest_location
}

func main() {
	data, seeds := read_input()
	messages := make(chan uint64)
	threadcount := 8
	for i:= 0; i < threadcount; i++ {
		go worker(threadcount,i,data,messages,seeds)
	}
	var lowest_location uint64 = ^uint64(0)
	for i:= 0; i < threadcount; i++ {
		msg := <-messages
		if msg < lowest_location {
			lowest_location = msg
		}
	}
	fmt.Println(lowest_location)
}