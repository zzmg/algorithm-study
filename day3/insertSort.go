package main

import (
	"fmt"
)

func main() {
	a := []int{2,1,5,8,6,4,3}
	insertSort(a)
}


func insertSort(a []int)  {
	if len(a) <=1 {
		fmt.Printf("can not to sort!!!!")
		return
	}
	// first move then insert
	for i:=1;i<len(a);i++ {
		current := a[i] //1
		//index := i
		j := i -1
		for ;j>= 0 ;j--  {
			if a[j] > current {
				a[j+1] = a[j]
				fmt.Println(a)
				//index--
			} else {
				break
			}
			//fmt.Println(a)
		}
		a[j+1] = current
	}
	fmt.Println(a)
}