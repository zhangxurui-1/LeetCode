package code

import (
	"fmt"
	"sort"
)

type Interval struct {
	left  int
	right int
}

func merge(intervals [][]int) [][]int {
	var (
		itvs = []Interval{}
		ans  = [][]int{}
	)
	for _, value := range intervals {
		itvs = append(itvs, Interval{left: value[0], right: value[1]})
	}

	sort.Slice(itvs, func(i, j int) bool {
		return itvs[i].left < itvs[j].left
	})

	for _, itv := range itvs {
		l := len(ans)
		if l == 0 || itv.left > ans[l-1][1] {
			ans = append(ans, []int{itv.left, itv.right})
		} else {
			ans[l-1][1] = max(ans[l-1][1], itv.right)
		}
	}

	return ans
}

func Test56() {
	var intervals = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}
