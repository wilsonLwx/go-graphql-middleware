package main

import (
	"fmt"
)


func Adder() func(int)int{
	var x int
	return func(d int)int{
		x += d
		return x
	}
}

func testClosure1(){
	f := Adder()
	ret := f(1)
	fmt.Printf("f(1):%d\n",ret)
	ret = f(20)
	fmt.Printf("f(20):%d\n",ret)
	ret = f(100)
	fmt.Printf("f(100):%d\n",ret)
}

func main(){
	testClosure1()
}