package channel

import (
	"fmt"
	"time"
)

/*
All functions will be called separately in main:

which means main(){ChannelTest1()}
Type out the running result of the 7 runs,

A: stdout:5          B: nothing
C: deadlock fatal error
 */

func ChannelTest1 (){
	ch1 := make(chan int, 10)
	fmt.Println(<-ch1)
	ch1 <- 5
}
func ChannelTest2 (){
	ch1 := make(chan int)
	ch1 <- 5
	fmt.Println(<-ch1)
}
func ChannelTest3 (){
	ch1 := make(chan int, 10)
	ch1 <- 5
	fmt.Println(<-ch1)
}
func ChannelTest4 (){
	ch1 := make(chan int)
	go fmt.Println(<-ch1)
	ch1 <- 5
	time.Sleep(1 * time.Second)
}
func ChannelTest5 (){
	ch1 := make(chan int, 10)
	go fmt.Println(<-ch1)
	time.Sleep(1 * time.Second)
	ch1 <- 5
	time.Sleep(1 * time.Second)
}
func ChannelTest6 (){
	ch1 := make(chan int)
	ch1 <- 5
	go fmt.Println(<-ch1)
	time.Sleep(1 * time.Second)
}
func ChannelTest7 (){
	ch1 := make(chan int, 10)
	ch1 <- 5
	go fmt.Println(<-ch1)
	time.Sleep(1 * time.Second)
}




































func ChannelTest0 (){
	ch1 := make(chan int)
	go func(a int) {
		fmt.Println(a)
	}(<-ch1)
	ch1 <- 5
	time.Sleep(1 * time.Second)
}

