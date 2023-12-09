package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func stringToInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
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
	for _, s := range arr {
		numeric := ""
		for _, c := range s {
			if '0' <= c && c <= '9' {
				numeric += string(c)
			}
		}
		if len(numeric) == 1 {
			numeric += numeric
		}
		ans += stringToInt(string(numeric[0])) * 10
		ans += stringToInt(string(numeric[len(numeric)-1]))
	}
	fmt.Print(ans)
}
