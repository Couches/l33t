package main

// Also not a terrible solution. Beats 100% at 2.20mb, but only 75% at 1ms.
// I also like how it was able to actually pass the tests. A great start.
func plusOne(digits []int) []int {
  start := len(digits) - 1
  digits[start] += 1
  for i := start; i > 0; i-- {
    if digits[i] > 9 {
      digits[i - 1] += 1
      digits[i] = 0
    } else {
      return digits
    }
  }

  if digits[0] > 9 {
    digits[0] = 0
    digits = append([]int{1}, digits...)
  }

  return digits
}

// Ok solution. Didn't pass because some cases were higher than integer max. Whatever.
func plusOneOld(digits []int) []int {
  sum := 0
  for _, digit := range digits {
    sum *= 10
    sum += digit
  }

  sum += 1

  res := []int{}

  for sum > 0 {
    res = append([]int{sum % 10}, res...)
    sum /= 10
  }

  return res
}
