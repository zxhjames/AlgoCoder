package main

import (
	"fmt"
	"time"
	"strconv"
)

func main(){
	var today string = fmt.Sprintf("%s",time.Now().Format("2006") +"-"+ time.Now().Format("01") + "-"+strconv.Itoa(11))
	fmt.Println(today)
	now:=time.Now()
	q := time.Date(now.Year(),now.Month(),12,0,0,0,0,now.Location())
	fmt.Println(q.Format("2006"),q.Format("01"),q.Format("02"))
}
