Если работаем с обычной WaitGroup’ой, то приходится самим собирать ошибки при помощи каналов/общей памяти под mutex’ом.

Errgroup решает эту проблему. Как только одна горутина вернёт ошибку, отменится контекст egctx. Соответственно, другие горутины могут проверять его завершение через `case <-egctx.Done()`.

```go
eg, egctx := errgroup.WithContext(ctx)

eg.Go(func() {
	time.Sleep(10 * time.Millisecond)
	return errors.New("users: timeout")
})

eg.Go(func() error {
	select {
	case <-time.After(time.Second):
		return nil
	case <-egctx.Done():
		fmt.Println("orders: отменён")
		return egctx.Err()
	}
})

err := g.Wait()
fmt.Println("wait:", err)
```

```
orders: canceled
Wait: users: timeout
```
