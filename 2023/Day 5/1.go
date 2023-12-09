package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func convert(s string) []int { //converts array of strings to array of integers
	arr := []int{0}
	for _, c := range s {
		if c == ' ' {
			arr = append(arr, 0)
		}
		arr[len(arr)-1] *= 10
		val, _ := strconv.Atoi(string(c))
		arr[len(arr)-1] += val
	}
	return arr
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)
	arr := make([]string, 1)
	for i := range input {
		if string(input[i]) == "\n" {
			arr = append(arr, "")
		} else {
			arr[len(arr)-1] += string(input[i])
		}
	}
	inp := [][][]int{}
	var seeds []int
	for i, x := range arr {
		if i == 0 {
			seeds = convert(x)
			continue
		}
		if x == "" {
			inp = append(inp, make([][]int, 0))
		} else {
			inp[len(inp)-1] = append(inp[len(inp)-1], convert(x))
		}
	}
	lengths := []int{}
	for _, seed := range seeds {
		val := seed
		for _, a := range inp {
			for _, b := range a {
				if val >= b[1] && val <= b[1]+b[2]-1 {
					val += b[0] - b[1]
					break
				}
			}
		}
		lengths = append(lengths, val)
	}
	mn := lengths[0]
	for _, x := range lengths {
		mn = min(mn, x)
	}
	fmt.Print(mn)
}
