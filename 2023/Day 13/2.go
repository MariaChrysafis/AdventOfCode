package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

func transpose(arr []string) []string {
	ans := make([]string, len(arr[0]))
	for j := 0; j < len(arr); j++ {
		for i := 0; i < len(arr[j]); i++ {
			ans[i] += string(arr[j][i])
		}
	}
	return ans
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
	arr := make([][]string, 1)
	for _, x := range inp {
		if len(x) == 0 {
			arr = append(arr, make([]string, 0))
		} else {
			arr[len(arr)-1] = append(arr[len(arr)-1], x)
		}
	}
	ans := 0
	for _, g := range arr {
		lineOfReflection := func(grid []string) []int {
			res := []int{}
			for t := 1; t <= 100; t += 99 {
				grid = transpose(grid)
				for i := 1; i < 2*len(grid)-2; i += 2 {
					//reflection over i?
					okay := true
					for j := range grid {
						if i-j >= 0 && i-j < len(grid) && grid[j] != grid[i-j] {
							okay = false
						}
					}
					if okay {
						res = append(res, t*(i+1)/2)
					}
				}
			}
			return res
		}
		found := false
		for i := range g {
			for j := range g[i] {
				if found {
					continue
				}
				orig := lineOfReflection(g)
				for t := 0; t < 2; t++ {
					if g[i][j] == '#' {
						g[i] = g[i][:j] + "." + g[i][j+1:]
					} else {
						g[i] = g[i][:j] + "#" + g[i][j+1:]
					}
					if t == 0 && len(lineOfReflection(g)) != 0 && !reflect.DeepEqual(orig, lineOfReflection(g)) {
						if orig[0] == lineOfReflection(g)[0] {
							ans += lineOfReflection(g)[1]
						} else {
							ans += lineOfReflection(g)[0]
						}
						found = true
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
