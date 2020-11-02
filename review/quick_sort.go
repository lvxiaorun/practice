package review

func Quick_sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	length := len(arr)
	//把第一个作为标志数
	cNum := arr[0]
	var left []int
	var right []int

	for i := 1; i < length; i++ {
		if cNum >= arr[i] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	//左边结果继续排序
	left = Quick_sort(left)
	right = Quick_sort(right)
	ret := append(append(left, cNum), right...)
	return ret
}

//func MergeSort(arr []int) ([]int) {
//	l := len(arr)
//	if l <= 1 {
//		return arr
//	}
//	mid := l/2
//	var left []int
//	var right []int
//	for i := range arr{
//		if i<mid{
//			left = append(left,arr[i])
//		}else {
//			right = append(right,arr[i])
//		}
//	}
//	left = MergeSort(left)
//	right = MergeSort(right)
//	var result []int
//	lleft := len(left)
//	lright := len(right)
//	i := lleft
//	j := lright
//	for i > 0 && j > 0{
//		if left[i-1] >= right[j-1]
//	}
//
//}
