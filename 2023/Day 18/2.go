package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func stringToInt(s string) int64 {
	val, _ := strconv.ParseInt(s[2:len(s)-2], 16, 64)
	return val
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
	move := func(x *int64, y *int64, a uint8, d int64) {
		if a == uint8('0') {
			*x += d
		} else if a == uint8('2') {
			*x -= d
		} else if a == uint8('3') {
			*y += d
		} else if a == uint8('1') {
			*y -= d
		}
	}
	x, y := int64(0), int64(0)
	points := [][]int64{{0, 0}}
	boundary := int64(0)
	for _, i := range inp {
		arr := strings.Split(i, " ")
		move(&x, &y, arr[2][len(arr[2])-2], stringToInt(arr[2]))
		points = append(points, []int64{x, y})
		boundary += stringToInt(arr[2])
	}
	area := int64(0)
	for i := 0; i < len(points); i++ {
		area += points[i][0] * points[(i+1)%len(points)][1]
		area -= points[(i+1)%len(points)][0] * points[i][1]
	}
	area = max(area, -area)
	fmt.Println(area/2 + boundary/2 + 1)
}
