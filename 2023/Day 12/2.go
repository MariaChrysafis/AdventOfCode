package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func stringToInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}
func stringArrayToIntArray(s []string) []int {
	arr := []int{}
	for _, x := range s {
		arr = append(arr, stringToInt(x))
	}
	return arr
}
func merge(a []int, b []int) []int {
	ans := []int{}
	for _, x := range a {
		ans = append(ans, x)
	}
	for _, x := range b {
		ans = append(ans, x)
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
	ans := 0
	for _, d := range inp {
		a := strings.Split(d, " ")
		s := a[0] + "?" + a[0] + "?" + a[0] + "?" + a[0] + "?" + a[0] + "."
		arr := stringArrayToIntArray(strings.Split(a[1], ","))
		arr = merge(merge(merge(arr, arr), merge(arr, arr)), arr)
		arr = append(arr, 30)
		dp := make([][][]int, len(s)+1)
		for i := 0; i < len(dp); i++ {
			dp[i] = make([][]int, len(arr))
			for j := 0; j < len(arr); j++ {
				dp[i][j] = make([]int, arr[j]+1)
			}
		}
		//dp[index in string][index in array][consecutive]
		dp[0][0][0] = 1
		for i := 1; i < len(dp); i++ {
			for j := 0; j < len(arr); j++ {
				for k := 0; k <= arr[j] && k <= i; k++ {
					dp[i][j][k] = 0
					period := func() {
						if k == 0 {
							dp[i][j][k] += dp[i-1][j][0]
							if j != 0 {
								dp[i][j][k] += dp[i-1][j-1][arr[j-1]]
							}
						}
					}
					hashtag := func() {
						if k != 0 {
							dp[i][j][k] += dp[i-1][j][k-1]
						}
					}
					if s[i-1] == '#' || s[i-1] == '?' {
						hashtag()
					}
					if s[i-1] == '.' || s[i-1] == '?' {
						period()
					}
				}
			}
		}
		for _, x := range dp[len(s)][len(arr)-1] {
			ans += x
		}
	}
	fmt.Println(ans)
}
