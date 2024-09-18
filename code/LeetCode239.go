package code

import "fmt"

type Stack struct {
	data   []int
	top    int
	bottom int
	length int
}

func (st *Stack) Reset(cap int) {
	st.data = make([]int, cap)
	st.top = -1
	st.bottom = 0
	st.length = 0
}

func (st *Stack) Push(value int) {
	st.top++
	st.data[st.top] = value
	st.length++
}

func (st *Stack) Pop() int {
	if st.length == 0 {
		return -1
	}
	var value int = st.data[st.top]
	st.top--
	st.length--
	return value
}

func (st *Stack) BottomPop() int {
	if st.length == 0 {
		return -1
	}
	var value int = st.data[st.bottom]
	st.bottom++
	st.length--
	return value
}

func maxSlidingWindow(nums []int, k int) []int {
	var (
		stack Stack
		index int = 0
		i     int = 0
		ans   []int
	)
	stack.Reset(len(nums))

	// init
	for i = 0; i < k; i++ {
		if stack.length == 0 {
			stack.Push(i)
		} else {
			index = stack.data[stack.top]
			for nums[index] < nums[i] {
				stack.Pop()
				if stack.length == 0 {
					break
				}
				index = stack.data[stack.top]
			}
			stack.Push(i)
		}

	}

	stack.bottom = stack.top - stack.length + 1
	maxIndex := stack.data[stack.bottom]
	ans = append(ans, nums[maxIndex])

	for ; i < len(nums); i++ {
		if i-stack.data[stack.bottom] >= k {
			stack.BottomPop()
		}

		if stack.length == 0 {
			stack.Push(i)
		} else {
			index = stack.data[stack.top]
			for nums[index] < nums[i] {
				stack.Pop()
				if stack.length == 0 {
					break
				}
				index = stack.data[stack.top]
			}
			stack.Push(i)
		}

		stack.bottom = stack.top - stack.length + 1
		maxIndex := stack.data[stack.bottom]
		ans = append(ans, nums[maxIndex])
	}
	return ans
}

func Test239() {
	var nums = []int{1, 3, 1, 2, 0, 5}
	fmt.Println(maxSlidingWindow(nums, 3))
}
