package main

import (
	"fmt"
	"io/ioutil"
)

type Position struct {
	x int
	y int
}

func contains(pos []Position, p Position) bool {
	for _, y := range pos {
		if y.x == p.x && y.y == p.y {
			return true
		}
	}
	return false
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	input := string(content)
	inp := make([]string, 1)
	for i := range input {
		if string(input[i]) == "\n" {
			inp = append(inp, "")
		} else {
			inp[len(inp)-1] += string(input[i])
		}
	}
	possible := make([]Position, 0)
	for i := range inp {
		for j := range inp[i] {
			if inp[i][j] == 'S' {
				possible = append(possible, Position{i, j})
			}
		}
	}
	for t := 0; t < 64; t++ {
		new_possible := make([]Position, 0)
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if abs(dx)+abs(dy) != 1 {
					continue
				}
				for _, p := range possible {
					if p.x+dx >= 0 && p.x+dx < len(inp) && p.y+dy >= 0 && p.y+dy < len(inp[0]) {
						if inp[p.x+dx][p.y+dy] == '#' {
							continue
						}
						if contains(new_possible, Position{p.x + dx, p.y + dy}) {
							continue
						}
						new_possible = append(new_possible, Position{p.x + dx, p.y + dy})
					}
				}
			}
		}
		fmt.Println(len(new_possible))
		possible = new_possible
	}
}
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
func abs(x int) int {
	return max(x, -x)
}
