package back_track

func permuteUnique(nums []int) [][]int {
	k := len(nums)
	temp := make([]int, 0)
	res := make([][]int, 0)
	state := make([]bool, k)

	var dfs func()
	dfs = func() {

		if len(temp) == k {
			ele := make([]int, k)
			copy(ele, temp)
			res = append(res, ele)
			return
		}

		hash := make(map[int]bool, 0)
		/*
			相比于46题的无重复元素数组，除了元素在调用路径内只能使用一次
			重复元素带来的问题 => 同一调用栈内相同元素不能被使用，需要在栈内引入map做去重判断
		*/

		for i := 0; i < k; i++ {
			if !state[i] && !hash[nums[i]] {
				state[i] = true
				hash[nums[i]] = true
				/* 为什么hash要放在里面？
				每次调用都需要遍历所有元素，去判断当前索引在当前调用path下是否使用(state) => 数组内某个重复元素在当前路径下已被选取一次，但其他该元素仍可被该路径调用
				如果放在if外面会出现 => 当前路径为[2,1]，在第三层调用时，先遍历第一个1时，因为第一个1已经被选取，所以跳过它的处理，但是hash由于在if外面，所以会添加到hash导致第二个1没办法选取。
				*/
				temp = append(temp, nums[i])
				dfs()
				temp = temp[:len(temp)-1]
				state[i] = false
			}
		}
	}
	dfs()
	return res
}
