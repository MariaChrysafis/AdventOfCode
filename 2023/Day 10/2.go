package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"slices"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
	s := ""
	for j := 0; j < len(inp[0]); j++ {
		s += "."
	}
	inp = append(inp, s)
	inp = append([]string{s}, inp...)
	for i := range inp {
		inp[i] = "." + inp[i] + "."
	}
	n, m := len(inp), len(inp[0])
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = -1
		}
	}
	sx, sy := 0, 0
	for i := range inp {
		for j := range inp[0] {
			if inp[i][j] == 'S' {
				sx, sy = i, j
			}
		}
	}
	q := [][]int{{sx, sy}}
	dist[sx][sy] = 0
	path := [][]int{}
	for len(q) != 0 {
		x, y := q[0][0], q[0][1]
		path = append(path, []int{x, y})
		q = q[1:]
		if inp[x][y] == '.' {
			continue
		}
		mx, my := make(map[string][]int), make(map[string][]int)
		for _, s := range []string{"L", "J", "|", "S"} {
			mx[s] = append(mx[s], -1)
		}
		for _, s := range []string{"F", "7", "|", "S"} {
			mx[s] = append(mx[s], 1)
		}
		for _, s := range []string{"L", "F", "-"} {
			my[s] = append(my[s], 1)
		}
		for _, s := range []string{"J", "7", "-"} {
			my[s] = append(my[s], -1)
		}
		for _, dx := range mx[string(inp[x][y])] { //going north-south
			if slices.Contains(mx[string(inp[x+dx][y])], -dx) && dist[x+dx][y] == -1 && inp[x+dx][y] != '.' {
				q = append(q, []int{x + dx, y})
				dist[x+dx][y] = 1
			}
		}
		for _, dy := range my[string(inp[x][y])] {
			if slices.Contains(my[string(inp[x][y+dy])], -dy) && dist[x][y+dy] == -1 && inp[x][y+dy] != '.' {
				q = append(q, []int{x, y + dy})
				dist[x][y+dy] = 1
			}
		}
	}
	area := 0
	for i := 0; i < len(path); i++ {
		area += (path[i][1] + path[(i+1)%len(path)][1]) * (path[i][0] - path[(i+1)%len(path)][0])
	}
	fmt.Println(abs(area)/2 - len(path)/2 + 1)
}
