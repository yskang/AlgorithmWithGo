package leetcode

import "fmt"

func RepeatedSubstringPattern() {
	fmt.Println(repeatedSubstringPattern("abab"))
}

func repeatedSubstringPattern(str string) bool {
	문자열길이 := len(str)

	약수들:= make([]int, 0)

	약수들 = 약수구하기(문자열길이)

	fmt.Println(약수들)

	var 결과 bool

	for _, 약수 := range(약수들) {
		결과 = true
		for 기준점 := 0 ; 기준점 < 문자열길이 ; 기준점 += 약수 {
			fmt.Println("검사", 기준점)
			if str[0:약수] != str[기준점:기준점 +약수] {
				결과 = false
			}
		}
		if 결과 == true {
			return true
		}
	}

	return false
}

func 약수구하기(숫자 int) []int {
	약수들 := make([]int, 0)
	for 테스트숫자 := 1 ; 테스트숫자 <= 숫자/2 ; 테스트숫자++ {
		if 숫자 % 테스트숫자 == 0 {
			약수들 = append(약수들, 테스트숫자)
		}
	}

	return 약수들
}

