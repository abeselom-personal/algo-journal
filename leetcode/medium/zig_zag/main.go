package main

import (
	"strconv"
)

func main() {
	convert("PAYPALISHIRING", 3)
}

func convert(s string, numRows int) string {
	temp := make(map[string][]string)
	index := 1
	isDescending := true
	for _, val := range s {
		key := strconv.Itoa(index)
		temp[key] = append(temp[key], string(val))

		if isDescending {
			index++
		} else {
			index--
		}
		if index == numRows {
			isDescending = false
		}
		if index == 1 {
			isDescending = true
		}
	}
	var result string
	for i := 1; i <= numRows; i++ {
		key := strconv.Itoa(i)
		for _, char := range temp[key] {
			result += char
		}
	}
	return result
}
