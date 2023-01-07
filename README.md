# killCatcher

[![Go Report Card](https://goreportcard.com/badge/github.com/MiroslavRepka/killCatcher)](https://goreportcard.com/report/github.com/MiroslavRepka/killCatcher)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)


Simple go module for detecting the pod deletion from Kubernetes cluster

## How to use it

Integration of this go module is quite simple. Inside your `main`, define `killCatcher` instance, and call `Listen()` in separate goroutine. It is up to you how you will manage those goroutines. 

In this example, we will use `errorGroup`.

```go
func main(){
    kc := killCatcher.New(postSigterm)
    var eg errgroup.Group
    eg.Go(killCatcher.Listen)
    eg.Go(yourApp)
    if err := eg.Wait(); err != nil {
        fmt.Printf("Got error in one of the goroutines : %v\n", err)
        os.Exit(1)
    }
    os.Exit(0)
}

func postSigterm() error{
    //logic to execute after SIGTERM
}

func yourApp() error {
    //main logic of your app
}
```

Lastly, do not forget to define `terminationGracePeriodSeconds` in you manifest file. By default, its set to `30s`, however, this might not suit your needs. Example can be found [here](test/pod.yaml).
