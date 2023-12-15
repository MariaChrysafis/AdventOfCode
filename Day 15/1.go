package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), ",")
	ans := 0
	for _, s := range input {
		res := 0
		for _, c := range s {
			res += int(c)
			res *= 17
			res %= 256
		}
		ans += res
	}
	fmt.Println(ans)
}
