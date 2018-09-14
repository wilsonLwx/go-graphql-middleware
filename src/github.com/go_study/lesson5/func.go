package main

import (
	"fmt"
)

func calc_v1(a int, b int ) (sum int, sub int){
	sum = a + b
	sub = a - b
	return
}

func calc_v2(a ...int ) (int,int) {
	sum := 0
	sub := 0
	for i:= 0; i < len(a);i++{
		sum += a[i]
		sub -= a[i]
	}
	return sum,sub
}

func calc_v3(a int,b ...int) int {
	sum := 0
	for i := 0; i < len(b);i++{
		sum += b[i]
	}
	return a + sum
}
func main(){
	// sum,sub := calc_v1(12,1)
	// sum,sub := calc_v2(20,1,3,6)
	sum := calc_v3(20)
	// fmt.Printf("sum=%d,sub=%d\n",sum,sub)
	fmt.Printf("sum=%d\n",sum)
}