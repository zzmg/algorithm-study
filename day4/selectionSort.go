package main

import "fmt"

func main() {
	a := []int{2,4,1,6,8,10}
	selectSort(a)
}

func selectSort(a []int)  {
	if len(a) <=1 {
		return
	}
	for i:= 0;i < len(a) ;i++ {
		minIndex := i
		for j := i+1 ;j < len(a);j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i],a[minIndex] = a[minIndex],a[i]
	}
	fmt.Println(a)
}
