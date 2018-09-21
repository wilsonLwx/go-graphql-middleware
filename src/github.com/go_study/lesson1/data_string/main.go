package main

import (
	"strings"
	"fmt"
)

func main(){
	var c string
	c = "床前明月光，疑是地上霜"
	s := "举头望明月，低头思故乡"
	// fmt.Printf("c = %s\n",c)
	// c = c + "\n" +s
	// fmt.Printf("%s\n",c)
	m := fmt.Sprintf("%s\n%s\n",c,s)
	fmt.Printf("m:\n%s\n",m)

	ips := "10.108.34.30;10.108.34.31;10.108.34.32;"
	ipArray := strings.Split(ips,";")
	for i := 0; i<len(ipArray)-1; i++{
		fmt.Printf("ipArray = %s\n",ipArray[i])
	}
	result := strings.Contains(ips,"你好")
	fmt.Println("result =",result)

	sTr := "http://www.baidu.com"

	if strings.HasPrefix(sTr,"http"){
		fmt.Println("sTr is http url")
	}else{
		fmt.Println("sTr is not http url")
	}

	if strings.HasSuffix(sTr,"com"){
		fmt.Println("sTr is com url")
	}else{
		fmt.Println("sTr is not http url")
	}
	index := strings.Index(sTr,"baidu")
	fmt.Printf("index = %d\n",index)

	var str_array []string = []string{ "I" ,"love", "you"}
	new_str := strings.Join(str_array, " ")
	fmt.Println(new_str)
}