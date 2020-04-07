package algorithm_study

import (
	"fmt"
)

type Boy struct {
	no int
	next *Boy
}

func AddBoy(num int) *Boy {
	head := &Boy{}
	curBoy := &Boy{}
	if num < 1{
		return head
	}
	for i := 1;i<=num;i++{
		boy := &Boy{
			no:   i,
			next: nil,
		}
		if i == 1 {
			head = boy
			curBoy = boy
			curBoy.next = head
		} else {
			//curBoy = boy
			//head.next = boy
			//curBoy.next = head
			curBoy.next = boy
			curBoy = boy
			curBoy.next = head
		}
	}
	return head
}

func ShowBoy(head *Boy)  {
	temp := head
	for {
		fmt.Println("boy:",temp.no)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func playGame(first *Boy,startNo,countnum int)  {
	if first.next == nil {
		fmt.Println("就一个小孩玩游戏，玩不了")
		return
	}
	tail := first
	for {
		if tail.next == first {
			break
		}
		tail = tail.next
	}
	startNode := first
	for   {
		if startNode.no == startNo {
			break
		}
		startNode = startNode.next
		tail = tail.next
	}
	for {
		for i:=1;i<=countnum;i++ {
			startNode = startNode.next
			tail = tail.next
		}
		fmt.Println("退出的小孩：",startNode.no)
		if startNode == tail {
			fmt.Println("最后剩下的小孩编号：",startNode.no)
			break
		}
		startNode = startNode.next
		tail.next = startNode
	}

}
func main() {
	head := AddBoy(100)
	ShowBoy(head)
	playGame(head,3,2)

}
