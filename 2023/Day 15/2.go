package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func stringToInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

type Key struct {
	s string
	x int
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	get_hash := func(s string) int {
		ans := 0
		for _, c := range s {
			ans = ((ans + int(c)) * 17) % 256
		}
		return ans
	}
	input := strings.Split(string(content), ",")
	hashtable := make([][]Key, 256)
	for i := 0; i < 256; i++ {
		hashtable[i] = make([]Key, 0)
	}
	for _, s := range input {
		if strings.Count(s, "-") != 0 {
			s = s[:len(s)-1]
			hash := get_hash(s)
			for i, key := range hashtable[hash] {
				if key.s == s {
					hashtable[hash] = append(hashtable[hash][:i], hashtable[hash][i+1:]...)
					break
				}
			}
		} else {
			arr := strings.Split(s, "=")
			hash := get_hash(arr[0])
			found := false
			val := stringToInt(arr[1])
			for i, key := range hashtable[hash] {
				if key.s == arr[0] {
					hashtable[hash][i].x = val
					found = true
				}
			}
			if !found {
				hashtable[hash] = append(hashtable[hash], Key{arr[0], val})
			}
		}
	}
	ans := 0
	for j := range hashtable {
		for ind, key := range hashtable[j] {
			ans += (j + 1) * (ind + 1) * key.x
		}
	}
	fmt.Print(ans)
}
