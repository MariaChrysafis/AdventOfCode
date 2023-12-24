package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type Part struct {
	x int
	m int
	a int
	s int
}

func getField(v *Part, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}

func StringToInt(s string) int {
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
	m := make(map[string]string)
	ans := 0
	for _, x := range inp {
		if len(x) == 0 {
			continue
		}
		if x[0] != '{' {
			arr := strings.Split(x[:len(x)-1], "{")
			m[arr[0]] = arr[1]
		} else {
			arr := strings.Split(x[1:len(x)-1], ",")
			values := []int{}
			for _, y := range arr {
				values = append(values, StringToInt(y[2:]))
			}
			part := Part{values[0], values[1], values[2], values[3]}
			workflow := "in"
			for workflow != "A" && workflow != "R" {
				res := strings.Split(m[workflow], ",")
				for i, y := range res {
					if i == len(res)-1 {
						workflow = y
					} else {
						clause := strings.Split(y, ":")
						val := getField(&part, clause[0][0:1])
						if clause[0][1] == '<' && val < StringToInt(clause[0][2:]) {
							workflow = clause[1]
							break
						} else if clause[0][1] == '>' && val > StringToInt(clause[0][2:]) {
							workflow = clause[1]
							break
						}
					}
				}
			}
			if workflow == "A" {
				ans += part.a + part.x + part.s + part.m
			}
		}
	}
	fmt.Println(ans)
}
