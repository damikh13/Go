Главная идея — ограничить количество горутин каким-то числом — N.

Пусть у нас есть канал с задачами, которые нужно обработать, и канал с результатами, куда нужно сложить то, что получилось в результате обработки. Давайте запустим 3 воркера, основной код которых — это возведение в квадрат:

```go
func main() {
	jobs := make(chan int)
	results := make(chan int)
	var wg sync.WaitGroup

	for wID := range 3 {
		wg.Go(func() { // 1
			for j := range jobs {
				results <- j * j
			}
		})
	}

	go func() { // 2
		defer close(jobs)
		for j := range 9 {
			jobs <- j
		}
	}()

	go func() { // 3
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("got result %v.\n", r)
	}
}
```

Как видно, возникает много сложностей при синхронизации при использовании каналов:

1. Воркеры должны запускаться только через `wg.Go()`, т.к. нужно дождаться их завершения и только после этого закрывать канал results. Иначе канал закрыт, а воркер пытается в него писать — `panic`.
2. Заполнение канала с заданиями происходит в отдельной горутине. Если заполнять их в main’е, то заблокируемся на попытке записать 4-ю: j1, j2, j3 достаны воркерами, возведены в квадрат и ждут, пока читатель будет готов прочесть; читатель ещё не готов — он в самом конце; jobs ← j4 блокируется (нет ещё одного воркера, который смог бы достать). Выходит, пытаемся положить 4-ю задачу на исполнение, а не можем, т.к. все воркеры заняты тем, что ждут, пока их результаты не прочтут.
3. Ожидание завершение работы всех воркеров и закрытие канала results тоже должно выполняться в отдельной горутине, т.к. иначе мы начнём ждать завершение горутин до того, как начали читать из results (⇒ до того, как воркеры смогли отдавать свои результаты)

Поэтому, если нам заранее известно, сколько всего работы нужно сделать, можем использовать более простой подход: results ← slice.

Работа = input[atomicIDX]

```go
input := []int{1, 2, 3, 4, 5}
N := len(input)
results := make([]int, N)
var wg sync.WaitGroup

var atomIDX atomic.Int64
for range N {
	wg.Go(func() {
		idx := atomIDX.Add(1) - 1
		results[idx] = input[idx] * 10
	})
}

wg.Wait()
for _, r := range results { // 10, 20, ..., 50
	fmt.Println(r)
}
```

Работа = <- jobs

```go
input := []int{1, 2, 3, 4, 5}
N := 9
jobs := make(chan int)
results := make([]int, N)
var wg sync.WaitGroup
```

```go
for range 3 {
	wg.Go(func() {
		for j := range jobs {
			results[j] = input[j] * 10
		}
	})
}
```

```go
for j := range N {
	jobs <- j
}
close(jobs)
```

```go
wg.Wait()

for i, v := range results {
	fmt.Printf("results[%v] = %v.\n", i, v)
}
```

Можем прикрутить контекст к последнему примеру. Будем проверять, отменён ли контекст. В случае, если отменён, выдача задач прекратится.

```go
func process(ctx context.Context, workersNum int, input []int, fn func(context.Context, int) int) ([]int, error) {
	results := make([]int, len(input))
	jobs := make(chan int) // индексы
	var wg sync.WaitGroup

	for range workersNum {
		wg.Go(func() {
			for j := range jobs {
				results[j] = fn(ctx, input[j])
			}
		})
	}

feed:
	for idx := range input {
		if err := ctx.Err(); err != nil {
			break
		}

		select {
		case jobs <- idx:
		case <-ctx.Done():
			break feed
		}
	}
	close(jobs)

	wg.Wait()
	return out, ctx.Err()
}
```

Такая проверка нужна в случае, если одновременно отменился контекст и воркер стал готов читать, чтобы не гадать с тем, что рандомно выберет select:

```go
if err := ctx.Err(); err != nil {
	break
}
```

Или отменён контекст (в этом случае прекращаем выдачу), или пришёл воркер, готовый принять задачу:

```go
select {
case jobs <-idx:
case <-ctx.Done():
	break feed
}
```
