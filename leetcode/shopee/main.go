package main

import (
	"fmt"
)

//func solution (test int) int {
//
//	return test
//}

//func main2() {
//	var userNum int
//	var transcNum int
//	var userMap map[string]int
//
//	fmt.Scanf("%d %d", &userNum, &transcNum)
//	userMap = make(map[string]int)
//	for i := 0; i < userNum; i++ {
//		var userName string
//		var deposit int
//		fmt.Scanf("%s %d\n", &userName, &deposit)
//		userMap[userName] = deposit
//	}
//	for i := 0; i < transcNum; i++ {
//		var userNameA string
//		var userNameB string
//		var money int
//		fmt.Scanf("%s %s %d\n", &userNameA, &userNameB, &money)
//		if deposit, ok := userMap[userNameA]; ok && deposit >= money {
//			userMap[userNameA] = deposit - money
//			if depositB, ok := userMap[userNameB]; ok {
//				userMap[userNameB] = depositB + money
//			} else {
//				userMap[userNameB] = money
//			}
//		}
//
//	}
//
//	keys := make([]string, 0, len(userMap))
//	for k := range userMap {
//		keys = append(keys, k)
//	}
//	sort.Strings(keys)
//
//	for _, k := range keys {
//		fmt.Printf(k + " " + strconv.Itoa(userMap[k]) + "\n")
//	}
//
//}


func main() {
	var testNum int
	fmt.Scanf("%d", &testNum)
	result := make([]string, testNum)
	//Start
	for i := 0; i < testNum; i++ {
		success := true
		var totalNum int
		fmt.Scanf("%d", &totalNum)
		//fmt.Println(totalNum)
		isOuterAndInnerCrossed := make([][2]bool, totalNum)
		numberStartAndEndPosition := make([][]int, totalNum)

		for j := 0; j < 2*totalNum; j++ {
			var cur int
			fmt.Scanf("%d", &cur)
			if numberStartAndEndPosition[cur-1] == nil {
				numberStartAndEndPosition[cur-1] = make([]int, 2)
				numberStartAndEndPosition[cur-1][0] = j
			} else {
				numberStartAndEndPosition[cur-1][1] = j
			}

		}

		for k := 0; k < totalNum; k++ {
			upcross := false
			downcross := false

			for m := 0; m < totalNum; m++ {
				if numberStartAndEndPosition[m][0] < numberStartAndEndPosition[k][0] && (numberStartAndEndPosition[m][1] < numberStartAndEndPosition[k][1] && numberStartAndEndPosition[m][1] > numberStartAndEndPosition[k][0]) || (numberStartAndEndPosition[m][0] > numberStartAndEndPosition[k][0] && numberStartAndEndPosition[m][0] < numberStartAndEndPosition[k][1]) && numberStartAndEndPosition[m][1] > numberStartAndEndPosition[k][1] {
					if isOuterAndInnerCrossed[m][0] == true {
						upcross = true
					}
					if isOuterAndInnerCrossed[m][1] == true {
						downcross = true
					}
				}
			}
			if upcross && downcross {
				success = false
				break
			} else if upcross {
				isOuterAndInnerCrossed[k][1] = true
			}else if downcross {
				isOuterAndInnerCrossed[k][0] = true
			} else {
				isOuterAndInnerCrossed[k][0] = true
			}

		}

		if success {
			result[i] = "yes"
		} else {
			result[i] = "no"
		}

	}

	for i := 0; i < testNum; i++ {
		fmt.Println(result[i])
	}

}
