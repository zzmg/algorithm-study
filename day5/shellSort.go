package main

import "fmt"

func main() {
	a := []int{2,6,1,7,3,8,10}
	shellSort(a)
}

func shellSort(a []int)  {
	n := len(a)
	h := 1
	for h < n/3 {
		h = 3*h +1
	}
	for h >=1{
		for i := h ;i<n;i++ {
			for j:= i; j>=h && a[j] < a[j-h] ;j = j-h {
				a[j],a[j-h] = a[j-h],a[j]
			}
		}
		h /= 3

	}
	fmt.Println(a)
}