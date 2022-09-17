package killCatcher

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	sleepTime = 5 * time.Second
)

type KillCatcher struct {
	postKillFunc func() error
}

func New(f func() error) *KillCatcher {
	return &KillCatcher{postKillFunc: f}
}

func (k *KillCatcher) Listen() error {
	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGTERM)
	defer func() {
		signal.Stop(term)
	}()
	for {
		select {
		case <-term:
			fmt.Printf("== Received SIGTERM signal, executing provided function ==\n")
			if err := k.postKillFunc(); err != nil {
				fmt.Printf("== Error while executing function post SIGTERM : %v\n", err)
				return fmt.Errorf("error while executing function post SIGTERM : %v", err)
			}
			return nil
		default:
			fmt.Printf("== DEBUG TICK ==\n")
			time.Sleep(sleepTime)
		}
	}
}
