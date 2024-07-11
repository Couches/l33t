package main

// 0ms, beats 100%. 2.35mb, beating 64.4%
func removeElement(nums []int, val int) int {
  // This is solved using a very similar approach to removeDuplicates.
  length := len(nums)
  pointer := 0

  for i := 0; i < length; i++ {
    nums[pointer] = nums[i]
    if nums[pointer] != val {
      pointer++
    }
  }

  return pointer + 1
}
