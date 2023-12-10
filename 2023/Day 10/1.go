package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"slices"
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
	s := ""
	for j := 0; j < len(inp[0]); j++ {
		s += "."
	}
	inp = append(inp, s)
	inp = append([]string{s}, inp...)
	for i := range inp {
		inp[i] = "." + inp[i] + "."
	}
	n := len(inp)
	m := len(inp[0])
	dist := make([][]int, n)
	ans := 0
	for i := range inp {
		for j := range inp[0] {
			if inp[i][j] == 'S' {
				for x := 0; x < n; x++ {
					dist[x] = make([]int, m)
					for y := 0; y < m; y++ {
						dist[x][y] = -1
					}
				}
				q := [][]int{{i, j}}
				dist[i][j] = 0
				for len(q) != 0 {
					x, y := q[0][0], q[0][1]
					ans = max(ans, dist[x][y])
					q = q[1:]
					if inp[x][y] == '.' {
						dist[x][y] = -1
						continue
					}
					mx, my := make(map[string][]int), make(map[string][]int)
					//north (x - 1)
					for _, s := range []string{"L", "J", "|", "S"} {
						mx[s] = append(mx[s], -1)
					}
					//south (x + 1)
					for _, s := range []string{"F", "7", "|", "S"} {
						mx[s] = append(mx[s], 1)
					}
					//east (y + 1)
					for _, s := range []string{"L", "F", "-", "S"} {
						my[s] = append(my[s], 1)
					}
					//west
					for _, s := range []string{"J", "7", "-", "S"} {
						my[s] = append(my[s], -1)
					}
					for _, dx := range mx[string(inp[x][y])] { //going north-south
						if slices.Contains(mx[string(inp[x+dx][y])], -dx) && dist[x+dx][y] == -1 && inp[x+dx][y] != '.' {
							q = append(q, []int{x + dx, y})
							dist[x+dx][y] = dist[x][y] + 1
						}
					}
					for _, dy := range my[string(inp[x][y])] {
						if slices.Contains(my[string(inp[x][y+dy])], -dy) && dist[x][y+dy] == -1 && inp[x][y+dy] != '.' {
							q = append(q, []int{x, y + dy})
							dist[x][y+dy] = dist[x][y] + 1
						}
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
