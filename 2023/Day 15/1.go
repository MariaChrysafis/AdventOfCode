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
	get_hash := func(s string) int {
		ans := 0
		for _, c := range s {
			ans = ((ans + int(c)) * 17) % 256
		}
		return ans
	}
	ans := 0
	for _, x := range strings.Split(string(content), ",") {
		ans += get_hash(x)
	}
	fmt.Println(ans)
}
