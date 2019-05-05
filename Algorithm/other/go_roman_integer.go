package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	record := map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}
	str := strings.Split(s, "")
	sum := 0
	length := len(str)
	for i := 0; i < length; i++ {
		if str[i] == "I" || str[i] == "X" || str[i] == "C" {
			if i+1 < length {
				if v, ok := record[str[i]+str[i+1]]; ok {
					i++
					sum += v
					continue
				}
			}
		}
		sum += record[str[i]]
	}
	return sum
}

func main() {
	fmt.Println(romanToInt("CMI"))
}
