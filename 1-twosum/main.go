package main

func main() {
	var nums, target = []int{3, 2, 4}, 6

	twoSum1(nums, target)
}

func twoSum3(nums []int, target int) []int { // 4ms 3.4MB
	maps := make(map[int]int, len(nums))
	for i, _ := range nums {
		val, ok := maps[target-nums[i]]
		if !ok {
			maps[nums[i]] = i
		} else {
			return []int{i, val}
		}
	}
	return []int{-1, -1}
}
func twoSum2(nums []int, target int) []int { // 4ms 3.8MB
	maps := make(map[int]int)
	for i, v := range nums {
		_, ok := maps[v]
		if !ok {
			maps[target-v] = i
		} else {
			return []int{maps[v], i}
		}
	}
	return []int{-1, -1}
}

func twoSum1(nums []int, target int) []int { //runtime, memeory 36ms 2.9MB
	var returnindex []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				returnindex = append(returnindex, i)
				returnindex = append(returnindex, j)
				return returnindex
			}
		}
	}
	return []int{-1, -1}
}
