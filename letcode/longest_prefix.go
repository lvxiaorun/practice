package letcode

import (
	"strings"
)

func LongestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	if strs[0] == "" {
		return ""
	}
	var cursor int
j:
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < l; j++ {
			if len(strs[j]) <= cursor {
				break j
			}
			if strs[j][i] != strs[0][i] {
				break j
			}
		}
		//等于的话就延长longest
		cursor = i

	}
	return strs[0][0:cursor]
}

func LongestCommonPrefix2(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	if strs[0] == "" {
		return ""
	}
	var cursor int //已经确认相同的数字个数
	l1 := len(strs[0])
	var higest = l1   //当前未确认的最大数字位数
	k := (l1 + 1) / 2 //初始化k k代表当前需要第几个数字
x:
	for cursor < higest {
		for i := 1; i < l; i++ {
			if !strings.HasPrefix(strs[i], strs[0][0:k]) {
				//有不包含的情况
				higest = k - 1                //k以及后面的直接被排除了
				k = (cursor + higest + 1) / 2 //+1是因为怕永远到不了higest
				//fmt.Println("tiao", "cursor", cursor, "highest", higest, "k", k)
				continue x
			}
		}
		//没有被break就说明k之前的全部一样
		cursor = k
		k = (cursor + higest + 1) / 2

		//fmt.Println("cursor", cursor, "highest", higest, "k", k)

	}
	return strs[0][0:cursor]
}
