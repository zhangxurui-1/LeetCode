package code

func minWindow(s string, t string) string {
	var (
		mp         = make(map[rune]int)
		record     = make(map[rune][]int)
		start  int = 0
		end    int = len(s) - 1
	)

	for _, c := range t {
		if _, exist := mp[c]; !exist {
			mp[c] = 1
		} else {
			mp[c]++
		}
	}

	for i, e := range s {
		if _, exist := mp[e]; exist {
			mp[e]--

		}
	}

}
