package contest_372

import (
	"strconv"
)

/*
You are given three strings s1, s2, and s3. You have to perform the following operation on these three strings as many times as you want.

In one operation you can choose one of these three strings such that its length is at least 2 and delete the rightmost character of it.

Return the minimum number of operations you need to perform to make the three strings equal if there is a way to make them equal, otherwise, return -1.
*/

func findMinimumOperations(s1 string, s2 string, s3 string) int {
	// find the longest common prefix
	var i int
	for i = 0; i < len(s1) && i < len(s2) && i < len(s3); i++ {
		if s1[i] != s2[i] || s2[i] != s3[i] {
			break
		}
	}
	if i == 0 {
		return -1
	}
	var min = len(s1) - i + len(s2) - i + len(s3) - i
	return min
}

/*
There are n balls on a table, each ball has a color black or white.
You are given a 0-indexed binary string s of length n, where 1 and 0 represent black and white balls, respectively.
In each step, you can choose two adjacent balls and swap them.
Return the minimum number of steps to group all the black balls to the right and all the white balls to the left.

Input: s = "100"
Output: 2
Explanation: We can group all the black balls to the right in the following way:
- Swap s[0] and s[1], s = "010".
- Swap s[1] and s[2], s = "001".
It can be proven that the minimum number of steps needed is 2.
*/
func minimumSteps(s string) int64 {
	sum := 0
	start := 0
	end := len(s) - 1
	sB := []byte(s)

	var nextWhite int
	for start <= end {
		if sB[start] == '0' {
			start++
			continue
		}
		if nextWhite <= start {
			nextWhite = start + 1
		}
		for nextWhite <= end {
			if sB[nextWhite] == '0' {
				break
			}
			nextWhite++
		}
		if nextWhite > end {
			break
		}
		sB[start], sB[nextWhite] = sB[nextWhite], sB[start]
		sum += nextWhite - start
	}

	return int64(sum)
}

/*
Given three integers a, b, and n, return the maximum value of (a XOR x) * (b XOR x) where 0 <= x < 2n.
Since the answer may be too large, return it modulo 109 + 7.
Note that XOR is the bitwise XOR operation.
*/
func maximumXorProduct(a int64, b int64, n int) int {
	var min int64
	if a < b {
		min = a
	} else {
		min = b
	}
	binaryStr := strconv.FormatInt(min, 2)
	if len(binaryStr) < n {
		return 0
	}

	return 0

}
