package lcs

// LCS 最长公共子串
func LCS(src string, dst string) string {
	if len(src) == 0 || len(dst) == 0 {
		return ""
	}
	maxLen := 0
	end := 0
	dp := make([]int, len(dst)+1)
	for i := 1; i < len(src)+1; i++ {
		leftUp := 0
		for j := 1; j < len(dst)+1; j++ {
			up := dp[j]
			if src[i-1] == dst[j-1] {
				dp[j] = leftUp + 1
				if dp[j] > maxLen {
					maxLen = dp[j]
					end = i
				}
			} else {
				dp[j] = 0
			}
			leftUp = up
		}
	}
	return src[end-maxLen : end]
}

func simple(src string, dst string) int {
	m := len(src)
	n := len(dst)

	memo := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		memo[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if src[i-1] == dst[j-1] {
				memo[i][j] = memo[i-1][j-1] + 1
			}
		}
	}

	longest := 0
	for i := range memo {
		for j, e2 := range memo[i] {
			if longest < memo[i][j] {
				longest = e2
			}
		}
	}

	return longest
}
