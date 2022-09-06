package main

import (
	"algo_study/algo"
	"fmt"
)

func main() {

	/*
		tree := &object.TreeNode{
			Val: 1,
			Left: &object.TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &object.TreeNode{
				Val: 5,
				Left: &object.TreeNode{
					Val:   10,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		}
		ans := bfs.MinDepth(tree)
		fmt.Println(ans)
	*/
	/*
		coins := []int{1, 2, 5}
		ans := dp.CoinChange(coins, 11)
		fmt.Println(ans)

	*/

	/*
		nums := []int{5, 4, 6, 2}
		ans := back_trace.FullPermutaiton(nums)
		fmt.Println(ans)

	*/
	//N-皇后问题
	/*
		ans := back_trace.SolveNQueens(4)
		fmt.Println(ans)

	*/
	/*
		head := &object.ListNode{
			Val: 1,
			Next: &object.ListNode{
				Val: 2,
				Next: &object.ListNode{
					Val: 3,
					Next: &object.ListNode{
						Val: 4,
						Next: &object.ListNode{
							Val: 5,
							Next: &object.ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		}
		ans := link.ReverseKGroup(head, 2)
		for ans != nil {
			fmt.Println(ans.Val)
			ans = ans.Next
		}

	*/

	//ans := sliding_window.LengthOfLongestSubstring("pwwkew")
	nums := []int{1, 1, 3, 5, 3, 5, 1, 100, 0, -1}
	ans := algo.Find2thItem(nums)
	fmt.Println(ans)
	nums = []int{1, 1, 1, 1}
	ans = algo.Find2thItem(nums)
	fmt.Println(ans)

	nums = []int{1, 3, 5, 2, 10, 9, 1, 11, 4}
	ans1 := algo.GetPrice([]int{1, 2, 3})
	fmt.Println(ans1)

	sum := 1 + 2 + 1.1
	fmt.Println(sum)
}
