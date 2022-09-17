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
	killCatcher := killCatcher.New(postKill)
	var eg errgroup.Group
	eg.Go(killCatcher.Listen)
	eg.Go(logic)
	if err := eg.Wait(); err != nil {
		fmt.Printf("Got error while waiting for stuff to return : %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func logic() error {
	for {
		fmt.Printf("Main logic is here \n")
		time.Sleep(sleepTime)
	}
}

func postKill() error {
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d after SIGTERM in post SIGTERM function \n", i)
		time.Sleep(sleepTime)
	}
	fmt.Printf("Post SIGTERM function ended \n")
	return nil
}
