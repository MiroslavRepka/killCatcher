package main

import (
	"fmt"
	"os"
	"time"

	"github.com/MiroslavRepka/killCatcher"
	"golang.org/x/sync/errgroup"
)

const (
	sleepTime = 3 * time.Second
)

func main() {
	ch := make(chan struct{}, 1)
	killCatcher := killCatcher.New(postKill(ch))
	var eg errgroup.Group
	eg.Go(killCatcher.Listen)
	eg.Go(logic(ch))
	if err := eg.Wait(); err != nil {
		fmt.Printf("Got error while waiting for stuff to return : %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func logic(ch <-chan struct{}) func() error {
	return func() error {
		for {
			select {
			case <-ch:
				fmt.Printf("SIGTERM in main logic received")
				time.Sleep(sleepTime)
				return nil
			default:
				fmt.Printf("Main logic is here \n")
				time.Sleep(sleepTime)
			}
		}
	}
}

func postKill(ch chan<- struct{}) func() error {
	return func() error {
		for i := 0; i < 5; i++ {
			fmt.Printf("Iteration %d after SIGTERM in post SIGTERM function \n", i)
			time.Sleep(sleepTime)
		}
		fmt.Printf("Post SIGTERM function ended, sending info to channel \n")
		ch <- struct{}{}
		return nil
	}
}
