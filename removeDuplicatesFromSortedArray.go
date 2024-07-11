package main

// Runtime of 4ms, beating 85%. 4.5mb, beating 90%
func removeDuplicates(nums []int) int {
  length := len(nums)
  if length <= 1 {
    return 1
  }

  pointer := 0

  for i := 0; i < length; i++ {
    if nums[pointer] != nums[i] {
      pointer++
    }
    nums[pointer] = nums[i]
  }

  return pointer + 1
}
