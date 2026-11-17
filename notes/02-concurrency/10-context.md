Мы запустили горутину, которая делает что-то тяжёлое — ходит в БД или проводит сложные вычисления. Но потом выясняется, что результат её работы уже не нужен — пользователь закрыл вкладку. Что делать с такой горутиной? В Go нет подходящей конструкции, чтобы просто “убить” её. Поэтому делают так: горутина сама проверяет, просили ли её остановиться, и, если просили, завершается.

- `Done()` → канал, который закрывается, когда работу нужно прекращать
- `Err()` →
  - `nil`, если канал жив
  - `context.Canceled` == `errors.New("context canceled")`, если канал был отменён через cancel()
  - `context.DeadlineExceeded` == `deadlineExceededError{}`, если истекло время

Контексты образуют дерево. Корень этого дерева создаётся одним из двух способов:

1. от context.Background()
2. от context.TODO() ← когда нам в текущей функции нужен контекст, но мы ещё не успели написать код для его проброски
3. от `r.Context()` ← в `http.HandleFunc("/user", UserHandler)` + `func UserHandler(w http.ResponseWriter, r *http.Request) { ctx := WithTimeout(r.Context(), 3 * time.Millisecond) }`. Отменяется, когда клиент разорвал соединение/обработчик закончил работу.

Далее от корня порождаются производные контексты:

```go
ctx, cancel := context.WithCancel(parent)
ctx, cancel := context.WithTimeout(parent, 3*time.Second) // time.Duration
ctx, cancel := context.WithDeadline(parent, specificMoment) // time.Time (напр., 13:15)
ctx := context.WithValue(parent, key, value)
```

Каждый контекст знает только своего родителя. Отмена родителя ⇒ отмена всех потомков. Но не наоборот.

Пример:

```go
func query(ctx context.Context) (string, error) {
	select {
	case <-time.After(200 * time.Millisecond):
		return "rows", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50 * time.Millisecond)
	defer cancel()

	res, err := query(ctx)
	fmt.Printf("res: %q, err: %v\n", res, err)
	fmt.Println(errors.Is(err, context.DeadlineExceeded))
}
```

Важное правило: вызываем defer cancel() для того,

Правила использования:

1. Передавать первым аргументом функции (`func Foo(ctx context.Context, …)`).
2. Не хранить в полях структур ← контекст привязан к одному вызову/запросу, а не к объекту
3. Если ещё не понятно, какой именно контекст взять, пишем `context.TODO()`.
4. Пробрасывать дальше. Функция, получившая контекст, передаёт его дальше всем, кого вызывает.
5. `defer cancel()` ← чтобы в том случае, когда работа закончилась раньше таймера, ресурсы всё равно освободились.
6. Стараться избегать WithValue. Можно использовать только для тех данных, которые проходят через всю цепочку вызовов (request ID, trace ID). Передавая значения в контексте, избегаем засорения сигнатур: `func getUser(**ctx context.Context**, ~~requestID, traceID string~~, userID, id int) …`.

   ```go
   ctx = context.WithValue(ctx, key, value)
   v, ok := ctx.Value(key).(string)
   ```

   Если всё же используем, помним о следующей фишке: так как сравнение ключей производится не только по значению, но и по типу, то, чтобы в контексте строчки `"id"` (из пакета mypkg) и `"id"` (из пакета libpkg) были двумя разными ключами с двумя разными значениями, нужно определить для каждой из них кастомный тип:

   ```go
   ctx := context.Background()

   ctx = context.WithValue(ctx, "id", "req-42")
   ctx = context.WithValue(ctx, "id", "trace-7")

   fmt.Println(ctx.Value("id").(string)) // trace-7
   ```

   ```go
   type myKey // с маленькой буквы => неэкспортируемый

   ctx := context.Background()

   ctx = context.WithValue(ctx, myKey("id"), "req-42")
   ctx = context.WithValue(ctx, "id", "trace-7")

   fmt.Println(ctx.Value(myKey("id"))) // req-42
   fmt.Println(ctx.Value("id")) // trace-7
   ```
