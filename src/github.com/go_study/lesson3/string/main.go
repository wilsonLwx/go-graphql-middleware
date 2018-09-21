package main

import (
	"fmt"
)


func testReverseStringV1(){
	var bytes []byte = []byte("I love you!")
	for i := 0; i < len(bytes)/2; i++{
		tmp := bytes[len(bytes)-i-1]
		bytes[len(bytes)-i-1] = bytes[i]
		bytes[i] = tmp
	}
	fmt.Printf("%s\n",string(bytes))
}


func testReverseStringV2(){
	var r []rune = []rune("Hello 中国")
	for i := 0; i < len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp 
	}
	fmt.Println(string(r))
}

func testHuiwenString(){
	var str = "上海自来水来自海上"
	var r []rune = []rune(str)
	for i := 0; i < len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
	}
	str2 := string(r)
	if str == str2 {
		fmt.Printf("%s is hui wen\n", str2)
	}else{
		fmt.Printf("%s is not hui wen\n", str2)
	}
}

func main(){
	// testReverseStringV2()
	testHuiwenString()
}