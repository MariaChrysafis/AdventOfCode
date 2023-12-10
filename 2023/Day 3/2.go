package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func isNumeric(s string) bool {
	return "0" <= s && s <= "9"
}
func stringToInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
	ans := 0
	parts := [][]int{}
	for x := 0; x < len(arr); x++ {
		for j := 0; j < len(arr[x]); j++ {
			//check if this is a number
			if "0" <= string(arr[x][j]) && string(arr[x][j]) <= "9" {
				l, r := j, j
				for l != 0 && isNumeric(string(arr[x][l-1])) {
					l -= 1
				}
				for r != len(arr[0])-1 && isNumeric(string(arr[x][r+1])) {
					r += 1
				}
				if l != j {
					continue
				}
				okay := true
				for y := l; y <= r; y++ {
					for dx := -1; dx <= 1; dx++ {
						for dy := -1; dy <= 1; dy++ {
							if dx+x >= 0 && dx+x < len(arr) && dy+y >= 0 && dy+y < len(arr[0]) {
								if arr[x+dx][y+dy] != '.' && !isNumeric(string(arr[x+dx][y+dy])) {
									okay = false
								}
							}
						}
					}
				}
				if !okay {
					parts = append(parts, []int{x, l, r})
				}
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == '*' {
				cnt := 0
				prod := 1
				res := []int{}
				for _, p := range parts {
					if abs(p[0]-i) <= 1 {
						if (p[1] <= j && j <= p[2]) || abs(p[1]-j) <= 1 || abs(p[2]-j) <= 1 {
							cnt += 1
							prod *= stringToInt(arr[p[0]][p[1] : p[2]+1])
							res = append(res, stringToInt(arr[p[0]][p[1]:p[2]+1]))
						}
					}
				}
				if cnt == 2 {
					ans += prod
				}
			}
		}
	}
	fmt.Println(ans)
}
