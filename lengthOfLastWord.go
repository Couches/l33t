package main

import (
	"strings"
)

// I like this one better. Beats 75% at 1ms, 2.38mb beating 78%.
func lengthOfLastWord(s string) int {
  length := len(s)
  start := 0

  for i := 1; i < length; i++ {
    c := s[length - i]
    if c != ' ' {
      start = length - i
      break;
    }
  }

  i := 0
  for i = 1; i <= start; i++ {
    c := s[start - i]
    if c == ' ' {
      return i
    }
  }

  return i
}

// Kind of a cheap solution. 75% at 1ms, 2.41mb 24.24%. Feels cheaty.
func lengthOfLastWordBS(s string) int {
  fields := strings.Fields(s)
  return len(fields[len(fields)-1])
}
