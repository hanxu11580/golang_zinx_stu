package main

import (
	"fmt"
	"time"
)

type T struct {
	Name string
}

type I interface {
	Rename(name string)
}

func (t *T) Error() string {
	temp := fmt.Sprintf("我是%s，我也不知道，我有什么错", t.Name)
	return temp
}

func say(s string, delay int, c chan string) {
	var res string
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		res += s
	}
	c <- res
}

type Arr struct {
	Str_Arr []int
}

func main() {
	// c := make(chan int, 2)
	// c <- 1
	// c <- 2
	// close(c)
	// for i := range c {
	// 	fmt.Println(i)
	// }
}
