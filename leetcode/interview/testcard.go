package main

import (
	"fmt"
	"math/rand"
	"time"
)

func testCards(results []map[string][]int64) map[string]int64 {
	counter := make(map[string]int64)
	for k, _ := range results[0] {
		counter[k] = 0
	}
	fmt.Println("len results", len(results))
	totalTrue := 0
	totalFalse := 0
	for _, result := range results {
		occur := false
		for k, v := range result { //执行了 100000 次
			for _, card := range v {
				if card == 54 {
					counter[k] += 1
					if occur {
						fmt.Println("abnormal", result)
					}
					occur = true
					totalTrue++
				}
			}
		}
		if !occur {
			totalFalse++
		}
	}
	fmt.Println("totalTrue", totalTrue, "totalFalse", totalFalse)
	return counter
}

// 扑克牌问
// A拿到的概率是两倍

func DeliverCards(cards []int64, person []string) (result map[string][]int64) {

	// new rand
	newRand := rand.New(rand.NewSource(time.Now().UnixMicro()))

	// Assume person = ["A", "B", "C", "D"]

	// cards = [1~54] , 53 小王， 54 大王

	currentCard := cards
	stillHaveJoker := true

	currentPersonIndex := 0

	result = make(map[string][]int64, len(person)*3)

	for len(currentCard) > 0 {
		deliveredIndex := int64(0)
		// A 的概率两倍
		if stillHaveJoker && currentPersonIndex == 0 {
			currentCard = append(currentCard, 54)
			deliveredIndex = newRand.Int63n(int64(len(currentCard)))
			result[person[currentPersonIndex]] = append(result[person[currentPersonIndex]], currentCard[deliveredIndex])
			//fmt.Println("A", "Current Card: ", currentCard, "Delivered Index: ", deliveredIndex, "Result Person: ", person[currentPersonIndex], "person: ", result[person[currentPersonIndex]])
			// 如果抽到的是 joker
			if currentCard[deliveredIndex] == 54 {
				stillHaveJoker = false
			}
			if deliveredIndex == int64(len(currentCard)-1) {
				//fmt.Println("got joker")
				deliveredIndex--
			}
			currentCard = currentCard[:len(currentCard)-1]
		} else {
			deliveredIndex = rand.Int63n(int64(len(currentCard)))
			if currentCard[deliveredIndex] == 54 {
				stillHaveJoker = false
			}
			result[person[currentPersonIndex]] = append(result[person[currentPersonIndex]], currentCard[deliveredIndex])
			//fmt.Println("Current Card: ", currentCard, "Delivered Index: ", deliveredIndex, "Result Person: ", person[currentPersonIndex], "person: ", result[person[currentPersonIndex]])
		}

		// remove the delivered card
		if deliveredIndex == int64(len(currentCard)-1) {
			//fmt.Println("deliveredIndex", deliveredIndex, "len(currentCard)", len(currentCard))
			currentCard = currentCard[:len(currentCard)-1]
		} else {
			//fmt.Println("else","deliveredIndex", deliveredIndex, "len(currentCard)", len(currentCard))
			currentCard = append(currentCard[:deliveredIndex], currentCard[deliveredIndex+1:]...)
		}

		currentPersonIndex = (currentPersonIndex + 1) % len(person)
	}
	return result
}
