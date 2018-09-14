package main


import (
	"fmt"
	"time"
)

func testTime(){
	now := time.Now()
	fmt.Printf("current time:%v\n",now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n",year,month,day,hour,minute,second)
	timestamp := now.Unix()
	fmt.Printf("timestamp is %d\n",timestamp)
}

func testTimestamp(timestamp int64){
	timeObj := time.Unix(timestamp,0)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n",year,month,day,hour,minute,second)
}

func testTicker(){
	ticker := time.Tick(time.Second)
	for i := range ticker{
		fmt.Printf("%v\n",i)
	}
}

func testFormat(){
	now := time.Now()
	//timeStr := now.Format("2006/01/02 15:04:05")
	timeStr := fmt.Sprintf("%02d\\%02d\\%02d %02d:%02d:%02d\n",
		now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Printf("%s\n",timeStr)
}
func main(){
	// testTime()
	// timestamp := time.Now().Unix()
	// testTimestamp(timestamp)
	// testTicker()
	testFormat()
}