package bit

func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ { // 按照32-bit的格式计算每一位的值，所有出现3次的数字每一位数字的和一定是3的倍数，那么唯一的一个数字就是除以3的余数
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}

		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
