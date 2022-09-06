package other

func AdvantageCount(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	nums2Copy := make([]int, 0)
	nums2Copy = append(nums2Copy, nums2...)
	nums2Map := make(map[int]int)

	for i := 0; i < len(nums2); i++ {
		nums2Map[nums2[i]] = i
	}

	sort_asc(nums1)
	sort_desc(nums2Copy)

	left := 0
	right := len(nums1) - 1

	for i := 0; i < len(nums2Copy); i++ {
		maxVal := nums2Copy[i]
		index := nums2Map[maxVal]

		if maxVal < nums1[right] {
			ans[index] = nums1[right]
			right--
		} else {
			ans[index] = nums1[left]
			left++
		}
	}
	return ans
}

func sort_asc(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				temp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
			}
		}
	}
}

func sort_desc(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[j-1] {
				temp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
			}
		}
	}
}
