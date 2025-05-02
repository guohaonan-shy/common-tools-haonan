package every_day

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	pre := []byte{}
	for i := 0; i < len(num2); i++ {
		pre = append([]byte{'0'}, pre...)
		multiple := num2[i] - '0'
		multiStep := uint8(0)
		cur := make([]byte, 0)
		for j := len(num1) - 1; j >= 0; j-- {
			temp := (num1[j]-'0')*multiple + multiStep
			multiAns := temp % 10
			cur = append(cur, multiAns+'0')
			multiStep = temp / 10
		}
		if multiStep > 0 {
			cur = append(cur, multiStep+'0')
		}

		// sum
		pointer1, pointer2 := 0, 0
		addStep := uint8(0)
		for pointer1 < len(cur) && pointer2 < len(pre) {
			temp := (cur[pointer1] - '0') + (pre[pointer2] - '0') + addStep
			addAns := temp % 10

			pre[pointer2] = addAns + '0'
			addStep = temp / 10
			pointer1++
			pointer2++
		}

		for pointer1 < len(cur) {
			temp := cur[pointer1] - '0' + addStep
			ans := temp % 10
			pre = append(pre, ans+'0')
			pointer1++
			addStep = temp / 10
		}

		for pointer2 < len(pre) {
			temp := pre[pointer2] - '0' + addStep
			ans := temp % 10
			pre[pointer2] = ans + '0'
			pointer2++
			addStep = temp / 10
		}

		if addStep > 0 {
			pre = append(pre, addStep+'0')
		}
	}
	return reverseList(pre)
}

func reverseList(list []byte) string {
	left, right := 0, len(list)-1
	for left < right {
		list[left], list[right] = list[right], list[left]
		left++
		right--
	}
	return string(list)
}
