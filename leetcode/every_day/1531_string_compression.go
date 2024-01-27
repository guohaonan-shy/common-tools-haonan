package every_day

//func getLengthOfOptimalCompression(s string, k int) int {
//	compressed := string(s[0])
//	pre := s[0]
//	cnt := 1
//	for i := 1; i < len(s); i++ {
//		if pre == s[i] {
//			cnt++
//		} else {
//			if cnt != 1 {
//				compressed += strconv.Itoa(cnt)
//			}
//			pre = s[i]
//			cnt = 1
//			compressed += string(pre)
//		}
//	}
//
//	if cnt != 1 {
//		compressed += strconv.Itoa(cnt)
//	}
//
//	return 0
//}
