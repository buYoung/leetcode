package main

func main() {
	var nums, target = []int{3, 2, 4}, 6

	twoSum1(nums, target)
}

// twosum3의 과정 3 2 4에 6인경우
// 1회차   i = 0
// 0,false = maps[6 - 3] // maps[3]은 초기값이므로 0 과 false 반환
// false이므로 maps[3] = 0
// 2회차  i = 1
// 0,false = maps[6 - 2] // maps[4]는 초기값이므로 0 과 false 반환
// false 이므로 maps[2] = 1
// 3회차 i = 2
// 1, true = maps[6-4] // maps[2]는 2회차에서 값을 줬으므로 1 과 true 반환
// true 이므로 return []int{2,1}

// twosum3의 과정 2 7 11 15에 9인경우
// 1회차   i = 0
// 0,false = maps[9-2] // maps[7]은 초기값이므로 0 과 false 반환
// false이므로 maps[2] = 0
// 2회차  i = 1
// 0,true = maps[9-7] // maps[2]는 1회차에 값을 줬으므로 0 과 false 반환
// true 이므로 return []int{0,1}
// 3회차 i = 2
// 0, false = maps[9-11] // maps[-2]은 초기값이므로 0 과 false 반환
// false이므로 maps[11] = 2
// 4회차 i = 3
// 0, false = maps[9-15] // maps[-6]은 초기값이므로 0 과 false 반환
// false이므로 maps[15] = 3

func twoSum3(nums []int, target int) []int { // 4ms 3.4MB
	maps := make(map[int]int, len(nums)) //버퍼가 nums의 길이인 map[int]int을 만듭니다 (버퍼쓰는이유는 memory 관리)
	for i, _ := range nums {             // nums를 루프합니다.
		val, ok := maps[target-nums[i]] // map의 키를 정의합니다 (val : map[key]의 값 ok : map[key] 존재하는지)
		if !ok {                        // 새로운 key일경우
			maps[nums[i]] = i // 값을 넣어준다.
		} else { // 새로운 값이아니라면
			return []int{i, val} // 반환
		}
	}
	return []int{-1, -1} // 값이 없을경우
}

//twosum3와 비슷하므로 설명 생략
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

//returnindex []int로 선언
// loop를 많이 사용하므로 느릴수밖에없다. 단, 메모리는 적게먹는다.
func twoSum1(nums []int, target int) []int { //runtime, memeory 36ms 2.9MB
	for i := 0; i < len(nums); i++ { // nums를 0 ~ nums의 길이만큼 루프
		for j := i + 1; j < len(nums); j++ { // nums의 뒤에값도 비교할수있게 0~ nums의 길이만큼 루프
			if nums[i]+nums[j] == target { // nums[i] + nums[j]가 target과 같은지 비교
				return []int{i, j} // 반환
			}
		}
	}
	return []int{-1, -1}
}
