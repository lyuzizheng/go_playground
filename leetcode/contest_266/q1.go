package contest_266


//A substring is a contiguous (non-empty) sequence of characters within a string.
//
//A vowel substring is a substring that only consists of vowels ('a', 'e', 'i', 'o', and 'u') and has all five vowels present in it.
//
//Given a string word, return the number of vowel substrings in word.
func CountVowelSubstrings(word string) int {

	startposition := -1
	currentPosition := 0
	count := 0

	vowelsMap := map[string]int{
		"a": 0b1,
        "e": 0b10,
        "i": 0b100,
        "o": 0b1000,
        "u": 0b10000,
	}

	if len(word) <= 4 {
        return 0
    }

	fullfillConsition := 0b00000

	//iterate character in word
	for index, char := range word {
        currentPosition = index
		//check if current character is vowel
		if val, ok := vowelsMap[string(char)]; ok {
			if startposition == -1 {
				startposition = index
			}
			fullfillConsition = fullfillConsition | val
		} else {
			if fullfillConsition == 0b11111 {
				count += countOnlyVowelSubstring(word[startposition:currentPosition], vowelsMap)
				fullfillConsition = 0b00000
				startposition = -1
			} else {
				if fullfillConsition != 0b00000 {
                    fullfillConsition = 0b00000
                    startposition = -1
                }
			}
		}
    }
	if fullfillConsition == 0b11111 {
		count += countOnlyVowelSubstring(word[startposition:len(word)], vowelsMap)
	}
	return count

}

func countOnlyVowelSubstring(word string, vowelMap map[string]int) int {
	count := 0

	startposition := 0
	currentPosition := 0

	cowelsCount := map[string]int{
		"a": 0,
		"e": 0,
		"i": 0,
		"o": 0,
		"u": 0,
	}
	for i := 0; i < len(word); i++ {
		if val, ok := cowelsCount[string(word[i])]; ok {
			cowelsCount[string(word[i])] = val + 1
		}
		//if all values of cowelsCount are greater than 1
		for true {
			if cowelsCount["a"] >= 1 && cowelsCount["e"] >= 1 && cowelsCount["i"] >= 1 && cowelsCount["o"] >= 1 && cowelsCount["u"] >= 1 {
				currentPosition = i
				count += 1
				count += len(word) - 1 - currentPosition

				first := string(word[startposition])
				cowelsCount[first] = cowelsCount[first] - 1
				startposition ++

			} else {
				break
			}
		}


	}


	return count

}