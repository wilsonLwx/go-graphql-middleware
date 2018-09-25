package main

import (
	"fmt"
)

func insert_sort(a [8]int)[8]int{
	for i := 1; i < len(a); i++{
		for j := i; j > 0; j--{
			if a[j] < a[j-1]{
				a[j-1],a[j] = a[j],a[j-1] 
			}else{
				break
			}
		}
	} 
	return a
}

func select_sort(a [8]int)[8]int{
	for i := 0; i < len(a); i++{
		for j := i + 1; j < len(a);j++{
			if a[j] < a[i]{
				a[i],a[j] = a[j],a[i]
			} 
		}
	}
	return a
}

func main(){
	var i [8]int = [8]int{8,4,3,2,9,10,0,6}
	ret := select_sort(i)
	fmt.Println(i)
	fmt.Println(ret)
}