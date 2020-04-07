package algorithm_study

import (
	"fmt"
)

func main() {
	//cal(100)
	//insert(9)
	loop()

}
func loop()  {
	var b =  [3]int{0}
	for i := 0;i<=3 ;i++{
		b[i] = 0
		fmt.Printf("hello world")
	}
}
func print( a []int)  {
	for i:=0;i <len(a);i++{
		fmt.Println(a[i])
	}
	for _,v := range a {
		fmt.Println(v)
	}
}
func insert(val int)  {
	count := 0
	fmt.Printf("first count is ",count)
	a := []int{1,2,3,4,5}
	for {
		if count == len(a) {
			sum := 0
			for i := 0; i < len(a); i++ {
				sum = sum + a[i]
				fmt.Printf("this is %d,sum is %d", i, sum)
			}
			a[0] = sum
			count = 1
		}
		a[count] = val
		count++
		fmt.Printf("count is %d",count)
	}
}
func cal(n int)  {
	sum := 0
	for i:= 0;i<n;i++{
		sum = sum + f(i)
	}
	fmt.Println(sum)
}
func f(n int) int {
	sum := 0
	for i:=0;i<n;i++{
		sum = sum +i
	}
	return sum
}