// https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&envId=top-100-liked

package code

import "fmt"

func groupAnagrams(strs []string) [][]string {
	var (
		t       = [26]int{}
		ans     = [][]string{}
		k   int = 0
	)

	var m = make(map[[26]int]int)

	for _, s := range strs {
		for j := 0; j < len(s); j++ {
			t[s[j]-'a']++
		}
		if v, ok := m[t]; ok {
			ans[v] = append(ans[v], s)
		} else {
			m[t] = k
			ans = append(ans, []string{})
			ans[k] = append(ans[k], s)
			k++
		}

		for j := 0; j < len(t); j++ {
			t[j] = 0
		}
	}

	return ans
}

func Test2() {
	var (
		strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	)

	fmt.Println(groupAnagrams(strs))
}
