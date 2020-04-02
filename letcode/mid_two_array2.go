package letcode

import "fmt"

func FindMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	//两个都不为空
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2

	k := l/2 + 1 //如果l是奇数结果就是第k个 偶数第k k-1个
	var isOu = l%2 == 0
	if len(nums1) == 0 {
		if isOu {
			return float64(nums2[k-2]+nums2[k-1]) / 2
		} else {
			return float64(nums2[k-1])
		}
	}
	if len(nums2) == 0 {
		if isOu {
			return float64(nums1[k-2]+nums1[k-1]) / 2
		} else {
			return float64(nums1[k-1])
		}
	}

	if l == 2 {
		return (float64(nums1[0] + nums2[0])) / 2
	}
	left := nums1
	right := nums2
	vkpre := 0
	for {
		if k <= 1 {
			break
			//这时候能直接处理出来结果
		}
		lk := k / 2
		ll := len(left)
		lr := len(right)
		//如果有任何一个长度为0
		if len(left) == 0 {
			if right[k-2] > vkpre{
				vkpre = right[k-2]
			}
			right = right[k-1:]
			k = 1
			break
		}

		if len(right) == 0 {
			if left[k-2] > vkpre{
				vkpre = left[k-2]
			}
			left = left[k-1:]
			k = 1
			break
		}
		var lmid, rmid int
		var acceptll, acceptrl int
		if ll < lk {
			//没有这么多
			acceptll = ll - 1
		} else {
			acceptll = lk - 1
		}
		lmid = left[acceptll]

		if lr < lk {
			//没有这么多
			acceptrl = lr - 1
		} else {
			acceptrl = lk - 1
		}
		rmid = right[acceptrl]

		if rmid > lmid {
			//去掉左边的 并且改变k的值
			//vkpre = lmid 第二次裁剪的可能比第一次的小所以要判断
			if lmid > vkpre {
				vkpre = lmid
			}
			left = left[acceptll+1:]
			k = k - acceptll - 1
		} else {
			if rmid > vkpre {
				vkpre = rmid
			}
			right = right[acceptrl+1:]
			k = k - acceptrl - 1
		}
		//fmt.Println(left, right, k)
		fmt.Println("k", k)
		fmt.Println("left", left)
		fmt.Println("right", right)
		fmt.Println("vkpre", vkpre)
		fmt.Println("-----")
	}
	//可能有两个原因导致结束 k=1 或者其中一个长度只有0了 也有可能刚好两种情况都存在 下面的逻辑可以涵盖

	//fmt.Println(vkpre)
	if len(left) == 0 {
		if isOu {
			return float64(right[k-1]+vkpre) / 2
		} else {
			return float64(right[k-1])
		}
	}
	if len(right) == 0 {
		if isOu {
			return float64(left[k-1]+vkpre) / 2
		} else {
			return float64(left[k-1])
		}
	}

	//下面就是k=1的情况了 而且left和right都不是空
	//直接找出来接下来两个数
	var resultk int
	if right[0] > left[0] {
		resultk = left[0]
	} else {
		resultk = right[0]
	}
	if isOu {
		return float64(resultk+vkpre) / 2
	} else {
		return float64(resultk)
	}
	return 0

}
