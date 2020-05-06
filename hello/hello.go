package hello

import "strings"

// longestCommonPrefix LeetCode最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _,k := range strs {
		for strings.Index(k,prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix) - 1]
		}
	}
	return prefix
}


