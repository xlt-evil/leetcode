package main

import (
	"fmt"
	"math"
	"strings"
)

//最长回文子串
//Manacher 算法 马拉车

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	sl := make([]rune, 0)
	for _, v := range s {
		sl = append(sl, v)
	}
	max := -1
	length, str := getstrlength(sl)
	p := make([]int, length) //记录当前的回文长度
	var id, mx int
	for i := 1; i < length; i++ {
		if mx > i {
			p[i] = int(math.Min(float64(mx-i), float64(p[2*id-i]))) //找到i的对称坐标j  s = i - id ;j=id-s ===>> j = 2*id - i
		} else {
			p[i] = 1 //以当前字符为中心寻找
		}
		for str[i-p[i]] == str[i+p[i]] {
			p[i]++
		}
		if mx < i+p[i] {
			mx = i + p[i] //设置最远的回文标识
			id = i        //设置最远回文中心
		}
		max = int(math.Max(float64(max), float64(p[i]-1)))
	}
	index := Max(p)
	sub := string(str[index-max : index+max+1])
	substring := strings.Replace(sub, "#", "", -1)
	substring = strings.Replace(substring, "@", "", -1)
	return substring
}

func getstrlength(s []rune) (int, []rune) {
	sl := make([]rune, 0)
	sl = append(sl, '$')
	sl = append(sl, '#')
	j := 2
	for i := 0; i < len(s); i++ {
		sl = append(sl, s[i])
		j++
		sl = append(sl, '#')
		j++
	}
	sl = append(sl, '@')
	return j, sl
}

func Max(p []int) int {
	max := 0
	index := 0
	for i, _ := range p {
		if max < p[i] {
			max = p[i]
			index = i
		}
	}
	return index
}

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}
