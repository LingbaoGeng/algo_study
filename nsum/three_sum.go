package nsum

import "fmt"

var ans [][]int

func ThreeSum(nums []int) [][]int {
	ans = make([][]int, 0)
	sort(nums)
	fmt.Println(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		tuples := twoSum(nums, i+1, -nums[i])
		for _, val := range tuples {
			temp := make([]int, 0)
			temp = append(temp, nums[i])
			temp = append(temp, val...)
			ans = append(ans, temp)
		}
		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return ans
}

func sort(nums []int) {
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

func twoSum(nums []int, start, target int) [][]int {
	twoSumAns := make([][]int, 0)
	low := start
	high := len(nums) - 1
	for low < high {
		sum := nums[low] + nums[high]
		left := nums[low]
		right := nums[high]
		if sum == target {
			temp := make([]int, 0)
			temp = append(temp, left, right)
			//temp = append(temp,right)
			twoSumAns = append(twoSumAns, temp)
			for low < high && nums[low] == left {
				low++
			}
			for low < high && nums[high] == right {
				high--
			}
		} else if sum < target {
			for low < high && nums[low] == left {
				low++
			}
		} else if sum > target {
			for low < high && nums[high] == right {
				high--
			}
		}
	}
	return twoSumAns
}
