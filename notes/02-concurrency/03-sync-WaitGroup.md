Основная задача — дождаться завершения всех горутин из группы.

Используем, когда нам или неважен результат работы горутины, или есть какой-то особый способ этот результат забрать.

Во всех остальных случаях лучше использовать не waitgroup’ы, а каналы и select’ы!

Старый способ:

```go
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine is doing smt that takes ~ 1 sec")
		time.Sleep(1 * time.Second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine is doing smt that takes ~ 0.5 sec")
		time.Sleep(500 * time.Millisecond)
	}()

	wg.Wait()
	fmt.Println("code after wg.Wait()")
}

```

Новый (Go 1.25+) — wg.Go(func() {…})

```go
func main() {
	var wg sync.WaitGroup

	wg.Go(func() {
		fmt.Println(...)
		time.Sleep(1 * time.Second)
	})

	wg.Go(func() {
		fmt.Println(...)
		time.Sleep(500 * time.Milliseconds)
	})

	wg.Wait()
	fmt.Println("code after wg.Wait()")
}
```

- для передачи параметров используется анонимная функция-обёртка

Можно переиспользовать ту же переменную `wg` после Wait’а, необязательно создавать новую!
