package climbingstairs

/*
Problem:
	Paid Staircase

	You are climbing a paid staircase. It takes n steps to reach to the top and you have to
	pay p[i] to step on the i-th stair. Each time you can climb 1 or 2 steps.
	return the cheapest path to the top of the staircase.
*/

// Time complexity: O(n)
// Space complexity: O(n)
func paidStaircaseWithPath(n int, p []int) []int {
	Dp := make([]int, n+1)
	from := make([]int, n+1)
	Dp[0] = 0
	Dp[1] = p[1]
	for i := 2; i <= n; i++ {
		Dp[i] = min(Dp[i-1], Dp[i-2]) + p[i]
		if Dp[i-1] < Dp[i-2] {
			from[i] = i - 1
		} else {
			from[i] = i - 2
		}
	}
	path := []int{}
	for curr := n; curr >= 0; curr = from[curr] {
		path = append(path, curr)
		if curr == 0 {
			break
		}
	}
	return reverse(path)
}

func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}
