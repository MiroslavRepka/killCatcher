package killCatcher

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	sleepTime = 5 * time.Second // interval in which the goroutine will check for SIGTERM
)

// killCatcher will execute provided function postKillFunc after the SIGTERM is received
// if the function returns an error, it will be propagated to the main program for handling
type killCatcher struct {
	postKillFunc func() error //function which will be executed after SIGTERM
}

// New returns new killCatcher with provided function which will be executed after SIGTERM
func New(f func() error) *killCatcher {
	return &killCatcher{postKillFunc: f}
}

// Listen listens for SIGTERM and executes provided function in killCatcher if received
func (k *killCatcher) Listen() error {
	//create channel to listen for SIGTERM
	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGTERM)
	defer func() {
		signal.Stop(term)
	}()
	// Listen for SIGTERM
	for {
		select {
		// the SIGTERM received
		case <-term:
			fmt.Printf("== Received SIGTERM signal, executing provided function ==\n")
			// execute postKillFunc function
			if err := k.postKillFunc(); err != nil {
				fmt.Printf("== Error while executing function post SIGTERM : %v\n", err)
				return fmt.Errorf("error while executing function post SIGTERM : %v", err)
			}
			return nil
		// default is used so the execution is not blocked
		default:
			fmt.Printf("== DEBUG TICK ==\n")
			// sleep for sleepTime
			time.Sleep(sleepTime)
		}
	}
}
