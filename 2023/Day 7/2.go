package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func deepCopy(s string) string {
	var sb strings.Builder
	sb.WriteString(s)
	return sb.String()
}
func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)
	arr := make([]string, 1)
	for i := range input {
		if string(input[i]) == "\n" {
			arr = append(arr, "")
		} else {
			arr[len(arr)-1] += string(input[i])
		}
	}
	fmt.Println(arr)
	get_type1 := func(s string) int {
		cnt := make(map[string]int)
		for _, c := range s {
			cnt[string(c)] += 1
		}
		mx_value := 0
		for _, value := range cnt {
			mx_value = max(mx_value, value)
		}
		if mx_value == 5 {
			return 1
		} else if mx_value == 4 {
			return 2
		} else if mx_value == 3 && len(cnt) == 2 {
			return 3
		} else if mx_value == 3 {
			return 4
		} else if mx_value == 2 && len(cnt) == 3 {
			return 5
		} else if mx_value == 2 {
			return 6
		}
		return 7
	}
	get_type := func(s string) int {
		ans := 10000
		for c := '2'; c <= 'Z'; c++ {
			orig := strings.ReplaceAll(deepCopy(s), "J", string(c))
			ans = min(ans, get_type1(orig))
		}
		return ans
	}
	get_strength := func(s string) int {
		cards := "AKQT9876543210J"
		for i := 0; i < len(cards); i++ {
			if string(cards[i]) == s {
				return i
			}
		}
		return -1
	}
	sort.Slice(arr, func(i int, j int) bool {
		key1 := arr[i][:5]
		key2 := arr[j][:5]
		if get_type(key1) == get_type(key2) {
			for ind := range key1 {
				if key1[ind] != key2[ind] {
					return get_strength(string(key1[ind])) > get_strength(string(key2[ind]))
				}
			}
		}
		return get_type(key1) > get_type(key2)
	})
	ans := 0
	for ind, _ := range arr {
		val, err := strconv.Atoi(arr[ind][6:])
		if err != nil {
			panic(err)
		}
		ans += val * (ind + 1)
	}
	for _, x := range arr {
		fmt.Println(x, get_type(x[:5]))
	}
	fmt.Println(ans)
}
