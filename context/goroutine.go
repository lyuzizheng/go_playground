package context

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func TestContext(ctx context.Context)  {
	time.Sleep(time.Second * 1)

	fmt.Println("After sleep I wake up")
	fmt.Println("I sleep again")

	time.Sleep(time.Second * 5)

	fmt.Println("Get up")



}

func TestContextCancel()  {
	ctx := context.Background()
	ctx1, cancel := context.WithTimeout(ctx, time.Second*1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		TestContext(ctx1)

	}()
	cancel()
	wg.Wait()



}


func TestContextParent()  {
	parent, pCancel := context.WithCancel(context.Background())
	child, _ := context.WithCancel(parent)
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go work(wg, child)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	defer signal.Stop(c)

	select {
	case <-c:
		pCancel()
		fmt.Println("Waiting everyone to finish...")
		wg.Wait()
		fmt.Println("Exiting")
		os.Exit(0)
	}



}


func work(wg *sync.WaitGroup, ctx context.Context) {
	done := false
	wg.Add(1)
	for !done {
		fmt.Println("Doing something...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			done = true
		default:

		}
	}
	wg.Done()
}