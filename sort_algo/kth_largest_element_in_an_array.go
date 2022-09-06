package sort_algo

func FindKthLargest(nums []int, k int) int {
	low := 0
	high := len(nums) - 1
	k = len(nums) - k

	for low <= high {
		p := partition1(nums, low, high)
		if p < k {
			low = p + 1
		} else if p > k {
			high = p - 1
		} else {
			return nums[p]
		}
	}
	return -1
}

func partition1(nums []int, low, high int) int {
	piviot := nums[low]
	i := low + 1
	j := high

	for i <= j {
		for i < high && nums[i] < piviot {
			i++
		}

		for j > low && nums[j] > piviot {
			j--
		}

		if i >= j {
			break
		}
		swap1(nums, i, j)
	}
	swap1(nums, low, j)
	return j
}

func swap1(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
