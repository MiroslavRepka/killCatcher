package killCatcher

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
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

// Listen listens for SIGTERM/os.Interrupt and executes provided function in killCatcher if received
func (k *killCatcher) Listen() error {
	// create channel to listen for SIGTERM/Interrupt
	term := make(chan os.Signal, 1)
	// Listen for signal
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	// close channel before exit
	defer signal.Stop(term)
	// signal received
	<-term
	// execute postKillFunc function
	if err := k.postKillFunc(); err != nil {
		return fmt.Errorf("error while executing function post SIGTERM : %v", err)
	}
	return nil

}
