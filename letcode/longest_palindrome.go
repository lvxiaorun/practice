package letcode

import (
	"strings"
)

//1 2 3 4 5 6 7 8
func LongestPalindrome(s string) string {
	f := func(str string, length int, index int) string {
		//index从0开始
		var maxStr = string(str[index])
		left := 1
		right := 1
		for {
			if index == 0 {
				return maxStr
			}
			ll := index - left
			rl := index + right
			if ll >= 0 && rl < length {
				if str[ll] == str[rl] {
					maxStr = string(str[ll]) + maxStr + string(str[rl])
					left++
					right++
				} else {
					return maxStr
				}
			} else {
				//直接返回了
				return maxStr
			}
		}
		return maxStr
	}

	if len(s) == 0 {
		return ""
	}
	ss := []byte(s)
	sstr := []string{}
	for _, item := range ss {
		sstr = append(sstr, string(item))
	}
	dstr := strings.Join(sstr, " ")
	var result string
	for i := range dstr {
		fr := f(dstr, len(dstr), i)
		if len(strings.Replace(fr, " ", "", -1)) > len(result) {
			result = strings.Replace(fr, " ", "", -1)
		}
	}
	return result
}

//12345678
func LongestPalindrome2(s string) string {

	return ""
}
