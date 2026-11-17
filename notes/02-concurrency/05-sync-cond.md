Представим ситуацию: функция должна начать выполняться, когда наступило какое-то событие. К примеру, она должна добавить элемент в очередь, когда освободится место, или изменится состояние какого-то флага, или счётчик дойдёт до какого-то определённого числа.

Обычно такую ситуацию разрешают при помощи каналов. Используется следующий факт: “Чтение из закрытого канала не блокируется.”

```go
var wg sync.WaitGroup
ch := make(chan struct{})

for i := range 3 {
	wg.Go(func() {
		<-ch
		fmt.Printf("gorouitne %v is running!\n", i)
	})
}

fmt.Println("waiting 500 ms...")
time.Sleep(500 * time.Millisecond)
fmt.Println("... done")
close(ch)

wg.Wait()
```

Но таким способом нельзя воспользоваться повторно, так как закрытие уже закрытого канала — паника!

Тогда используют Cond. Его базовый синтакс:

```go
var wg sync.WaitGroup
var mu sync.Mutex
cond := sync.NewCond(&mu)
ready := false

for i := range 3 {
	wg.Go(func() {
		mu.Lock()
		for !ready {
			cond.Wait() // unlock, sleep... ...wake, lock
		}
		mu.Unlock()

		fmt.Printf("goroutine %v started working.\n", i)
	})
}

fmt.Println("sleeping for 500 ms...")
time.Sleep(500 * time.Millisecond)
fmt.Println("... done.")

mu.Lock()
ready = true
mu.Unlock()

cond.Broadcast()
wg.Wait()
```

Почему используется for, а не if? Вроде бы, если условие не готово, то просто блокируемся. Станет готово — освободимся.

На самом деле, в этом блоке происходят следующие вещи:

```go
mu.Lock()
for !ready {
    cond.Wait()   // a) встать в очередь + unlock, чтобы могло случиться mu.Lock(); ready = true; mu.Unlock()
                  // b) ожидаем Broadcast'а или Signal'а
                  // c) дождались! просыпаемся
                  // d) пробуем Lock() <- можем ждать других какое-то время
                  // e) re-lock succeeded, Wait returns
}
mu.Unlock()
```

И может случиться так, что условие (ready) изменилось после пробуждения. Значит, его нужно проверить снова.

Используя базовый синтакс (`newCond(&mu) + for cond.Wait()`), напишем пример, когда подойдёт только Cond — пробуждение/засыпание **несколько** раз:

```go
type Gate struct {
    mu     sync.Mutex
    cond   *sync.Cond
    paused bool
}
func NewGate() *Gate {
    g := &Gate{}
    g.cond = sync.NewCond(&g.mu)
    return g
}
func (g *Gate) Pause() {
    g.mu.Lock()
    g.paused = true
    g.mu.Unlock()
}
func (g *Gate) Resume() {
    g.mu.Lock()
    g.paused = false
    g.mu.Unlock()
    g.cond.Broadcast() // wake ALL paused workers
}
func (g *Gate) WaitIfPaused() {
    g.mu.Lock()
    defer g.mu.Unlock()
    for g.paused {
        g.cond.Wait()
    }
}
```

```go
gate := NewGate()
var wg sync.WaitGroup
ctx, cancel := context.WithCancel(context.Background())

for id := range 3 {
	wg.Go(func() {
		for step := 1; ; step++ {
			wg.WaitIfPaused() // runs cond.Wait() inside for paused loop

			select {
			case <-ctx.Done():
				return
			default:
			}

			fmt.Printf("worker: %v, step %v\n", id+1, step)
			time.Sleep(300 * time.Millisecond)
		}
	})
}
```

```go
time.Sleep(1 * time.Second)
fmt.Println(">>> PAUSE")
gate.Pause()

time.Sleep(1 * time.Second)
fmt.Println(">>> RESUME")
gate.Resume()

time.Sleep(1 * time.Second)
fmt.Println(">>> PAUSE again")
gate.Pause()

time.Sleep(1 * time.Second)
fmt.Println(">>> RESUME again")
gate.Resume()

time.Sleep(1 * time.Second)
cancel()
wg.Wait()
```
