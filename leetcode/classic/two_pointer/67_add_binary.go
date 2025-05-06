package two_pointer

func addBinary(a string, b string) string {
	p1, p2 := len(a)-1, len(b)-1
	var step uint8
	res := []byte{}
	for p1 >= 0 && p2 >= 0 {
		temp := (a[p1] - '0') + (b[p2] - '0') + step

		step = temp / 2
		ans := temp % 2
		res = append(res, ans+'0')
		p1--
		p2--
	}

	for ; p1 >= 0; p1-- {
		temp := (a[p1] - '0') + step
		step = temp / 2
		ans := temp % 2
		res = append(res, ans+'0')
	}

	for ; p2 >= 0; p2-- {
		temp := (b[p2] - '0') + step
		step = temp / 2
		ans := temp % 2
		res = append(res, ans+'0')
	}

	if step == 1 {
		res = append(res, 1+'0')
	}
	reverseBytes(res)
	return string(res)
}

func reverseBytes(bytes []byte) {
	left, right := 0, len(bytes)-1
	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}
	return
}
