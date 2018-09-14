package main

import (
	"fmt"
)

func exchange(x,y *int){
	fmt.Printf("before x=%d y=%d\n",*x,*y)
	*x,*y = *y,*x
	fmt.Printf("after x=%d y=%d\n",*x,*y)
}
func main(){
	var x int = 89
	var y int = 22
	exchange(&x,&y)
	fmt.Printf("in main x=%d y=%d\n",x,y)
}