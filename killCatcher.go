package killCatcher

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	sleepTime = 5 * time.Second
)

type KillCatcher struct {
	postKillFunc func() error
	executed     bool
}

func New(f func() error) *KillCatcher {
	return &KillCatcher{postKillFunc: f, executed: false}
}

func (k *KillCatcher) Listen(errgroup errgroup.Group) {
	errgroup.Go(
		func() error {
			term := make(chan os.Signal, 1)
			kill := make(chan os.Signal, 1)
			signal.Notify(term, syscall.SIGTERM)
			signal.Notify(kill, syscall.SIGKILL)
			defer func() {
				signal.Stop(term)
				signal.Stop(kill)
			}()
			for {
				select {
				case <-term:
					fmt.Printf("== Received SIGTERM signal, executing provided function ==\n")
					defer func() {
						k.executed = true
					}()
					if err := k.postKillFunc(); err != nil {
						fmt.Printf("== Error while executing function post SIGTERM : %v\n", err)
						return fmt.Errorf("error while executing function post SIGTERM : %v", err)
					}
					return nil
				case <-kill:
					fmt.Printf("== Received SIGKILL signal, exiting ==\n")
					if !k.executed {
						return fmt.Errorf("the post SIGTERM function did not return")
					}
					return nil
				default:
					fmt.Printf("== DEBUG TICK ==\n")
					time.Sleep(sleepTime)
				}
			}
		})
}
