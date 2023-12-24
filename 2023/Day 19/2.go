package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type Part struct {
	X int
	M int
	A int
	S int
}

func getField(v *Part, field string) int {
	return int(reflect.Indirect(reflect.ValueOf(v)).FieldByName(strings.ToUpper(field)).Int())
}

func setField(obj any, field string, value any) {
	reflect.Indirect(reflect.ValueOf(obj)).FieldByName(strings.ToUpper(field)).Set(reflect.ValueOf(value))
}

func StringToInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func cost(part_lower Part, part_higher Part) int {
	prod := 1
	for _, val := range []string{"x", "m", "a", "s"} {
		if getField(&part_lower, val) <= getField(&part_higher, val) {
			prod *= getField(&part_higher, val) - getField(&part_lower, val) + 1
		} else {
			return 0
		}
	}
	return prod
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
	myMap := make(map[string]string)
	for _, x := range inp {
		if len(x) != 0 && x[0] != '{' {
			arr := strings.Split(x[:len(x)-1], "{")
			myMap[arr[0]] = arr[1]
		}
	}
	var dfs func(Part, Part, string) int
	dfs = func(partLower Part, partHigher Part, workflow string) int {
		switch workflow {
		case "A":
			return cost(partLower, partHigher)
		case "R":
			return 0
		}
		res := strings.Split(myMap[workflow], ",")
		ans := 0
		for i, y := range res {
			if i == len(res)-1 {
				ans += dfs(partLower, partHigher, y)
			} else {
				arr := strings.Split(y, ":")
				field := y[0:1]
				value := StringToInt(arr[0][2:])
				new_part_lower, new_part_higher := partLower, partHigher
				if strings.Count(arr[0], ">") > 0 {
					setField(&partHigher, field, min(value, getField(&partHigher, field)))
					setField(&new_part_lower, field, max(value+1, getField(&new_part_lower, field)))
				} else {
					setField(&partLower, field, max(value, getField(&partLower, field)))
					setField(&new_part_higher, field, min(value-1, getField(&new_part_higher, field)))
				}
				ans += dfs(new_part_lower, new_part_higher, arr[1])
			}
		}
		return ans
	}
	fmt.Println(dfs(Part{1, 1, 1, 1}, Part{4000, 4000, 4000, 4000}, "in"))
}

func min(val int, value int) int {
	if val < value {
		return val
	}
	return value
}

func max(val int, value int) int {
	return -min(-val, -value)
}
