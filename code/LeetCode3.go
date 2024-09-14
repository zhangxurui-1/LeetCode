package code

import "fmt"

type T struct {
}

func lengthOfLongestSubstring(s string) int {
	var (
		vis     = make(map[byte]int)
		ans int = 0
		j   int = 0
		i   int = 0
	)
	for i < len(s) {
		for ; j < len(s); j++ {
			v, exist := vis[s[j]]
			if exist && v > 0 {
				ans = max(ans, j-i)

				// 重定位i
				for i < j {
					vis[s[i]] = vis[s[i]] - 1
					if s[i] == s[j] {
						i++
						break
					}
					i++
				}

				break
			} else {
				vis[s[j]] = 1
			}
		}
		if j == len(s) {
			ans = max(ans, j-i)
			return ans
		}
	}

	return ans
}

func Test3() {
	var s string = "dvdf"
	fmt.Println(lengthOfLongestSubstring(s))
}
