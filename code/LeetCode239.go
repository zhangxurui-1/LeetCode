package code

import (
	"container/heap"
	"fmt"
)

/*
单调队列
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
*/

// 优先队列，大根堆
type ElementHeap []element

func (eh ElementHeap) Len() int {
	return len(eh)
}

func (eh ElementHeap) Less(i, j int) bool {
	return eh[i].value > eh[j].value
}

func (eh ElementHeap) Swap(i, j int) {
	eh[i], eh[j] = eh[j], eh[i]
}

func (eh *ElementHeap) Push(x interface{}) {
	*eh = append(*eh, x.(element))
}

func (eh *ElementHeap) Pop() interface{} {
	if len(*eh) == 0 {
		return nil
	}
	x := (*eh)[len(*eh)-1]
	*eh = (*eh)[:len(*eh)-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	var (
		eh      = &ElementHeap{}
		ans     = []int{}
		i   int = 0
	)
	for ; i < k; i++ {
		*eh = append(*eh, element{value: nums[i], loc: i})
	}
	heap.Init(eh)
	ans = append(ans, (*eh)[0].value)

	for ; i < len(nums); i++ {
		heap.Push(eh, element{value: nums[i], loc: i})
		for i-(*eh)[0].loc >= k {
			heap.Pop(eh)
		}
		ans = append(ans, (*eh)[0].value)
	}

	return ans
}

func Test239() {
	var nums = []int{1, 3, 1, 2, 0, 5}
	fmt.Println(maxSlidingWindow(nums, 3))
}
