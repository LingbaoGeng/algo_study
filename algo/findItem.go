package algo

func Find2thItem(nums []int) []int {
	result := make([]int, 0)
	countMap := make(map[int]int)
	for _, val := range nums {
		countMap[val]++
	}
	maxCount := 0
	secondCount := 0
	for _, value := range countMap {
		if value > maxCount {
			secondCount = maxCount
			maxCount = value
		} else if value > secondCount {
			secondCount = value
		}
	}

	for index, value := range countMap {
		if value == secondCount {
			result = append(result, index)
		}
	}
	return result
}
