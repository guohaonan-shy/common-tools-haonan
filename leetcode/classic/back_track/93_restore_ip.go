package back_track

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	temp := make([]string, 0)
	res := make([]string, 0)
	var dfs func(cur int)

	dfs = func(cur int) {
		if len(temp) == 4 && cur == len(s) {
			ip := strings.Join(temp, ".")
			res = append(res, ip)
			return
		}
		// 这个终止条件: 1. 组成了一个合理ip，但是字符没有使用完成 2. 仍处于ip组成过程中，但是无可用字符
		if (len(temp) == 4 && cur < len(s)) || cur >= len(s) {
			return
		}

		if s[cur] == '0' {
			temp = append(temp, "0")
			dfs(cur + 1)
			temp = temp[:len(temp)-1]
			return
		}

		code := 0
		for i := cur; i < cur+3 && i < len(s); i++ {
			code = code*10 + int(s[i]-'0')
			if code > 255 {
				break
			}

			temp = append(temp, strconv.Itoa(code))
			dfs(i + 1)
			temp = temp[:len(temp)-1]
		}
	}

	dfs(0)
	return res
}
