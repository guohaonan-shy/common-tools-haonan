package binary_search

// 因为两个数组是有序的，所以中位数元素的位置可以事先知道，问题转换为求解数组中第k小的元素
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 明确中位数元素的位置
	var median1, median2 = -1, -1
	length1, length2 := len(nums1), len(nums2)

	if (length1+length2)%2 == 1 {
		median1 = (length1 + length2 + 1) / 2
	} else {
		median1 = (length1 + length2) / 2
		median2 = median1 + 1
	}

	res1 := findKSmallestElement(nums1, nums2, median1)
	var res2 = res1
	if median2 != -1 {
		res2 = findKSmallestElement(nums1, nums2, median2)
	}

	return float64(res1+res2) / 2

}

func findKSmallestElement(nums1 []int, nums2 []int, k int) int {

	left, right := 0, 0

	for left < len(nums1) && right < len(nums2) && k > 1 {

		var leftIdx, rightIdx int
		var removeK int
		// 因为我们需要找的是第k小的元素，所以我们要排除第k/2小的元素，通过二分不断逼近
		// 此处有个case，当时两个数组长度差距较大时，即一个数组并没有k/2个元素，那么直接跳到最后一个元素进行比较

		// 当A[k/2-1] 小于 B[k/2-1]，即比A[k/2−1]小的数最多只有A的前 k/2−1个数和B的前 k/2−1个数，即比 A[k/2−1]小的数最多只有k−2个，因此A[k/2−1]不可能是第k个数
		leftIdx = min(left+k/2, len(nums1)) - 1
		rightIdx = min(right+k/2, len(nums2)) - 1

		if nums1[leftIdx] > nums2[rightIdx] {
			removeK = rightIdx - right + 1
			right = rightIdx + 1
		} else {
			removeK = leftIdx - left + 1
			left = leftIdx + 1
		}
		k = k - removeK
	}

	// 边缘case
	if left >= len(nums1) {
		return nums2[right+k-1]
	} else if right >= len(nums2) {
		return nums1[left+k-1]
	} else {
		if nums1[left] < nums2[right] {
			return nums1[left]
		} else {
			return nums2[right]
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
