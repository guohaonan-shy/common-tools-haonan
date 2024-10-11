package every_day

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	for i := range nums2 {
		nums2[i] *= k
	}

	set := make(map[int]int64, 0)
	maximum := 0
	for i := range nums1 {
		set[nums1[i]] += 1
		if nums1[i] > maximum {
			maximum = nums1[i]
		}
	}

	var cnt int64 = 0
	record := map[int]int64{}
	for i := range nums2 {

		if record[nums2[i]] > 0 {
			cnt += record[nums2[i]]
			continue
		}

		var local int64 = 0
		cur := nums2[i]
		for cur <= maximum {
			if set[cur] > 0 {
				cnt += set[cur]
				local += set[cur]
			}
			cur += nums2[i]
		}
		record[nums2[i]] = local
	}
	return cnt
}
