package letcode

/*[1,  5,  9],
  [10, 11, 13],
  [12, 13, 15]*/
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	//因为是n*n的 所以行和列是一样的
	if n == 0 {
		return 0
	}
	//二分法查找

}
