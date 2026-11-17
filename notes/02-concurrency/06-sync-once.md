Если нам нужно, чтобы какое-то действие выполнилось всего 1 раз, то можем использовать `sync.Once`:

```go
var wg sync.WaitGroup
var once sync.Once
var counter int

for i := range 100 {
	wg.Go(func() {
		once.Do(func() {
			counter++
			fmt.Printf("counter was increased by %v-th goroutine!\n", i)
		})
	})
}

wg.Wait()
fmt.Printf("counter = %v.\n", counter)
```

Если 99 горутин начали выполняться чуть позже единственной, то их исполнение блокируется на блоке `.Do()`, пока она не завершит свой вызов.

Можно подумать, что Once считает количество вызовов уникальных функций. Но это не так. Ведётся подсчёт того, сколько раз было сделано `.Do()`:

```go
var once sync.Once
var wg sync.WaitGroup

counter := 0

increment := func() {
	counter++
}
decrement := func() {
	counter--
}

once.Do(increment)
once.Do(decrement)

fmt.Println("ctouner:", counter) // 1, NOT 0!
```

Так как вызовы блокируются, то можно словить deadlock:

```go
var once sync.Once
once.Do(func() { // начало исполняться => все остальные .Do() должны блокировать
	once.Do(func() { // блокирован
	})

	// никогда не закочним исполняться!
})
```

Вот реальный пример использования — долгая операция:

```go
var once sync.Once
var config map[string]string

func getConfig() map[string]string {
    once.Do(func() {
        config = loadConfig() // slow
    })
    return config // everyone sees the fully loaded config
}
```

Как можно заметить, используется подход: объявляем рядом переменные once и ту, что мы хотим проинициализировать (config, db, connection, etc.). Начиная с Go 1.21+, придумали более прикладной подход — `OnceFunc`, `OnceValue`, `OnceValues`:

OnceFunc:

```go
setup := sync.OnceFunc(func() {
	fmt.Println("setting up...")
})

setup() // setting up...
setup() // nothing
setup() // nothing
```

OnceValue:

```go
getNumber := sync.OnceValue(func() {
	fmt.Println("computing...")
	time.Sleep(1 * time.Second)
	return 42
})

fmt.Println(getNumber()) // computing... 42
fmt.Println(getNumber()) // 42 <- instantly
fmt.Println(getNumber()) // 42 <- instantly
```

- если имеем getSlice(), то повторный вызов выдаст оригинал! возможно мутация.

OnceValues:

```go
readConfig := sync.OnceValues(func() (*Config, error) {
    data, err := os.ReadFile("config.json")
    if err != nil {
        return nil, err
    }
    var c Config
    return &c, json.Unmarshal(data, &c)
})

cfg, err := readConfig() // reads the file
cfg, err = readConfig()  // same cfg and same err, no second read
```

Обычный once глушит panic’у. Новые функции — нет:

```go
var once sync.Once
try := func() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("caught: %v.\n", r)
		}
	}()

	once.Do(func() {
		panic("boom")
	})
}

try() // caught: boom
try() // nothing
try() // nothing
```

```go
get := sync.OnceFunc(func() {
	fmt.Println("here")
	panic("boom")
})

try := func() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("caught: %v.\n", r)
		}
	}()

	get()
}

try() // here; caught: boom
try() // caught: boom <- f is not called again! value is repeated
try() // caught: boom
```
