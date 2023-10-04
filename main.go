package main

import (
	"fmt"
	"math"
)

// [-2,1,-3,4,-1,2,1,-5,4]
// Prefix Sum:  -2, -1, -4, 0, -1, 1, 2, -3, 1

func FindSum(numArr []int) int {
	if numArr == nil {
		return 0
	}

	var prefixSumArr = make([]int, len(numArr))
	var result int

	singleMax := math.MinInt
	// find prefix sum
	curMin := math.MaxInt
	curMax := math.MinInt
	curSum := 0
	for i := 0; i < len(numArr); i++ {
		if singleMax < numArr[i] {
			singleMax = numArr[i]
		}
		prefixSumArr[i] = curSum + numArr[i]
		if prefixSumArr[i] < curMin { // find a new min
			curMin = prefixSumArr[i]
			fmt.Println("cur Min", curMin)
			curMax = math.MinInt // reset max
		}
		if prefixSumArr[i] > curMax { // find a new max after the new min
			curMax = prefixSumArr[i]
			fmt.Println("cur max", curMax)

		}
		curSum = prefixSumArr[i]

		// fins the biggest increase since
		if curMax > curMin {
			if (curMax - curMin) > result {
				result = curMax - curMin
			}
		}

	}

	// all negative just return a single max of the array
	if result == 0 {
		return singleMax
	}

	fmt.Println()

	return result
}

func main() {

	//fmt.Print(FindSum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	//fmt.Print(FindSum([]int{0, 0, 1, 1, 2, -100, 20, -10, 20}))
	//fmt.Print(FindSum([]int{-5, -1, -3, -4, -10, -1, -90}))

	fmt.Println(findPermuteAInB("abcdea", "okojieafabcdeaf"))
	fmt.Println(findPermuteAInB("abcdea", "okojieafbcdeaf"))

	fmt.Println(findPermuteAInB("abcdea", "okojieafbcadeaf"))
}

// A: "abcde"
// B: "okojieaf bcdea f"
func findPermuteAInB(A string, B string) bool {
	needMatchNum := len(A)

	AMap := map[rune]int{}

	// build a occurance map
	for _, myChar := range A {
		myRune := rune(myChar)
		val := AMap[myRune]
		AMap[myRune] = val + 1 // increase the occurance
	}

	// find match
	curMatchNumber := 0
	bMap := NewMap(AMap)
	for _, myChar := range B {

		if val := bMap[myChar]; val != 0 { // match
			bMap[myChar] = val - 1
			curMatchNumber++ // increase match count
		} else { // not match
			if curMatchNumber != 0 { // optimise
				bMap = NewMap(AMap)
			}
			curMatchNumber = 0 // reset current match
		}
		if curMatchNumber == needMatchNum {
			return true
		}
	}

	return false
}

func NewMap(src map[rune]int) map[rune]int {
	BMap := map[rune]int{}
	for k, v := range src {
		BMap[k] = v
	}
	return BMap
}
