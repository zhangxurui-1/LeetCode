package code

import "fmt"

/*
// 双指针
func trap(height []int) int {
	var (
		i           int = 0
		j           int = len(height) - 1
		leftHeight  int = 0
		rightHeight int = 0
		lastHeight  int = 0
		ans         int = 0
	)

	// 初始定位i,j
	for ; i < len(height)-1 && height[i] <= height[i+1]; i++ {
	}

	for ; j > 0 && height[j] <= height[j-1]; j-- {
	}

	for i < j {

		leftHeight = height[i]
		rightHeight = height[j]

		if i >= j {
			break
		}

		lastHeight = min(leftHeight, rightHeight)
		for k := i + 1; k < j; k++ {
			ans += lastHeight - min(lastHeight, height[k])
			height[k] = max(lastHeight, height[k])
		}

		if leftHeight >= rightHeight {
			for ; j > 0 && height[j] <= height[j-1]; j-- {
			}
		}
		if leftHeight <= rightHeight {
			for ; i < len(height)-1 && height[i] <= height[i+1]; i++ {
			}
		}
	}

	return ans
}
*/

// dp
/*
func trap(height []int) int {
	var (
		L              int = len(height)
		leftMaxHeight      = make([]int, L)
		rightMaxHeight     = make([]int, L)
		ans            int = 0
	)

	leftMaxHeight[0] = 0
	rightMaxHeight[L-1] = 0
	for i := 1; i < len(height); i++ {
		leftMaxHeight[i] = max(leftMaxHeight[i-1], height[i-1])
	}
	for i := L - 2; i >= 0; i-- {
		rightMaxHeight[i] = max(rightMaxHeight[i+1], height[i+1])
	}

	for i := 0; i < L; i++ {
		ans += max(min(leftMaxHeight[i], rightMaxHeight[i])-height[i], 0)
	}

	return ans
}
*/

// 单调栈
/*
func trap(height []int) int {
	var (
		stack      = make([]int, len(height))
		index  int = 0
		ans    int = 0
		left   int = 0
		bottom int = 0
	)
	for i, h := range height {
		if index < 2 {
			if index == 0 || h <= height[stack[index-1]] {
				// push
				stack[index] = i
				index++
			} else {
				stack[index-1] = i
			}
		} else {
			if h <= height[stack[index-1]] {
				// push
				stack[index] = i
				index++
			} else {
				for h > height[stack[index-1]] && index >= 2 {
					// pop
					left = stack[index-2]
					bottom = stack[index-1]
					ans += (min(h, height[left]) - height[bottom]) * (i - left - 1)
					index--
				}

				// push
				if index < 2 && h > height[stack[index-1]] {
					stack[index-1] = i
				} else {
					// push
					stack[index] = i
					index++
				}
			}
		}
	}
	return ans
}
*/

// 双指针，思想是dp
func trap(height []int) int {
	var (
		l              int = 0
		r              int = len(height) - 1
		leftMaxHeight  int = 0
		rightMaxHeight int = 0
		ans            int = 0
	)
	for l < r {
		leftMaxHeight = max(height[l], leftMaxHeight)
		rightMaxHeight = max(height[r], rightMaxHeight)
		if leftMaxHeight < rightMaxHeight {
			ans += leftMaxHeight - height[l]
			l++
		} else {
			ans += rightMaxHeight - height[r]
			r--
		}
	}
	return ans
}

func Test42() {
	var height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
