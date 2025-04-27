package hash_table

/*
	这个题目如果对空间复杂度没有要求的话，那么我们可以通过引入hashmap的方法进行求解。 第一个缺失的正整数一定是在 1～n+1 中间，因此当我们把所有元素放到hashmap内，然后从1~n+1判断是否存在即可。
	那么为什么第一个缺失值一定在1～n+1 ？
	1. 假设数组是[1,2,3,..., n]，那么第一个缺失的值是n+1
	2. 假设数组是[1,3,..., n, n+1]，那么第一个缺失的值是2
	3. 假设数组元素全部是负值，那么第一个缺失值就是1
	总之，我们可以将元素看做1～n以内的以及1～n以外的，只要存在至少一个1～n以外的元素，那么目标值一定是在1～n内。最极端数组内没有一个1～n范围外的元素，那目标值是n+1。
*/

func firstMissingPositive(nums []int) int {
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = len(nums) + 1
		}
	}

	for i := range nums {
		if abs(nums[i]) > len(nums) {
			continue
		}

		nums[abs(nums[i])-1] = -abs(nums[abs(nums[i])-1])
	}

	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
