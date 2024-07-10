package main

// Not a bad implementation. Scores around 2ms, beating 69 *nice* percent on average.
// Memory being 2.5mb, beating 67%
func longestCommonPrefix(strs []string) string {
	shortest := len(strs[0])
	prefix := ""

	for _, str := range strs {
		if len(str) < shortest {
			shortest = len(str)
		}
	}

  if shortest == 0 { return "" }

	for i := range shortest {
    curr := strs[0][i]
		for _, str := range strs {
			if str[i] != curr {
				return prefix
			}
		}
		prefix += string(curr)
	}

	return prefix
}
