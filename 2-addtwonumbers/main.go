package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

//문제 자체가 난이도가있습니다.. listnode 를 어떻게 값을 주는지 고민만 5분~10분했네요
//다른답들을 보고나도 의문이 안풀려서 그냥 값을 넣으니까 이게 맞는방식이네요 ㅋㅋㅋ...
//문제 : 영어를 잘못하므로 패스
// l1의 val을 2 ,4 ,3 을주고 l2를 5 ,6 ,4를 주고 결과는 7 , 0, 8이 나와야한다.
// type 은 위와같이 만들고

func main() {
	var l1 = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	var l2 = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	log.Println(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)     // List node의 dummy를 만든다
	p, q, cur := l1, l2, dummy // p,q,cur에 각각 l1,l2,dummy를 넣어준다.
	carry := 0                 // carry는 0부터 시작한다 (leetcode의 solution에 나와있음!)
	for p != nil || q != nil { // solution을보면 p와 q는 l1과 l2의 val값으로 초기화한다.
		var x, y int  // x와 y를 임의적으로 초기화
		if p != nil { // p가 nil(값없는상태 포인터라 nil 가능)이면 x 는 p의 val값으로 초기화
			x = p.Val
		} else { // 아니면 x 는 0 (솔루션에 나와요!)
			x = 0
		}
		if q != nil { // q가 nil(값없는상태 포인터라 nil 가능)이면 y 는 q의 val값으로 초기화
			y = q.Val
		} else {
			y = 0
		}
		sum := carry + x + y // sum은 솔루션에서 x+y+carry라고 나옵니다. x,y를 초기화할때 carry도 초기화를했는데 필요가없더라구요
		// 그래서 뺏습니다 여기서 쓰도록
		carry = sum / 10                    //sum을 10으로 나눈값을 carry에 대입
		cur.Next = &ListNode{sum % 10, nil} // 다음 노드의 val값은 sum의 나머지를 구해줍니다 (%을 잘안쓰니까 저도까먹었어요)
		// 자세한건 구글검색 golang operator
		//여기서  nil이든 new(ListNode)든 값은 나옵니다. 포인트값이라 nil을줬어요

		cur = cur.Next // 더미인 cur을 cur.next로 지정합니다.
		if p != nil {
			p = p.Next
		} // p가 nil이면 p = p의 next를 대입
		if q != nil {
			q = q.Next
		}
	}

	if carry > 0 { // carry가 0보다크면
		cur.Next = &ListNode{carry, nil} // cur.next 를 next값이 없는상태로 val만 있는 ListNode의 포인터를 대입
	}
	return dummy.Next
}
