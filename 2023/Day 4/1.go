package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func cleanUp(arr []string) []string {
	new_arr := []string{}
	for _, x := range arr {
		x = strings.ReplaceAll(x, " ", "")
		if x == "" {
			continue
		}
		new_arr = append(new_arr, x)
	}
	return new_arr
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
	ans := 0
	for _, x := range inp {
		arr := cleanUp(strings.Split(x, " "))[1:]
		arr[0] = arr[0][1:]
		res := 0
		for i := range arr {
			arr[i] = strings.ReplaceAll(arr[i], " ", "")
		}
		boundary := 0
		for i := 0; i < len(arr); i++ {
			if arr[i] == "|" {
				boundary = i
			}
		}
		for i := 0; i < boundary; i++ {
			found := false
			for j := boundary + 1; j < len(arr); j++ {
				found = found || (arr[j] == arr[i])
			}
			if found {
				res += 1
			}
		}
		if res != 0 {
			ans += (1 << (res - 1))
		}
	}
	fmt.Println(ans)
}
