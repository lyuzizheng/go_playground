package daily_contest

func generateParenthesis(n int) []string {

	cur := make([]byte, 0, n*2)
	sum := 0
	res := make([]string, 0)

	return generate(res,cur,sum,n)

}


func generate(result []string, current []byte, sum int, n int) []string {
	if sum < 0 {
		return result
	} else if sum > 2 * n - len(current){
		return result
	} else if sum == 0 {
		if len(current) == 2*n {
			result = append(result, string(current))
			return result
		}
	}

	if len(current) > 2*n {
		return result
	}

	result = generate(result, append(current, '('), sum + 1, n)
	result = generate(result, append(current, ')'), sum - 1, n)
	return result
}
