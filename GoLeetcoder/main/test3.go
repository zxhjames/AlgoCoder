package main

import "fmt"

func main(){
	c:=make(chan int,3)
	c <- 1
	c <- 2
	c <-100
	close(c)
	k:= <-c
	fmt.Println(k)
	x1,ok:= <- c
	fmt.Println(x1,ok)
}