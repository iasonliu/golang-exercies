package main

import "fmt"

func main() {
	arrList := []int{3, 5, 6, 3, 2, 1, 4}
	// [14, 3, 2, 5]
	var rs []int
	var tmp int
	for i, v := range arrList {
		fmt.Println(i, v, tmp, rs)
		if v < arrList[i+1] {
			tmp += v
		} else {
			rs = append(rs, tmp)
			tmp = v
		}
	}
}
