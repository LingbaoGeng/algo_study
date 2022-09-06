package bfs

func OpenLock(deadends []string, target string) int {
	visited := make(map[string]bool)
	queue := make([]string, 0)
	deadendsMap := make(map[string]bool, len(deadends))
	for _, value := range deadends {
		deadendsMap[value] = true
	}
	step := 0
	originStr := "0000"
	queue = append(queue, originStr)
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			if _, ok := deadendsMap[queue[i]]; ok {
				continue
			}
			if queue[i] == target {
				return step
			}
			for j := 0; j < 4; j++ {
				up := plusOne(queue[i], j)
				if !visited[up] {
					queue = append(queue, up)
					visited[up] = true
				}
				down := minusOne(queue[i], j)
				if !visited[down] {
					queue = append(queue, down)
					visited[down] = true
				}
			}
		}
		step++
		queue = queue[size:]
	}
	return -1
}

func plusOne(str string, index int) string {
	byteSlice := []byte(str)
	if byteSlice[index] == '9' {
		byteSlice[index] = '0'
	} else {
		byteSlice[index] += 1
	}
	return string(byteSlice)
}

func minusOne(str string, index int) string {
	byteSlice := []byte(str)
	if byteSlice[index] == '0' {
		byteSlice[index] = '9'
	} else {
		byteSlice[index] -= 1
	}
	return string(byteSlice)
}
