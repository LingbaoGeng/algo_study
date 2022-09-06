package tree

var memo [][]int

func NumTrees(n int) int {
	memo = make([][]int, n+1)
	for i, _ := range memo {
		memo[i] = make([]int, n+1)
	}
	return count(1, n)
}

func count(low, high int) int {
	if low > high {
		return 1
	}
	if memo[low][high] != 0 {
		return memo[low][high]
	}
	res := 0

	for i := low; i <= high; i++ {
		left := count(low, i-1)
		right := count(i+1, high)
		res += left * right
	}
	memo[low][high] = res
	return res
}