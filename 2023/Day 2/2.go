package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func split(s string, c string) []string {
	arr := []string{""}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == c {
			arr = append(arr, "")
		} else {
			arr[len(arr)-1] += string(s[i])
		}
	}
	return arr
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
	for _, x := range arr {
		okay := true
		mx := make(map[string]int)
		res := 1
		for _, game := range split(split(x, ":")[1:][0][1:], ";") {
			m := make(map[string]int)
			for _, s := range split(game, ",") {
				if string(s[0]) == " " {
					s = s[1:]
				}
				x := split(s, " ")
				m[x[1]] += stringToInt(x[0])
			}
			for key, value := range m {
				mx[key] = max(mx[key], value)
			}
		}
		for _, value := range mx {
			res *= value
		}
		if okay {
			ans += res
		}
	}
	fmt.Println(ans)
}
