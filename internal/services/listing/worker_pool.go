package listing

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

type Result struct {
	url string
	err error
}

const workersCount = 3

func rpcGenerator(ctx context.Context, rpcs []string) <-chan string {
	resultCh := make(chan string)

	go func() {
		defer close(resultCh)

		for _, rpc := range rpcs {
			select {
			case <-ctx.Done():
				return
			case resultCh <- rpc:
			}
		}
	}()

	return resultCh
}

func startWorkerPool(ctx context.Context, numWorkers int, jobsCh <-chan string, resultCh chan<- *Result) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for workerId := 1; workerId <= numWorkers; workerId++ {
		go func() {
			worker(ctx, workerId, jobsCh, resultCh)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()
}

func worker(ctx context.Context, workerId int, jobsCh <-chan string, resultCh chan<- *Result) {
	for {
		select {
		case <-ctx.Done():
			return
		case rpc, ok := <-jobsCh:
			if !ok {
				return
			}

			fmt.Println("Worker", workerId, "checking", rpc)

			_, err := http.Get(rpc)

			resultCh <- &Result{
				url: rpc,
				err: err,
			}

			if err != nil {
				return
			}
		}
	}
}
