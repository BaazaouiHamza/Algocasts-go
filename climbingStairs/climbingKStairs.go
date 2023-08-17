package climbingstairs

/*

Problem:
	Climbing Stairs(K stepsn,space optimized)

	You are climbing a stair case.It takes n steps to reach to the top.
	Each time you can  climb 1 .. k steps.
	In how many distinct ways can you climb to the top?
*/

// Time Complexity: O(nk)
// Space Complexity: O(n)
func climbKStairs(n int, k int) int {
	dp := make([]int, k)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j < k; j++ {
			if i < j {
				continue
			}
			dp[i%k] += dp[(i-j)%k]
		}
	}

	return dp[n%k]
}
