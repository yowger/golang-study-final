package main

import (
	"context"
	"fmt"
	"time"
)

// controlling timeouts
// cancelling go routines
//  and passing metadata across your Go applications

func main() {
	ctx := context.Background()
	exampleTimeout(ctx)
}

func exampleTimeout(ctx context.Context) {

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Called the API")
	case <-ctxWithTimeout.Done():
		fmt.Println("on no timed out expired", ctxWithTimeout.Err())
		// do some logic to handle the timeout
		// logging, setting response (try again)
	}
}

// ctx := context.Background()
// ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 4*time.Second)
