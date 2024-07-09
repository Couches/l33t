package main

func twoSum(nums []int, target int) []int {
  others := make(map[int]int)

  for i, num := range nums {
    other, ok := others[target - num]
    if ok {
      return []int{other, i}
    }
    others[num] = i
  }

  return []int{}
}
