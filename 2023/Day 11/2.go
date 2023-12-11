package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func swap(arr []string) []string {
	ans := make([]string, len(arr[0]))
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			ans[j] += string(arr[i][j])
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
	part1 := func(repeat int) int {
		t := 2
		for t != 0 {
			t--
			inp = swap(inp)
			l := len(inp)
			for i := l - 1; i >= 0; i-- {
				cur := inp[i]
				if strings.Index(cur, "#") <= -1 {
					for x := 0; x < repeat-1; x++ {
						if i != l-1 {
							inp = append(inp[:i+1], inp[i:]...)
							inp[i] = cur
						} else {
							inp = append(inp, inp[i])
						}
					}
				}
			}
		}
		galaxies := [][]int{}
		for i := range inp {
			for j := range inp[i] {
				if inp[i][j] == '#' {
					galaxies = append(galaxies, []int{i, j})
				}
			}
		}
		sum := 0
		for _, g1 := range galaxies {
			for _, g2 := range galaxies {
				sum += abs(g1[0]-g2[0]) + abs(g1[1]-g2[1])
			}
		}
		return sum / 2
	}
	orig := inp
	x1 := part1(2)
	inp = orig
	x2 := part1(3)
	mult := 1000000
	fmt.Print((mult-2)*(x2-x1) + x1)
}
