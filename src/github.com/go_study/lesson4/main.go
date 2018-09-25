package main

import (
    "fmt"
)

func main(){
    // 九九乘法表

    for i := 1; i <= 9; i++ {
        for j:= 1; j <= i; j++{
            fmt.Printf("%dx%d=%02d ",j,i,j *i)
        }
        fmt.Println("")
    }
}