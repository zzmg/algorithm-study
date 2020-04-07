package main

import (
	"fmt"
)

func main() {
	a := []int{2,4,1,5,8,10}
	bubbleSort(a)
}

func bubbleSort(a []int) {
	if len(a) <= 1 {
		fmt.Printf("can not to sort")
		return
	}
	flag := false //define if or not to reversal
	for i:=0;i<len(a);i++ {
		for j:= 0;j <len(a)-1;j++{
			if a[j] > a[j+1] {
				a[j],a[j+1] =a[j+1],a[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println(a)
}