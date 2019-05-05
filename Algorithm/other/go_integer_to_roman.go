package main

import (
	"fmt"
)

func intToRoman(num int) string {
	record := map[int]string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M", 4: "IV", 9: "IX", 40: "XL", 90: "XC", 400: "CD", 900: "CM"}
	key := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	length := len(key)
	result := ""
	roman := []string{}
	i := 1
	for num != 0 {
		a := num % (10 * i)
		num -= a
		i *= 10
		for a != 0 {
			for i := length - 1; i >= 0; i-- {
				if a-key[i] >= 0 {
					result += record[key[i]]
					a = a - key[i]
					break
				}
			}
		}
		roman = append(roman, result)
		result = ""
	}
	length = len(roman)
	for i := length - 1; i >= 0; i-- {
		result += roman[i]
	}
	return result
}

func main() {
	fmt.Println(intToRoman(68))
}
