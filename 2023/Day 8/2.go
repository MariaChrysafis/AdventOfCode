package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func gcd(a int, b int) int {
	if a == 0 || b == 0 {
		return max(a, b)
	}
	return gcd(max(a, b)%min(a, b), min(a, b))
}
func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
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
	get_distance := func(start string, end string) int { //shortest distance between two nodes
		for key, _ := range left {
			cur := key
			reachable[key] = -1
			for steps, c := range directions {
				if c == 'L' {
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
		visited := make(map[string]bool)
		for reachable[start] == -1 {
			if visited[start] == true {
				return -1
			}
			visited[start] = true
			ans += len(directions)
			start = endUp[start]
		}
		return ans + reachable[start]
	}
	starting := make([]string, 0)
	ending := make([]string, 0)
	for key, _ := range left {
		if key[2] == 'A' {
			starting = append(starting, key)
		}
		if key[2] == 'Z' {
			ending = append(ending, key)
		}
	}
	ans := 0
	for _, x := range starting {
		for _, y := range ending {
			d := get_distance(x, y)
			if get_distance(x, y) != -1 {
				if ans == 0 {
					ans = d
				} else {
					ans = lcm(ans, d)
				}
			}
		}
	}
	fmt.Println(ans)
}
