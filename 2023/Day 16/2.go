package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Position struct {
	x  int
	y  int
	dx int
	dy int
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
	get := func(start Position) int {
		q := make([]Position, 0)
		q = append(q, start)
		inBounds := func(pos Position) bool {
			return pos.x >= 0 && pos.y >= 0 && pos.x < len(inp) && pos.y < len(inp[0])
		}
		move := func(pos Position, dx int, dy int) Position {
			return Position{pos.x + dx, pos.y + dy, dx, dy}
		}
		vis := make([][][][]bool, 3)
		for i := 0; i < 3; i++ {
			vis[i] = make([][][]bool, 3)
			for j := 0; j < 3; j++ {
				vis[i][j] = make([][]bool, len(inp))
				for k := 0; k < len(inp); k++ {
					vis[i][j][k] = make([]bool, len(inp[0]))
				}
			}
		}
		for len(q) != 0 {
			pos := q[0]
			q = q[1:]
			if !inBounds(pos) || vis[pos.dx+1][pos.dy+1][pos.x][pos.y] {
				continue
			}
			vis[pos.dx+1][pos.dy+1][pos.x][pos.y] = true
			if inp[pos.x][pos.y] == '.' || (inp[pos.x][pos.y] == '|' && pos.dy == 0) || (inp[pos.x][pos.y] == '-' && pos.dx == 0) {
				q = append(q, move(pos, pos.dx, pos.dy))
			} else if inp[pos.x][pos.y] == '|' {
				q = append(q, move(pos, 1, 0))
				q = append(q, move(pos, -1, 0))
			} else if inp[pos.x][pos.y] == '-' {
				q = append(q, move(pos, 0, 1))
				q = append(q, move(pos, 0, -1))
			} else {
				if inp[pos.x][pos.y] == '/' {
					q = append(q, move(pos, -pos.dy, -pos.dx))
				} else if inp[pos.x][pos.y] == '\\' {
					q = append(q, move(pos, pos.dy, pos.dx))
				}
			}
		}
		ans := 0
		for i := 0; i < len(inp); i++ {
			for j := 0; j < len(inp[i]); j++ {
				found := false
				for x := 0; x < 3; x++ {
					for y := 0; y < 3; y++ {
						found = found || vis[x][y][i][j]
					}
				}
				if found {
					ans += 1
				}
			}
		}
		return ans
	}
	ans := 0
	for j := 0; j < len(inp[0]); j++ {
		ans = max(ans, get(Position{0, j, 1, 0}))
		ans = max(ans, get(Position{len(inp) - 1, j, -1, 0}))
	}
	for i := 0; i < len(inp); i++ {
		ans = max(ans, get(Position{i, 0, 0, 1}))
		ans = max(ans, get(Position{i, len(inp[0]) - 1, 0, -1}))
	}
	fmt.Println(ans)
}
