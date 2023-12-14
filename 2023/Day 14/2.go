package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func reverse(s string) string {
	ans := ""
	for _, v := range s {
		ans = string(v) + ans
	}
	return ans
}
func reflect1(arr []string) []string {
	for i := 0; i*2 < len(arr); i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr
}
func reflect2(arr []string) []string {
	for i := range arr {
		arr[i] = reverse(arr[i])
	}
	return arr
}
func transpose(arr []string) []string {
	ans := make([]string, len(arr[0]))
	for j := 0; j < len(arr); j++ {
		for i := 0; i < len(arr[j]); i++ {
			ans[i] += string(arr[j][i])
		}
	}
	return ans
}
func compress(arr []string) string {
	s := ""
	for _, x := range arr {
		s += x
	}
	return s
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
	go_north := func() int {
		ans := 0
		for i := range inp {
			for j := range inp[i] {
				x := i
				for x != 0 && inp[x-1][j] == '.' && inp[x][j] == 'O' {
					inp[x] = inp[x][:j] + "." + inp[x][j+1:]
					inp[x-1] = inp[x-1][:j] + "O" + inp[x-1][j+1:]
					x -= 1
				}
				if inp[x][j] == 'O' {
					ans += len(inp) - x
				}
			}
		}
		return ans
	}
	print := func() {
		for _, x := range inp {
			fmt.Println(x)
		}
		fmt.Println()
	}
	cost := func(arr []string) int {
		ans := 0
		for i := range arr {
			for j := range arr[i] {
				if arr[i][j] == 'O' {
					ans += len(arr) - i
				}
			}
		}
		return ans
	}
	cycle := func() {
		//north
		go_north()
		//west
		inp = transpose(inp)
		go_north()
		inp = transpose(inp)
		//south
		inp = reflect1(inp)
		go_north()
		inp = reflect1(inp)
		//east
		inp = reflect2(inp)
		inp = transpose(inp)
		go_north()
		inp = transpose(inp)
		inp = reflect2(inp)
	}
	lastseen := make(map[string][]int)
	for t := 1; t <= int(1e9); t++ {
		fmt.Println(t)
		cycle()
		fmt.Println(cost(inp))
		a := lastseen[compress(inp)]
		if len(a) != 0 {
			fmt.Println(t, a[0])
			d := t - a[0]
			for t+d < int(1e9) {
				t += d
			}
		} else {
			lastseen[compress(inp)] = []int{t}
		}
		print()
	}
}
