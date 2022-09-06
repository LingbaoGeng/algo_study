package sort_algo

func MergeSort(nums []int, low, high int) {
	if low == high {
		return
	}

	mid := (low + high) / 2

	MergeSort(nums, low, mid)
	MergeSort(nums, mid+1, high)

	mergeTwoArray(nums, low, mid, high)
}

func mergeTwoArray(nums []int, low, mid, high int) {
	i, loc := low, low
	j := mid + 1

	temp := make([]int, len(nums))

	for i <= mid && j <= high {
		if nums[i] > nums[j] {
			temp[loc] = nums[j]
			j++
		} else {
			temp[loc] = nums[i]
			i++
		}
		loc++
	}
	for i <= mid {
		temp[loc] = nums[i]
		i++
		loc++
	}
	for j <= high {
		temp[loc] = nums[j]
		j++
		loc++
	}

	for k := low; k <= high; k++ {
		nums[k] = temp[k]
	}
}
