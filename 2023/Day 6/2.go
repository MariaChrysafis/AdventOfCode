package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

func convert(arr []string) []int { //converts array of strings to array of integers
	s := ""
	for _, x := range arr {
		s += x
	}
	x, _ := strconv.Atoi(s)
	return []int{x}
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
	r := regexp.MustCompile("[^\\s]+")
	time := convert(r.FindAllString(arr[0], -1))
	distance := convert(r.FindAllString(arr[1], -1))
	ans := 1
	for i := 0; i < len(distance); i++ {
		cnt := 0
		for x := 0; x <= time[i]; x++ {
			res := x * (time[i] - x)
			if res > distance[i] {
				cnt += 1
			}
		}
		ans *= cnt
	}
	fmt.Println(ans)
}
