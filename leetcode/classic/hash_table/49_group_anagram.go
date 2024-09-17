package hash_table

import (
	"sort"
)

// 假设字符数组长度个数为n，数组内最长的字符长为k，时间复杂度为O(N*KlogK)
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)

	dict := make(map[string][]string, 0)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})

		if _, ok := dict[string(s)]; ok {
			dict[string(s)] = append(dict[string(s)], str)
		} else {
			dict[string(s)] = []string{str}
		}

	}

	for _, value := range dict {
		res = append(res, value)
	}
	return res

}

// array数组可以比较大小，即可以作为map的key：
// 数组比较大小原则：1.数组长度必须相等 2.数组的元素类型必须支持比较操作。比较：对于相同位置上的元素，它们的值必须相等；否则不等
func groupAnagrams_withoutSort(strs []string) [][]string {
	dict := make(map[[26]int][]string, 0)

	for _, str := range strs {
		cnt := [26]int{}
		for i := range str {
			cnt[str[i]-'a']++
		}

		if _, ok := dict[cnt]; ok {
			dict[cnt] = append(dict[cnt], str)
		} else {
			dict[cnt] = []string{str}
		}
	}

	res := make([][]string, 0)
	for _, value := range dict {
		res = append(res, value)
	}
	return res
}

func isHappy(n int) bool {

	dict := make(map[int]struct{}, 0)
	for {
		squareSum := 0
		for i := 1000000000; i > 0; i = i / 10 {
			res, remain := n/i, n%i

			if res == 0 {
				continue
			} else {
				squareSum += res * res
				n = remain
			}
		}

		if squareSum == 1 {
			return true
		}

		if _, ok := dict[squareSum]; ok {
			return false
		} else {
			dict[squareSum] = struct{}{}
			n = squareSum
		}
	}
}

func groupAnagrams_v2_sorted(strs []string) [][]string {
	set := make(map[string][]string, 0)

	for _, s := range strs {

		bytes := []byte(s)

		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})

		if _, ok := set[string(bytes)]; ok {
			set[string(bytes)] = append(set[string(bytes)], s)
		} else {
			set[string(bytes)] = []string{s}
		}
	}

	res := make([][]string, 0)
	for _, val := range set {
		res = append(res, val)
	}
	return res
}

func groupAnagrams_v2_compared(strs []string) [][]string {
	set := make(map[[26]int][]string, 0)
	for _, s := range strs {
		cnts := [26]int{}

		for i := range s {
			cnts[s[i]-'a']++
		}

		if _, ok := set[cnts]; ok {
			set[cnts] = append(set[cnts], s)
		} else {
			set[cnts] = []string{s}
		}
	}

	res := make([][]string, 0)
	for _, values := range set {
		res = append(res, values)
	}
	return res
}
