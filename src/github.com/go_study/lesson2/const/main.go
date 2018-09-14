package main

import "fmt"

const(
	A = iota
	B
	C
	D = 8
	E
	F = iota
	G = iota
)

const(
	A1 = iota
	B1 = iota
)

func main(){
	fmt.Println("A =",A)
	fmt.Println("B =",B)
	fmt.Println("C =",C)
	fmt.Println("D =",D)
	fmt.Println("E =",E)
	fmt.Println("F =",F)
	fmt.Println("G =",G)
	fmt.Println("A1 =",A1)
}