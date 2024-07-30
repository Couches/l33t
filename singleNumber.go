package main

// Brute force solution. Wanted to get one out of the way and call it a night
func singleNumber(nums []int) int {
  counts := map[int]int {}
  for _, num := range nums {
    counts[num] = counts[num] + 1
  }

  for num, count := range counts {
    if count == 1 {
      return num
    }
  }

  return 0
}
