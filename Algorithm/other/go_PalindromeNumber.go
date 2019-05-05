package main

//判断数字是不是回文数
//负数一定不是回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	result := 0
	temp := x
	for temp != 0 {
		result = result*10 + temp%10
		temp = temp / 10
	}
	return x == result
}

func main() {

}
