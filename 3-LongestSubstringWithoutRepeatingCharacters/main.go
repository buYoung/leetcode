package main

import (
	"log"
)

// 이 문제는 엄청어려웠다. 이런문제도있구나 하게되는 수학과같은문제랄까..
// 처음에는 단순하게 반복안되는 문자열을 찾는줄알았는데 (이래서 영어랑 국어를 제대로 배워야한다..)
// 그게아니였다...
// 문제는 "반복되지않는 가장긴 서브 문자열 찾기이다"
// 뭔말인지 이해가 안갈수있다. (나도그랬으니까)
// 예를들어 abc는 반복되지않는 문자열이다.
// abcabc는 반복되는 문자열이다 여기서 반복안된문자열은 abc이다 (처음 찾는걸기준으로)
// abcabcbb는 반복안된문자열은 abc이다 왜 abc일까?
// 생각해보자  abc abc 는 결국 반복인데 그렇다면 반복되지않은 문자열을 저장해야한다.
// abc를 저장후 다시 a가나왔으니 다시 abcb 까지갔을꺼다.
// b가있으므로 중복이다 그러므로 4개가 아닌 3개이다 (abc) 문자열 찾기이므로 abc이고 갯수로는 3개이다

// 위의 과정만봐도 어렵다. 단순하게 해결한다면 당신은 천재

// 그다음예인 bbbbb이다
// 위에서 경험을 했다면 이건쉬울꺼다.
// b다음 b가있으므로 찾은 문자열은 b이고 갯수는 1개이다
// 이게 5번 반복되므로 결국은 b와 갯수 1개이다

// 그다음예 pwwkew
// pw다음 w가있으므로 찾은건 pw에 갯수 2개
// ww 이므로 찾은건 w에 1개
// wkew 이므로 찾은건 wke에 3개

//이해하면 쉬울꺼같지만 안쉽다.

func main() {
	s := "pwwkew"
	log.Println(lengthOfLongestSubstring2(s))
}

func lengthOfLongestSubstring2(s string) int {
	var index [128]int
	var result int
	for i, j := 0, 0; j < len(s); j++ {
		if index[s[j]] > i {
			i = index[s[j]]
		}
		if j-i+1 > result {
			result = j - i + 1
		}
		index[s[j]] = j + 1
		log.Println(index[s[j]], string(s[j]), s[j], result, j, i, j-i+1)
	}
	return result
}
