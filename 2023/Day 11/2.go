package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"slices"
	"strings"
)

func abs(x int) int {
	return max(x, -x)
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
	part1 := func(repeat int, arr []string) int {
		for t := 0; t < 2; t++ {
			arr = swap(arr)
			for i := len(arr) - 1; i >= 0; i-- {
				cur := arr[i]
				if strings.Index(cur, "#") == -1 {
					for x := 0; x < repeat-1; x++ {
						arr = slices.Insert(arr, i, cur)
					}
				}
			}
		}
		galaxies := [][]int{}
		for i := range arr {
			for j := range arr[i] {
				if arr[i][j] == '#' {
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
	mult := 1000000
	fmt.Print((mult-1)*part1(2, inp) - (mult-2)*part1(1, inp))
}
