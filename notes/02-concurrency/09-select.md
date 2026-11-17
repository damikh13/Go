При входе в `select` вычисляются все каналы из `case`’ов:

- если хотя бы один `case` готов, то выбирается он; если готово сразу несколько, срабатывает случайный
- если готовых `case`’ов нет, но есть `default`, то выполняется `default`. Иногда можно сделать пустой `default`, чтобы получить неблокирующий вызов:

  ```go
  // some logic

  // вдруг контекст отменили. Проверим. Если всё хорошо, пойдём дальше.
  select {
  case <-ctx.Done():
  	return
  default:
  }

  // some logic
  ```

- если готовых нет и нет `default`’а, то select блокируется, пока какой-то из `case`’ов не станет готов

`case` считается готовым, когда:

- в буфере канала есть значение
- отправитель начал ждать ← для небуф.
- канал закрыт (чтение из такого не блокируется и всегда возвращает zero, false)

---

Чтение

- в буфере есть место на запись
- получатель начал ждать

---

Запись

Допустим, есть следующая задача: “Функция slowCall() проводит долгие вычисления. Нужно дождаться или пока она закончит, или пока не истечёт время таймаута.”

```go
ch := make(chan string)

go func() {
	time.Sleep(200 * time.Millisecond)
	ch <- "1st"
}()

select {
case v := <-ch:
	fmt.Println(v)
case <-time.After(100 * time.Millisecond): // СРАБОТАЕТ
	fmt.Println("timeout")
}
```

После срабатывания ветки таймаута никто не читает `ch`, поэтому горутина навсегда застрянет на `ch <- "1st"` — **утечка горутины**. Есть 2 способа устранить утечку, т.е. убедиться, что функция не заблокировалась навсегда:

```go
func slowCall(ch chan string) {
	time.Sleep(200 * time.Millisecond)
	ch <- "res" // сразу запишется, не будет ждать
}

ch := make(chan string, 1) // задали буфер размером 1
go slowCall(ch)

select {
case v := <-ch:
	fmt.Println(v)
case <-time.After(100 * time.Millisecond):
	fmt.Println("timeout")
}
```

> Если истёк таймаут, то вычисления продолжат выполняться. А нам ведь их результат уже и не нужен!

Поэтому можем добавить контекст и вынести запись в канал в отдельный `case` (или контекст отменён, или пишем в канал):

```go
func slowCall(ctx context.Context) (string, error) {
	select {
	case <-time.After(200 * time.Millisecond):
		return "res", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
```

```go
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Millisecond)
	defer cancel()

	ch := make(chan string)

	go func() {
		v, err := slowCall(ctx)
		if err != nil { return } // отменили, слать ничего не нужно

		select {
		case <-ctx.Done(): // отменили, слать ничего не нужно
		case ch <- v: // читатель уже ждёт
		}
	}()

	select {
	case v := <-ch:
		fmt.Println(v)
	case <-ctx.Done():
		fmt.Println("timeout:", ctx.Err())
	}
}
```

# nil-канал и выключение ветки

Допустим, нам нужно читать из двух каналов. Как только значения закончатся в одном и он закроется, то дочитать другой. Но чтение из закрытого канала никогда не блокируется: оно возвращает zero, false. Получается, когда канал закроется, select всегда будет выбирать ветку, связанную именно с ним.

Факт: чтение из nil-канала и запись в него блокируются навсегда. Значит, select никогда не выполнит эту ветку. Получается, если обнулить канал, то можно выключить соотв. ветку:

```go
a := gen(2)
b := gen(1)

for a != nil || b != nil { // at least 1 alive
select {
case v, ok := <-a:
	if !ok { // closed
		a = nil
		continue
	}

	fmt.Printf("got value %v from a.\n", v)
case v, ok := <-b:
	if !ok { // closed
		b = nil
		continue
	}

	fmt.Printf("got value %v from b.\n", v)
}
}
```

```go
func gen(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)

		for i := range n {
			ch <- i
		}
	}()

	return ch
}
```
