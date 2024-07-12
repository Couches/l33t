package main

// 1ms, 73%. 2.2mb, 93.5%
func strStr(haystack string, needle string) int {
  if haystack == needle {
    return 0
  }
	nlen := len(needle)
	for i := 0; i < len(haystack)-nlen+1; i++ {
    substr := haystack[i:i+nlen]
    if needle == substr {
      return i
    }
	}

	return -1
}
