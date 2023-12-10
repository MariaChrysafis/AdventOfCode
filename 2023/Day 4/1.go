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
		new_arr := cleanUp(strings.Split(x, " "))
		new_arr = new_arr[1:]
		new_arr[0] = new_arr[0][1:]
		res := 0
		for i := range new_arr {
			new_arr[i] = strings.ReplaceAll(new_arr[i], " ", "")
		}
		new_arr = cleanUp(new_arr)
		boundary := 0
		for i := 0; i < len(new_arr); i++ {
			if new_arr[i] == "|" {
				boundary = i
			}
		}
		for i := 0; i < boundary; i++ {
			found := false
			for j := boundary + 1; j < len(new_arr); j++ {
				found = found || (new_arr[j] == new_arr[i])
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
