package kata

func LongestConsec(strarr []string, k int) string {
	// your code
	if len(strarr) == 0 || k > len(strarr) || k <= 0 {
		return ""
	}
	var currentStr, str string

	for i := 0; i < len(strarr)-k+1; i++ {
		for j := i; j < i+k; j++ {
			currentStr += strarr[j]
		}
		if len(currentStr) > len(str) {
			str = currentStr
		}
		currentStr = ""
	}
	return str
}
