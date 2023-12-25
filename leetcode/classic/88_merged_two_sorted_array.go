package classic

// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
// Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
// The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

type Node struct {
	Element int
	Next    *Node
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// merge two array into one
	//for i := 0; i < n; i++ {
	//	nums1[m-1+i] = nums2[i]
	//}
	// use the heapify whose complexity is O(m+n)
	//var slices IntSlice = nums1
	//heap.Init(&slices)
	// heapify is a process that makes an array as a maxheap or minheap, but array is not absolutely sorted

	// linked_list can reduce the complexity of insert into O(1)
	var (
		node1, node2 = &Node{Element: -10000000000, Next: nil}, &Node{Element: -10000000000, Next: nil}
		cur          *Node
	)

	cur = node1
	for i := 0; i < m; i++ {
		temp := &Node{
			Element: nums1[i],
			Next:    nil,
		}
		cur.Next = temp
		cur = cur.Next
	}

	cur = node2
	for i := 0; i < n; i++ {
		temp := &Node{
			Element: nums2[i],
			Next:    nil,
		}
		cur.Next = temp
		cur = cur.Next
	}

	prev1 := node1
	ptr1, ptr2 := node1.Next, node2.Next
	for ptr1 != nil && ptr2 != nil {
		if ptr1.Element <= ptr2.Element {
			prev1 = prev1.Next
			ptr1 = ptr1.Next
		} else {
			prev1.Next = ptr2

			temp2 := ptr2.Next
			ptr2.Next = ptr1
			ptr2 = temp2
			prev1 = prev1.Next
		}
	}

	if ptr2 != nil {
		prev1.Next = ptr2
	}

	cur = node1.Next
	for i := range nums1 {
		nums1[i] = cur.Element
		cur = cur.Next
	}

}

type IntSlice []int

func (r IntSlice) Len() int {
	return len(r)
}

func (r IntSlice) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r IntSlice) Swap(i, j int) {
	temp := r[i]
	r[i] = r[j]
	r[j] = temp
}

func (r *IntSlice) Push(x any) {
	*r = append(*r, x.(int))
}

func (r *IntSlice) Pop() any {
	old := *r
	n := len(old)
	x := old[n-1]
	*r = old[0 : n-1]
	return x
}

//// has some problems

func mergeByIndex(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, 0, m+n)
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	if i < m {
		result = append(result, nums1[i:]...)
	}

	if j < n {
		result = append(result, nums2[j:]...)
	}
	nums1 = result
}
