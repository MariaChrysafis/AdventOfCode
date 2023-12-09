package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

func convert(arr []string) []int { //converts array of strings to array of integers
	ans := make([]int, 0)
	for _, x := range arr {
		val, _ := strconv.Atoi(x)
		ans = append(ans, val)
	}
	return ans
}

func allZeroes(arr []int) bool {
	for _, x := range arr {
		if x != 0 {
			return false
		}
	}
	return true
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
	inp := make([][]int, 0)
	for _, x := range arr {
		r := regexp.MustCompile("[^\\s]+")
		inp = append(inp, convert(r.FindAllString(x, -1)))
	}
	ans := 0
	for _, arr := range inp {
		dist := make([][]int, 0)
		dist = append(dist, arr)
		for !allZeroes(dist[len(dist)-1]) {
			nxt := make([]int, 0)
			for i := 1; i < len(dist[len(dist)-1]); i++ {
				nxt = append(nxt, dist[len(dist)-1][i]-dist[len(dist)-1][i-1])
			}
			dist = append(dist, nxt)
		}
		for i, x := range dist {
			if i%2 == 0 {
				ans += x[0]
			} else {
				ans -= x[0]
			}
		}
	}
	fmt.Println(ans)
}
