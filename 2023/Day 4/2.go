package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func cleanUp(arr []string) []string {
	new_arr := []string{}
	for _, x := range arr {
		x = strings.ReplaceAll(x, " ", "")
		if x == "" {
			continue
		}
		new_arr = append(new_arr, x)
	}
	return new_arr
}
func intersection(a1 []string, a2 []string) int {
	ans := 0
	for _, x := range a1 {
		for _, y := range a2 {
			if x == y {
				ans += 1
			}
		}
	}
	return ans
}
func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)
	inp := make([]string, 1)
	for i := range input {
		if string(input[i]) == "\n" {
			inp = append(inp, "")
		} else {
			inp[len(inp)-1] += string(input[i])
		}
	}
	objectives := [][]string{}
	cards := [][]string{}
	copies := make([]int, len(inp))
	for i, x := range inp {
		copies[i] = 1
		arr := cleanUp(strings.Split(x, " "))[1:]
		arr[0] = arr[0][1:]
		for i := range arr {
			arr[i] = strings.ReplaceAll(arr[i], " ", "")
		}
		boundary := 0
		for i := 0; i < len(arr); i++ {
			if arr[i] == "|" {
				boundary = i
			}
		}
		objectives = append(objectives, arr[:boundary])
		cards = append(cards, arr[boundary+1:])
	}
	sum := 0
	for i := range objectives {
		l := intersection(objectives[i], cards[i])
		for x := i + 1; x < i+l+1; x++ {
			copies[x] += copies[i]
		}
		sum += copies[i]
	}
	fmt.Println(sum)
}
