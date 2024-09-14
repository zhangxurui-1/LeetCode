package code

import "fmt"

func findAnagrams(s string, p string) []int {
	var (
		pattern = make(map[byte]int)

		i   int = 0
		j   int = 0
		ans     = []int{}
	)

	for i = 0; i < len(p); i++ {
		if _, exist := pattern[p[i]]; !exist {
			pattern[p[i]] = 1
		} else {
			pattern[p[i]]++
		}
	}

	i = 0
	for i < len(s) {
		if i+len(p) > len(s) {
			break
		}

		for ; j < i+len(p); j++ {
			if c, exist := pattern[s[j]]; !exist {
				// 定位i=j+1
				for ; i <= j; i++ {
					if _, exist1 := pattern[s[i]]; exist1 {
						pattern[s[i]]++
					}
				}
				j = i
				break
			} else if c == 0 {
				for s[i] != s[j] {
					pattern[s[i]]++
					i++
				}
				pattern[s[i]]++
				i++
				break
			} else {
				pattern[s[j]]--
			}
		}

		if j == i+len(p) {
			flag := true
			for _, v := range pattern {
				flag = flag && (v == 0)
			}
			if flag {
				ans = append(ans, i)
			}
			pattern[s[i]]++
			i++
		}

	}
	return ans
}

func Test438() {
	var (
		s string = "abab"
		p string = "ab"
	)

	fmt.Println(findAnagrams(s, p))

}
