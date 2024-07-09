package main

import (
	"math"
)

func isPalindrome(x int) bool {
  if x < 0 {
    return false
  }

  digits := 0
  y := x

  for y > 0 {
    digits++
    y /= 10
  }

  left := int(math.Pow10(digits - 1))
  right := 1

  for i := 0; i < digits; i++ {
    a := x / left % 10
    b := x / right % 10

    if a != b {
      return false
    }

    left /= 10
    right *= 10
  }

  return true
}
