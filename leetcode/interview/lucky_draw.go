package main

import (
	"math/rand"
	"sort"
	"time"
)

var prizeMap = map[string]float64{
	"gold":    0.2,
	"knift":   0.2,
	"diamond": 0.1,
	"water":   0.1,
	"nothing": 0.4,
}

var prizeWeightage = map[int][]string{
	1:  []string{"diamand", "iPhone", "watch"},
	2:  []string{"gold"},
	10: []string{"nothing"},
}

func luckyDraw(prizeMap map[string]float64) string {
	random := rand.Float64()
	for k, v := range prizeMap {
		if random < v {
			return k
		}
		random -= v
	}
	return "nothing"
}

func makePrizeRange(prizeMap map[int][]string) (prizeRange []int, actualPrizes []string) {
	current := 0
	prizeRange = make([]int, 0)
	actualPrizes = make([]string, 0)
	for weight, prizes := range prizeMap {
		for _, prize := range prizes {
			current += weight
			prizeRange = append(prizeRange, current)
			actualPrizes = append(actualPrizes, prize)
		}
	}
	return
}

func getLuckyDrawFunc(prizeMap map[int][]string) func() string {
	//Prepare a linear arrangement of all probabilities
	current := 0
	prizeRange := make([]int, 0)
	actualPrizes := make([]string, 0)
	for weight, prizes := range prizeMap {
		for _, prize := range prizes {
			current += weight
			prizeRange = append(prizeRange, current)
			actualPrizes = append(actualPrizes, prize)
		}
	}
	rand.Seed(time.Now().UnixNano())
	//return the func that used to do lucky draw
	return func() string {
		random := rand.Intn(current)
		//Binary Search the prizeRange
		low := 0
		high := len(prizeRange) - 1
		for high >= low {
			mid := (low + high) / 2
			if random < prizeRange[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return actualPrizes[low]
	}
}

// Interface has the format of prize name and weight
type PrizeCollection [][]interface{}

var a = PrizeCollection{
	{"Dyson", 1},
	{"iPhone", 2},
	{"XBox", 3},
	{"Gold", 5},
	{"Knift", 10},
	{"Water", 7},
	{"Diamond", 12},
}

func getHopScotchFunc(collection PrizeCollection) func() string {
	//Reverse Sort the Weights
	sort.Slice(collection, func(i, j int) bool {
		return collection[j][1].(int) < collection[i][1].(int)
	})
	sum := 0
	collectionWeightSum := make([]int, len(collection))
	for i := 0; i < len(collection); i++ {
		sum += collection[i][1].(int)
		collectionWeightSum[i] = sum
	}

	rand.Seed(time.Now().UnixNano())
	//return the func that used to do lucky draw
	return func() string {
		target := rand.Intn(collectionWeightSum[len(collectionWeightSum)-1])
		guessIndex := 0
		for true {
			if collectionWeightSum[guessIndex] > target {
				break
			}
			currentWeight := collection[guessIndex][1].(int)
			hopDistance := target - collectionWeightSum[guessIndex]
			hopIndex := 1 + hopDistance/currentWeight
			guessIndex += hopIndex
		}
		return collection[guessIndex][0].(string)
	}
}

func getAliasFunc(collection PrizeCollection) func() string {
	totalPrizeNum := len(collection)
	sum := 0
	for i := 0; i < len(collection); i++ {
		sum += collection[i][1].(int)
	}
	average := float64(sum) / float64(totalPrizeNum)
	aliases := make([][]interface{}, totalPrizeNum)
	for i := 0; i < totalPrizeNum; i++ {
		aliases[i] = []interface{}{1.0, 0.0}
	}
	bigWeights := make([][]interface{}, 0)
	smallWeights := make([][]interface{}, 0)
	for index, prizeItem := range collection {
		if float64(prizeItem[1].(int)) < average {
			smallWeights = append(smallWeights, []interface{}{index, float64(prizeItem[1].(int)) / average})
		} else {
			bigWeights = append(bigWeights, []interface{}{index, float64(prizeItem[1].(int)) / average})
		}
	}
	bigWeightsPosition := 0
	for i := 0; i < len(smallWeights); i++ {
		aliases[smallWeights[i][0].(int)] = []interface{}{smallWeights[i][1].(float64), bigWeights[bigWeightsPosition][0].(int)}
		bigWeights[bigWeightsPosition][1] = bigWeights[bigWeightsPosition][1].(float64) - (1 - smallWeights[i][1].(float64))
		if bigWeights[bigWeightsPosition][1].(float64) <= 1 {
			smallWeights = append(smallWeights, bigWeights[bigWeightsPosition])
			bigWeightsPosition++
			if bigWeightsPosition >= len(bigWeights) {
				break
			}
		}
	}
	rand.Seed(time.Now().UnixNano())
	//return the func that used to do lucky draw
	return func() string {
		target := rand.Float64() * float64(len(collection))
		targetAlias := int(target)
		targetWeight := target - float64(targetAlias)
		if targetWeight < aliases[targetAlias][0].(float64) {
			return collection[targetAlias][0].(string)
		} else {
			return collection[aliases[targetAlias][1].(int)][0].(string)
		}
	}
}
