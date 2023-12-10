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
	for x := 0; x < len(arr); x++ {
		for j := 0; j < len(arr[x]); j++ {
			//check if this is a number
			if "0" <= string(arr[x][j]) && string(arr[x][j]) <= "9" {
				l := j
				r := j
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
							if dx+x < 0 || dx+x == len(arr) {
								continue
							}
							if dy+y < 0 || dy+y == len(arr[0]) {
								continue
							}
							if arr[x+dx][y+dy] != '.' && !isNumeric(string(arr[x+dx][y+dy])) {
								okay = false
							}
						}
					}
				}
				if !okay {
					str := ""
					for y := l; y <= r; y++ {
						str += string(arr[x][y])
					}
					ans += stringToInt(str)
				}
			}
		}
	}
	fmt.Println(ans)
}
