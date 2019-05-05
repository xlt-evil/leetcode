package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(addBinary("101111", "10"))
}

//二进制加法
func addBinary(a string, b string) string {
	as := strings.Split(a, "")
	bs := strings.Split(b, "")
	answer := ""
	alen := len(as)
	blen := len(bs)
	maxlength := func() int {
		if alen >= blen {
			return alen
		}
		return blen
	}()
	if alen > blen {
		tem := []string{}
		for i := 0; i < alen-blen; i++ {
			tem = append(tem, "0")
		}
		bs = append(tem, bs[:]...)
	} else if alen < blen {
		tem := []string{}
		for i := 0; i < blen-alen; i++ {
			tem = append(tem, "0")
		}
		as = append(tem, as[:]...)
	}
	count := 0
	for i := maxlength - 1; i >= 0; i-- {
		al, _ := strconv.Atoi(as[i])
		bl, _ := strconv.Atoi(bs[i])
		if al+bl+count >= 2 {
			//进位
			s := (al + bl + count) % 2
			nums := strconv.Itoa(s)
			answer = nums + answer
			count = 1
		} else {
			nums := strconv.Itoa(al + bl + count)
			answer = nums + answer
			count = 0
		}
	}
	if count != 0 {
		answer = "1" + answer
	}
	return answer
}
