**Mutex**’ы нужны для решения проблем атомарности. Если 2 горутины делают `i++`, то может случиться следующее: обе одновременно прочитают 5, обе увеличат это значение на 1 и обе запишут в переменную 6. Потеряется одно увеличение. Если бы мы защитили секцию с `i++`, то такого бы не произошло: первая пришедшая горутина заблокировала бы доступ к этой секции до тех пор, пока не закончит свою работу, а вторая бы ждала:

```go
var wg sync.WaitGroup
var mu sync.Mutex
var value int

increment := func() {
	mu.Lock() // заблокировать; если уже заблокировано, то ждать, пока не разблокируется и заблокировать.
	defer mu.Unlock()
	value++
	fmt.Printf("incrementing: %v\n", value)
}
decrement := func() {
	mu.Lock()
	defer mu.Unlock()
	value--
	fmt.Printf("decrementing: %v\n", value)
}

for range 5 {
	wg.Go(increment)
}
for range 5 {
	wg.Go(decrement)
}

wg.Wait()
fmt.Printf("the result after all arithmetics: %v\n", value) // 0
```

Стоит использовать mutex’ы только там, где это необходимо, т.к. `Lock()` запрещает другим горутинам трогать участок кода и они начинают копиться. Получается bottleneck.

Эту ситуацию можно оптимизировать с помощью `RWMutex`: он хорошо подходит, когда читающих горутин заметно больше, чем пишущих. Правда, выигрыш появляется не сразу. По замерам автора книги, `RWMutex` начинает обгонять обычный `Mutex` лишь при ~2$^{13}$ = 8192 читателях, а до этого порога проигрывает из-за более сложной внутренней реализации.

- `RLock()` — блокирует ресурс для записи. Не блокирует для чтения. Другие горутины всё ещё могут сделать `RLock()`.
- `RUnlock()` — сработает, ресурс был заблокирован для чтения. Не сработает, если он был заблокирован для записи.
- `Lock()` — блокирует ресурс и для записи, и для чтения.

```go
type SafeConfig struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *SafeConfig) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.data[key]
}

func (c *SafeConfig) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

func main() {
	var wg sync.WaitGroup
	sf := &SafeConfig{
		data: make(map[string]string),
	}

	for i := range 10 {
		wg.Go(func() {
			val := sf.Get("theme")
			fmt.Printf("reader gorouitine %v: got theme = %q\n", i+1, val)
		})
	}

	for i := range 2 {
		wg.Go(func() {
			sf.Set("theme", "dark")
			fmt.Printf("write goroutine %v set theme to dark\n", i+1)
		})
	}

	wg.Wait()
}
```

Если скопировать mutex, то получится другой!
