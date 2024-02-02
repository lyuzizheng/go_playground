package contest373

func areSimilar(mat [][]int, k int) bool {
	if len(mat) == 0 {
		return false
	}
	length := len(mat[0])
	if length == 0 {
		return false
	}
	evenOffset := k % length
	oddOffset := length - evenOffset
	if evenOffset == 0 {
		return true
	}
	// if k = 3, length = 5, offset = 2
	for i := 0; i < len(mat); i++ {
		for j := 0; j < length; j++ {
			// odd row shift right
			if i%2 == 1 {
				if mat[i][(j+oddOffset)%length] != mat[i][j] {
					return false
				}
			} else { // even row shift left
				if mat[i][(j+evenOffset)%length] != mat[i][j] {
					return false
				}
			}

		}
	}
	return true
}

func beautifulSubstrings(s string, k int) int {
	vowels := make(map[rune]bool)
	vowels['a'] = true
	vowels['e'] = true
	vowels['i'] = true
	vowels['o'] = true
	vowels['u'] = true
	vowelsPositions := make([]int, 0)
	for i, c := range s {
		if vowels[c] {
			vowelsPositions = append(vowelsPositions, i)
		}
	}
	for i := 0; i < len(vowelsPositions); i++ {
		for j := i + 1; j < len(vowelsPositions); j++ {
			numOfVowels := j - i + 1
			if numOfVowels*numOfVowels%k != 0 {
				continue
			}
			numOfConsonantsInBetween := vowelsPositions[j] - vowelsPositions[i] - numOfVowels + 1
			if numOfConsonantsInBetween > numOfVowels {
				continue
			}
			numOfConsonantsRequired := numOfVowels - numOfConsonantsInBetween
			// check if there are enough consonants
			var spaceBehind, spaceInfront int
			if j == len(vowelsPositions)-1 {
				spaceBehind = len(s) - vowelsPositions[j] - 1
				if i == 0 {
					spaceInfront = vowelsPositions[i]
				} else {
					spaceInfront = vowelsPositions[i] - vowelsPositions[i-1] - 1
				}
			} else {
				spaceBehind = vowelsPositions[j+1] - vowelsPositions[j] - 1
				if i == 0 {
					spaceInfront = vowelsPositions[i]
				} else {
					spaceInfront = vowelsPositions[i] - vowelsPositions[i-1] - 1
				}
			}
			if spaceBehind+spaceInfront < numOfConsonantsRequired {
				continue
			}

		}
	}

	return 0

}
