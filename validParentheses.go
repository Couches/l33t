package main

// Beats 73% at 1ms. Memory, 2.3mb beating 72.4%
func isValidParentheses(s string) bool {
  var stack []rune
  p_map := map[rune]rune {
    '}': '{',
    ')': '(',
    ']': '[',
  }

  for _, c := range s {
    other, ok := p_map[c]
    if ok {
      size := len(stack)
      if size == 0 {
        return false
      }

      if stack[size - 1] != other {
        return false
      }

      stack = stack[:size - 1]
    } else {
      stack = append(stack, c)
    }
  }

  if len(stack) > 0 {
    return false
  }
  
  return true
}
