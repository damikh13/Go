package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	eg, egctx := errgroup.WithContext(ctx)
	eg.SetLimit(10)
	eg.Go(func() error {
		time.Sleep(10 * time.Millisecond)
		return errors.New("users: timeout")
	})

	eg.Go(func() error {
		select {
		case <-time.After(time.Second):
			return nil
		case <-egctx.Done():
			fmt.Println("orders: canceled")
			return egctx.Err()
		}
	})

	err := eg.Wait()
	fmt.Println("Wait:", err)
}
