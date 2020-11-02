package letcode

//"abcabcbb"
func LengthOfLongestSubstring(s string) int {
	var currentlen, maxlen, noin int
	var m = map[byte]int{}
	for i, byt := range []byte(s) {
		value, ok := m[byt]
		if ok && (value > noin || noin == 0) {
			//value为上一次出现的索引 还要大于已经不考虑部分
			if currentlen > maxlen {
				maxlen = currentlen
			}
			//去掉上次出现位置以及更靠前的部分的继续
			currentlen = i - value
			m[byt] = i
			noin = value
		} else {
			currentlen++
			m[byt] = i
		}
	}
	if currentlen > maxlen {
		maxlen = currentlen
	}
	return maxlen
}
