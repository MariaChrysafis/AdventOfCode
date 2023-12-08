package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

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
	right := make(map[string]string)
	left := make(map[string]string)
	directions := arr[0]
	for i, x := range arr {
		if i > 1 {
			right[x[:3]] = x[12:15]
			left[x[:3]] = x[7:10]
		}
	}
	endUp := make(map[string]string)
	reachable := make(map[string]int)
	start := "AAA"
	end := "ZZZ"
	for key, _ := range left {
		cur := key
		reachable[key] = -1
		for steps, c := range directions {
			if string(c) == "L" {
				cur = left[cur]
			} else {
				cur = right[cur]
			}
			if cur == end {
				reachable[key] = steps + 1
				break
			}
		}
		endUp[key] = cur
	}
	reachable[end] = 0
	ans := 0
	for true {
		if reachable[start] != -1 {
			fmt.Println(ans + reachable[start])
			break
		}
		ans += len(directions)
		start = endUp[start]
	}
}
