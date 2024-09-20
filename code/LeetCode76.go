package code

import (
	"fmt"
)

func minWindow(s string, t string) string {
	var (
		sByte = []byte(s)
		tByte = []byte(t)
		mp    = make(map[byte]int)

		start int    = 0
		exist bool   = false
		i     int    = 0
		ans   string = ""
	)

	for _, c := range tByte {
		if _, exist := mp[c]; !exist {
			mp[c] = 1
		} else {
			mp[c]++
		}
	}

	for _, exist = mp[sByte[start]]; !exist && start < len(sByte); start++ {
		_, exist = mp[sByte[start]]
	}

	for i = start; i < len(sByte); i++ {
		c := sByte[i]
		_, exist = mp[c]
		if exist {
			mp[c]--
			flag := true && mp[c] <= 0
			if mp[c] <= 0 {
				for _, v := range mp {
					flag = flag && v <= 0
				}
			}
			if flag {
				ansByte := sByte[start : i+1]
				ans = string(ansByte)
				break
			}
		}
	}

	i++
	for ; i < len(sByte); i++ {
		c := sByte[i]
		_, exist = mp[c]
		if exist {
			mp[c]--
			if mp[c] < 0 {
				for start < len(sByte) {
					if v, ok := mp[sByte[start]]; ok && v >= 0 {
						break
					} else if ok {
						mp[sByte[start]]++
					}
					start++
				}
			}
			if len(ans) > i-start+1 {
				ansByte := sByte[start : i+1]
				ans = string(ansByte)
			}
		}
	}

	return ans
}

func Test76() {
	var (
		s string = "a"
		t string = "aa"
	)
	fmt.Println(minWindow(s, t))
}
