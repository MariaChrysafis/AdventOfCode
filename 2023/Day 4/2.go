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
	arr := make([]string, 1)
	for i := range input {
		if string(input[i]) == "\n" {
			arr = append(arr, "")
		} else {
			arr[len(arr)-1] += string(input[i])
		}
	}
	objectives := [][]string{}
	cards := [][]string{}
	copies := make([]int, len(arr))
	for i := range copies {
		copies[i] = 1
	}
	for _, x := range arr {
		new_arr := cleanUp(strings.Split(x, " "))
		new_arr = new_arr[1:]
		new_arr[0] = new_arr[0][1:]
		for i := range new_arr {
			new_arr[i] = strings.ReplaceAll(new_arr[i], " ", "")
		}
		new_arr = cleanUp(new_arr)
		boundary := 0
		for i := 0; i < len(new_arr); i++ {
			if new_arr[i] == "|" {
				boundary = i
			}
		}
		objectives = append(objectives, new_arr[:boundary])
		cards = append(cards, new_arr[boundary+1:])
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
