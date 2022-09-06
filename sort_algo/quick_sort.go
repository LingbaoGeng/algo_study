package sort_algo

func QuickSort(nums []int) []int {
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, low, high int) {
	if low >= high {
		return
	}
	p := partition(nums, low, high)
	sort(nums, low, p-1)
	sort(nums, p+1, high)
}

func partition(nums []int, low, high int) int {
	pivot := nums[low]
	i := low + 1
	j := high
	for i <= j {
		for i < high && nums[i] <= pivot {
			i++
		}

		for j > low && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		swap(nums, i, j)
	}
	swap(nums, low, j)
	return j
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
