package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"slices"
	"strconv"
	"strings"
)

type Pair struct {
	x int
	y int
}

func add(p1 Pair, p2 Pair) Pair {
	return Pair{p1.x + p2.x, p1.y + p2.y}
}

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
	inp := make([]string, 1)
	for i := range input {
		if string(input[i]) == "\n" {
			inp = append(inp, "")
		} else {
			inp[len(inp)-1] += string(input[i])
		}
	}
	points := make([]Pair, 0)
	edges := make([]Pair, 0)
	do_move := func(x *int, y *int, a string, d int) {
		for i := 0; i < d; i++ {
			edges = append(edges, Pair{*x, *y})
			if a == "R" {
				*x += 1
			} else if a == "L" {
				*x -= 1
			} else if a == "U" {
				*y += 1
			} else if a == "D" {
				*y -= 1
			}
		}
		points = append(points, Pair{*x, *y})
	}
	pos_x, pos_y := 0, 0
	res := 0
	for _, i := range inp {
		arr := strings.Split(i, " ")
		do_move(&pos_x, &pos_y, arr[0], stringToInt(arr[1]))
		res += stringToInt(arr[1])
	}
	ans := 0
	for x := -500; x <= 500; x++ {
		for y := -500; y <= 500; y++ {
			//look at vertical line
			if slices.Contains(edges, Pair{x, y}) {
				ans += 1
				continue
			}
			tot := 0
			for i := range points {
				//edge connects points[i] to points[i + 1]
				if points[i].x == points[(i+1)%len(points)].x && x >= points[i].x {
					//horizontal line
					l := points[i].y
					r := points[(i+1)%len(points)].y
					l, r = min(l, r), max(l, r)
					if l <= y && y < r {
						tot += 1
					}
				}
			}
			if tot%2 == 1 {
				ans += 1
			}
		}
	}
	fmt.Println(ans)
}
