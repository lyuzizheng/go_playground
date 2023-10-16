package contest_377

import (
	"fmt"
	"math"
	"strings"
)

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
    length := len(nums)
	if indexDifference > length {
		return []int{-1, -1}
	}
	for i := 0; i < length - indexDifference; i++ {
		for j := i + indexDifference; j < length; j++ {
			if math.Abs(float64(nums[i] - nums[j])) >= float64(valueDifference) {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}


func shortestBeautifulSubstring(s string, k int) string {

	var ones = make([]int,0, len(s))
	var chars = []rune(s)
	for i := 0; i < len(chars); i++ {
		if chars[i] == '1' {
			ones = append(ones, i)
		}
	} 
	if len(ones) < k {
		return ""
	}
	if len(ones) == k {
		return string(chars[ones[0]:ones[k+1]])
	}
	fmt.Println(ones)
	var min string
	for i, j := 0, k-1; j < len(ones); i, j = i+1, j+1 {
		temp := string(chars[ones[i]:ones[j]+1])
		fmt.Printf("candidate: %s \n", temp)
		if len(min) == 0 {
			min = temp
		}
		if len(temp) < len(min) {
			min = temp
		}
		if len(temp) > len(min) {
			continue
		}
		if output := strings.Compare(min, temp); output < 0 {
			min = temp
		}
	}
	return min
}


func constructProductMatrix(grid [][]int) [][]int {
	var product = 0
	var start = false
	var zeros = 0
	for i := 0; i < len(grid); i++ { 
		for j := 0; j < len(grid[i]); j++ {
			var temp = grid[i][j] % 12345
			if temp != 0 {
				if !start {
					start = true
					product = temp
				} else{
					product = product *temp % 12345
				}
			}
			if temp == 0 {
				zeros++
				if zeros > 1 {
					i = len(grid)
					break
				}
			}
		}
	}
	fmt.Println(product)
	
	
		for i := 0; i < len(grid); i++ { 
			for j := 0; j < len(grid[i]); j++ {
				var temp = grid[i][j] % 12345
				if zeros > 1 {
					grid[i][j] = 0
				} else if zeros == 1 {
					if temp != 0 {
						grid[i][j] = 0
					} else {
						grid[i][j] = product
					}
				} else if zeros == 0 {
					if temp == 0 {
						grid[i][j] = product
						continue
					}
					grid[i][j] = product / temp 
				}
				
			}
		}
		return grid
	}

	

	

		
	
	
    
