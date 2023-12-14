package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

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
	fmt.Println(ans)
}
