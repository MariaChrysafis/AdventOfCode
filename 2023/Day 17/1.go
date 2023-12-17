package main

import (
	"fmt"
	pq "gopkg.in/dnaeon/go-priorityqueue.v1"
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
	inp := make([][]int, 1)
	for i := range input {
		if string(input[i]) == "\n" {
			inp = append(inp, make([]int, 0))
		} else {
			inp[len(inp)-1] = append(inp[len(inp)-1], int(input[i])-int('0'))
		}
	}
	inBounds := func(position Position) bool {
		return position.x >= 0 && position.y >= 0 && position.x < len(inp) && position.y < len(inp[0])
	}
	q := pq.New[Position, int64](pq.MinHeap)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if max(-dx, dx)+max(-dy, dy) == 1 {
				q.Put(Position{0, 0, dx, dy}, 0)
			}
		}
	}
	vis := make([][][][]int64, len(inp))
	for i := 0; i < len(inp); i++ {
		vis[i] = make([][][]int64, len(inp[0]))
		for j := 0; j < len(inp[0]); j++ {
			vis[i][j] = [][]int64{{-1, -1, -1}, {-1, -1, -1}, {-1, -1, -1}}
		}
	}
	for q.Len() != 0 {
		tot := q.Get()
		pos, cost := tot.Value, tot.Priority
		if vis[pos.x][pos.y][pos.dx+1][pos.dy+1] != -1 {
			continue
		}
		vis[pos.x][pos.y][pos.dx+1][pos.dy+1] = cost
		for t := 0; t < 2; t++ {
			dx, dy := pos.dy*(2*t-1), pos.dx*(2*t-1)
			c := int64(0)
			for l := 1; l <= 3; l++ {
				next := Position{pos.x + dx*l, pos.y + dy*l, dx, dy}
				if !inBounds(next) {
					break
				}
				c += int64(inp[pos.x+dx*l][pos.y+dy*l])
				q.Put(next, c+cost)
			}
		}
	}
	fmt.Println(min(vis[len(inp)-1][len(inp[0])-1][2][1], vis[len(inp)-1][len(inp[0])-1][1][2]))
}
