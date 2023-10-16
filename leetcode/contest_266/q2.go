package contest_266

// Given a string word, return the sum of the number of vowels ('a', 'e', 'i', 'o', and 'u') in every substring of word.
func CountVowels(word string) int64 {
	vowelsMap := map[string]int{
		"a": 0b1,
		"e": 0b10,
		"i": 0b100,
		"o": 0b1000,
		"u": 0b10000,
	}
	var count int64
	strLen := len(word)
	for i := 0; i < strLen; i++ {
		//if word[i] is vowel, add it to count
		if _, ok := vowelsMap[string(word[i])]; ok {
			a := int64(strLen)
			index := i + 1
			if index > strLen>>1 {
				index = strLen - index + 1
			}
			for j := 1; j <= index; j++ {
				if a >= 1 {
					count += a
					a -= 2
				}
			}
		}
	}
	return count

}
