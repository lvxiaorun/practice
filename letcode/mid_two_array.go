package letcode

//1. k=5 l=5 7 r=2 4  6 8 10  k-2=3
//2. k=3 l=5 7 r=4 6 8 10 k-1=2
//3. k=2 l=5 7 r=6 8 10 k-1=1
// 1 3 5 7
// 2 4 6 8 10
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//两个都不为空
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2

	k := l / 2 //如果l是奇数结果就是第k+1个 偶数第k k+1个
	var isOu = l%2 == 0
	if len(nums1) == 0 {
		if isOu {
			return float64(nums2[k]+nums2[k-1]) / 2
		} else {
			return float64(nums2[k])
		}
	}
	if len(nums2) == 0 {
		if isOu {
			return float64(nums1[k]+nums1[k-1]) / 2
		} else {
			return float64(nums1[k])
		}
	}

	if l == 2 {
		return (float64(nums1[0] + nums2[0])) / 2
	}
	//nums1 := nums1
	//nums2 := nums2
	for len(nums1) > 0 && len(nums2) > 0 {
		if k <= 1 {
			break
			//这时候能直接处理出来结果
		}
		lk := k / 2
		ll := len(nums1)
		lr := len(nums2)
		var lmid, rmid int
		var acceptll, acceptrl int
		if ll < lk {
			//没有这么多
			acceptll = ll - 1
		} else {
			acceptll = lk - 1
		}
		lmid = nums1[acceptll]

		if lr < lk {
			//没有这么多
			acceptrl = lr - 1
		} else {
			acceptrl = lk - 1
		}
		rmid = nums2[acceptrl]

		if rmid > lmid {
			//去掉左边的 并且改变k的值
			nums1 = nums1[acceptll+1:]
			k = k - acceptll - 1
		} else {
			nums2 = nums2[acceptrl+1:]
			k = k - acceptrl - 1
		}
		//fmt.Println(nums1, nums2, k)
	}
	//fmt.Println(k)
	//fmt.Println(nums1)
	//fmt.Println(nums2)
	//可能有两个原因导致结束 k=1 或者其中一个长度只有0了 也有可能刚好两种情况都存在 下面的逻辑可以涵盖

	if len(nums1) == 0 {
		if isOu {
			return float64(nums2[k-1]+nums2[k]) / 2
		} else {
			return float64(nums2[k])
		}
	}
	if len(nums2) == 0 {
		if isOu {
			return float64(nums1[k-1]+nums1[k]) / 2
		} else {
			return float64(nums1[k])
		}
	}

	//下面就是k=1的情况了 而且nums1和nums2都不是空
	//直接找出来接下来两个数
	var resultk, resultk1 int
	if nums2[0] > nums1[0] {
		resultk = nums1[0]
		nums1 = nums1[1:]
		if len(nums1) == 0 {
			resultk1 = nums2[0]
		} else {
			if nums2[0] < nums1[0] {
				resultk1 = nums2[0]
			} else {
				resultk1 = nums1[0]
			}
		}
	} else {
		resultk = nums2[0]
		nums2 = nums2[1:]
		if len(nums2) == 0 {
			resultk1 = nums1[0]
		} else {
			if nums1[0] < nums2[0] {
				resultk1 = nums1[0]
			} else {
				resultk1 = nums2[0]
			}
		}
	}
	if isOu {
		return float64(resultk+resultk1) / 2
	} else {
		return float64(resultk1)
	}
	return 0

}
