package main

import (
	"fmt"
	"time"
	"sync"
)

func main(){
	//var today string = fmt.Sprintf("%s",time.Now().Format("2006") +"-"+ time.Now().Format("01") + "-"+strconv.Itoa(11))
	//fmt.Println(today)
	//now:=time.Now()
	//q := time.Date(now.Year(),now.Month(),12,0,0,0,0,now.Location())
	//fmt.Println(q.Format("2006"),q.Format("01"),q.Format("02"))

	fmt.Println(time.Now().Format("2006-01-02"))
}
func test(){
	var mutex sync.Mutex
	var s int
	fmt.Println(time.Now())
	mutex.Lock()
	for i:=0;i<10000000;i++{
		s++
	}
	mutex.Unlock()
	fmt.Println(time.Now())
}