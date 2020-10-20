package main

import (
	"log"
	"math"
	"sort"
)

// 설명 : 두개의 배열에서 합집합을 구하고 합집합의 중간값을 구해라
// 조건 : 복잡성 = O(log (m+n))
// 조건의 O는 BIG O라고 불리며 주로 시간복잡도에 사용됩니다.
// 정확히는모릅니다. (이문제 풀면서 처음봤습습다.)
// 저는 이과정을 풀지못했습니다.
// 다른사람의 풀이를보고 이해를 한것입니다.
// 풀이과정은 총 3가지이며  조건에 해당하는 풀이식은 findMedianSortedArrays3입니다.
// https://blog.chulgil.me/algorithm/ 요 블로그를보시면 조금 도움이됩니다. (BIG O에 대해)
func main() {
	nums1, nums2 := []int{1, 2}, []int{3, 4, 5}
	log.Println(findMedianSortedArrays2(nums1, nums2))
}

// O(m+n) - 1
// 아래의 함수는 두개의 어레이를 한개로합치고 숫자순서대로 Sort를합니다.
// Sort이후 중간값을 찾기위해 합친 어레이의 길이의 SHIFT 연산을 해줍니다.(BIT연산)
// 설명 : 예를들어 길이가 2인 배열1과 길이가 3인 배열을 합칩니다. 그럼 총길이 5인 배열이 되겠죠?
// 이 배열은 1,2,3,4,5인 숫자로 이루워진 배열이라고 가정하고 중간값은 3입니다.
// 5를 bit로 표현하면 0101(순서대로 나열하면 8,4,2,1)  3번째인 4와 4번째가 1이됩니다.  4는 0100
// 아래의 함수는 len(배열) >> 1을했습니다. 이것은 0101에서 >> 1 하면 010 , 0100 에서 >>1은 0010 5와같이 2가 됩니다.
// 즉, (배열[2] + 배열[2]) / 2.0이 됩니다.  이건 3+3 /2.0 이 돼고 결과적으로 3이라는 값이 출력이됩니다.
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	return float64(nums1[len(nums1)>>1]+nums1[(len(nums1)-1)>>1]) / 2.0
}

// O(m+n) - 2
// leetcode의 문제의내용중 이렇게 적혀있습니다. (정렬된 배열 1과 2의 크기는 m과 n으로 각각 표기한다.)
// 아래의 함수는 함수에다가 주석을 달겠습니다.
//ex) 할당값 배열1 = 1,2  배열2 = 4,5
// 1회차
// for 0 < 2 || 0 < 2 True이므로 loop
// 0 < 2 && 0 < 2 이고 배열1[0] < 배열2[0] 이므로 작은값인 배열1[0]을 배열3[0]에 넣습니다. K ++ i ++
// 2회차
// for 1 < 2 || 0 < 2 True이므로 loop
// 1 < 2 && 0 < 2 이고 배열1[1]보다 배열2[0]이 크므로 배열3[1]은 배열1[1]값 k ++ i ++
// 3회차
// 2 < 2 || 0 < 2 True 이므로 loop
// 2<2 && 0< 2 이므로  배열3[2]은 배열2[0] k++ i++
// 4회차
// 2<2 || 1<2 True이므로 loop
// 2<2 && 1<2 이므로 배열3[3]은 배열2[1] k++ i++
// 5회차
// 2<2 || 2<2 False이므로 loop종료
// 결과 배열3[(2+2) >> 1(2)] + 배열3[(2+2-1) >> 1(1)] 1,2,3,4 에서 2와 1에 해당하는 (2+3) / 2 인 2.5가 됩니다.

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2) // 배열 1과 2를 각각 m과 n으로 선언합니다.
	nums := make([]int, m+n)       // 배열 3을 int배열로만들고 길이를 m+n으로 선언합니다.
	i, j, k := 0, 0, 0             // i,j,k라는 변수들을 선언합니다 (배열들을 순환하기위한 변수)
	for i < m || j < n {           // i가 m보다 작거나  j가 n보다 작으면 for 루프를 진행합니다.
		if i < m && j < n { // i가 m보다 작고 j가 n보다 작을때
			if nums1[i] < nums2[j] { // 배열1[i] 보다 배열2[j]가 작으면
				nums[k] = nums1[i] // 배열3[k]는 배열1[i]값을 할당
				i++                //i는 사용되었으므로 +1
			} else { // if의 조건이아니면
				nums[k] = nums2[j] // 배열3[k]는 배열2[j]값으로 할당
				j++                // j가 사용되었으므로 +1
			}
		} else if i < m { // i보다 m이 작으면
			nums[k] = nums1[i] // 배열3[k]는 배열1[i]를 할당
			i++                //i는 사용되었으므로 +1
		} else {
			nums[k] = nums2[j] // 배열3[k]는 배열2[j]를 할당
			j++                //j가 사용되었으므로 +1
		}
		k++ //루프가 끝났으므로 k값 +1
	}
	return float64((nums[(m+n)>>1] + nums[((m+n-1)>>1)])) / 2.0 //배열 3의 m+n에서 right shift연산 1을 해줍니다 또 배열3의 m+n -1에서 right shift연산 1을하고 두값을 더한후 2로 나눕니다.
}

//O(lg(M+N))

func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2) // 배열1과 배열2의 길이를 변수 l로 선언
	if l%2 == 0 {                // l이 짝수면
		// find mid and mid - 1
		v1 := findKth(nums1, 0, nums2, 0, l/2)
		v2 := findKth(nums1, 0, nums2, 0, (l+2)/2)
		return (float64(v1) + float64(v2)) / 2.0
	} else { // 홀수면
		// find mid
		return float64(findKth(nums1, 0, nums2, 0, (l+1)/2))
	}
}

func findKth(A []int, pa int, B []int, pb int, k int) int {
	// A: [pa..len(A) - 1]
	// B: [pb..len(B) - 1]
	// k: kth [1..n], starting from 1
	m, n := len(A), len(B)
	if pa >= m {
		return B[pb+k-1]
	}
	if pb >= n {
		return A[pa+k-1]
	}

	if k == 1 {
		return int(math.Min(float64(A[pa]), float64(B[pb])))
	}
	va := math.MaxInt64
	if pa+k/2-1 < m {
		va = A[pa+k/2-1]
	}
	vb := math.MaxInt64
	if pb+k/2-1 < n {
		vb = B[pb+k/2-1]
	}

	if va <= vb {
		return findKth(A, pa+k/2, B, pb, k-k/2)
	} else {
		return findKth(A, pa, B, pb+k/2, k-k/2)
	}
}
