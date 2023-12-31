package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
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
func get_mn(arr []int, a [][][]int) int {
	if len(a) == 0 {
		return arr[0]
	}
	special := []int{arr[0], arr[1]}
	for _, b := range a[0] {
		y := b[1]
		length := b[2]
		if max(y, arr[0]) <= arr[1] {
			special = append(special, max(y, arr[0]))
		}
		if min(y+length-1, arr[1]) >= arr[0] {
			special = append(special, min(y+length-1, arr[1]))
		}
	}
	sort.Slice(special, func(i int, j int) bool {
		return special[i] < special[j]
	})
	intervals := [][]int{{arr[0], arr[0]}}
	for i := 1; i < len(special); i++ {
		if special[i-1]+1 <= special[i] {
			intervals = append(intervals, []int{special[i-1] + 1, special[i]})
		}
	}
	mn := int(1e9)
	for _, x := range intervals {
		offset := 0
		r := x[0]
		for _, b := range a[0] {
			if r >= b[1] && r <= b[1]+b[2]-1 {
				offset = b[0] - b[1]
				break
			}
		}
		x[0] += offset
		x[1] += offset
		if x[0] > x[1] {
			x[1], x[0] = x[0], x[1]
		}
		mn = min(mn, get_mn(x, a[1:]))
	}
	return mn
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
		} else {
			if x == "" {
				inp = append(inp, make([][]int, 0))
			} else {
				inp[len(inp)-1] = append(inp[len(inp)-1], convert(x))
			}
		}
	}
	ans := int(1e9)
	for i := 0; i < len(seeds); i += 2 {
		ans = min(ans, get_mn([]int{seeds[i], seeds[i] + seeds[i+1] - 1}, inp))
	}
	fmt.Print(ans)
}
