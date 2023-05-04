package StringEndsWith

import "strings"

func Solution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
