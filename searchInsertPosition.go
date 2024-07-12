package main

// Something about this solution I'm not super happy about. Regardless, 3ms 71%, 3.1mb 73.1%
func searchInsert(nums []int, target int) int {
  if len(nums) == 1 { 
    if nums[0] >= target { return 0 }
    return 1
  }
  mid := len(nums) / 2
  mid_v := nums[mid]
  if mid_v == target {
    return mid
  }

  if target < mid_v {
    return searchInsert(nums[:mid], target)
  } else {
    return mid + searchInsert(nums[mid:], target)
  }
}
