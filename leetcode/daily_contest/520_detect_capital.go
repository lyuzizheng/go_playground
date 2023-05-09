package daily_contest

const (
	CapStart = 0x41
	CapEnd          = 0x5a
	NonCapitalStart = 0x61
	NonCapitalEnd   = 0x7a
)

func DetectCapitalUse(word string) bool {

	if len(word) <= 1 {
		return true
	}


	if word[0] >= CapStart && word[0] <= CapEnd {
		result := true
		//Detect first char is capital
		for i := 1; i < len(word); i++ {
			if word[i] < NonCapitalStart || word[i] > NonCapitalEnd {
				result = false
				break
			}
		}
		if result {
			return result
		}

		//Detect if all chars are capital
		result = true
		for i := 1; i < len(word); i++ {
			if word[i] < CapStart || word[i] > CapEnd {
				result = false
				break
			}
		}
		return result

	} else {
		result := true
		for i := 1; i < len(word); i++ {
			if word[i] < NonCapitalStart || word[i] > NonCapitalEnd {
				result = false
				break
			}
		}
		return result
	}


}