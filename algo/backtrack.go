package algo

import (
	"errors"
	"slices"
	"sort"
)

// Bondee
// [1,2,3,4,6] 10
// [1,2,3,4],[1,3,6],[4,6]
var result [][]int
var currentTrack []int
var trackSum int

func findSum(nums []int, target int) {

	// sort the nums
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})


	backtrack(nums, 0, 10)

}

func backtrack(nums []int, start int, target int) {

	if trackSum == target {
		result = append(result, append(make([]int, 0), currentTrack...))
		return
	}

	if trackSum > target {
		return
	}

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		currentTrack = append(currentTrack, nums[i])
		trackSum += nums[i]
		backtrack(nums, i+1, target)
		currentTrack = currentTrack[:len(currentTrack)-1]
		trackSum -= nums[i]
	}
}

const TransferLimit = 100

type Account struct {
	Name   string
	Amount int64
}

// interface => ["b","a" 20]

// A = 20
// B = 150
// C = 150
// B -> A => 50
// C - >A = 30
func MakeTransfer(accounts []*Account, limit int64) ([][]interface{}, error) {
	if len(accounts) == 0 {
		return nil, errors.New("ddd")
	}

	belowLimit := []*Account{}
	aboveLimit := []*Account{}

	var moneyRequired int64 = 0
	var moneySpared int64 = 0

	for _, account := range accounts {
		if account.Amount < limit {
			belowLimit = append(belowLimit, account)
			moneyRequired += limit - account.Amount
		}
		if account.Amount > limit {
			aboveLimit = append(aboveLimit, account)
			moneySpared += account.Amount - limit
		}
	}

	if moneySpared < moneyRequired {
		return nil, nil
	}

	result := [][]interface{}{}
	spareAccountIndex := 0 //todo

	for _, account := range belowLimit {

		// loop finding the required amount
		for spareAccountIndex < len(aboveLimit) {

			requireAmount := limit - account.Amount // 20

			spareMoney := aboveLimit[spareAccountIndex].Amount - limit

			if spareMoney >= requireAmount {
				result = append(result, []interface{}{aboveLimit[spareAccountIndex].Name, account.Name, requireAmount})
				aboveLimit[spareAccountIndex].Amount -= requireAmount
				account.Amount += requireAmount
				break
			}

			if spareMoney < requireAmount {
				result = append(result, []interface{}{aboveLimit[spareAccountIndex].Name, account.Name, spareMoney})
				aboveLimit[spareAccountIndex].Amount -= spareMoney
				account.Amount += spareMoney // 70
				spareAccountIndex++
			}
		}
	}

	return result, nil

}

/*
input : [0,3,1,6,2,2,7]
output: 4, [3, 6. ]
*/
func FindLongestSequence(sequence []int64) int64 {
	// return 0 if nil or empty
	if len(sequence) == 0 {
		return 0
	}
	// use a dp slice to memorise the history
	currentMax := make([]int, len(sequence))
	currentMax[0] = 1
	maxLength := 1

	// traverse the sequence
	for i := 1; i < len(sequence); i++ {
		currentMax[i] = 1                             // assign current = 1
		for previous := 0; previous < i; previous++ { // find previous max
			if sequence[i] > sequence[previous] { // if current is bigger than previous one, it means we can add one to the i
				if currentMax[i] < currentMax[previous]+1 {
					currentMax[i] = currentMax[previous] + 1
				} // update i max
			}
		}
		if maxLength <= currentMax[i] {
			maxLength = currentMax[i]
		}
	}

	return int64(maxLength)
}
