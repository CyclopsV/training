package SplitStrings

func Solution(str string) []string {
	res := []string{}
	lenStr := len([]rune(str))
	if lenStr%2 != 0 {
		str = str + "_"
	}
	for i := 0; i < lenStr; i = i + 2 {
		res = append(res, str[i:i+2])
	}
	return res
}
