package letcode

func MinPathSum(grid [][]int) int{
	m := len(grid)
	if m == 0{
		return 0
	}
	n := len(grid[0])
	if n == 0{
		return 0
	}

	for i := 0;i<m;i++{
		for j := 0;j<n;j++{
			//比较左边和上面 如果没有取自己 有一个加上一个 两个比较后加上小的
			if i==0 && j == 0{
				continue
			}else if i == 0 && j != 0{
				grid[i][j] = grid[i][j] + grid[i][j-1]
			}else if i != 0 && j == 0{
				grid[i][j] = grid[i][j] + grid[i-1][j]
			}else {
				//两个都不是0
				grid[i][j] = grid[i][j] + min(grid[i-1][j],grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}

func min(a,b int) int{
	if a > b{
		return b
	}
	return a
}
