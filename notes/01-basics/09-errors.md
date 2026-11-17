В Go удобная логика работа с ошибками:

```go
func main() {
	user, err := getUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	profile, err := getUserProfile(user.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
}
```

```jsx
try {
  const user = getUser();
  try {
    const profile = getUserProfile();
  } catch (err) {
    console.log(err);
  }
} catch (err) {
  console.log(err);
}
```

> _Получается витвистая структура._

Сам тип `error` — это интерфейс, содержащий всего один метод — `Error() string`.

Базовый сценарий работы с ними выглядит вот так:

```jsx

func doSomething(fail bool) (res int, err error) {
	// some logic
	if fail {
		return 0, errors.New("doSomething failed")
	}
	return 42, nil
}

```

```jsx
func main() {
	res, err := doSomething(false)
	if err != nil {
		fmt.Println("error happened:", err)
		return
	}
	fmt.Println("res:", res)
}
```

При создании ошибки можем пользоваться одним из двух способов:

1. `errors.New(”file not found”)`
2. `fmt.Errorf(”failed to open config: %w”, err)` — оборачиваем ошибки при помощи `%w`. Потом её можно будет развернуть.

Чтобы работать с ошибками, которые были обёрнуты в несколько других, используются функции:

1. Is — ошибки всегда стоит сравнивать именно так (или через `!= nil`), а не через `==`.

   ```jsx
   base_error := errors.New("base level error")
   err2 := fmt.Errorf("wrapping base error: %w", base_error)

   if errors.Is(err2, base_error) {
   	fmt.Println("err2 is actually a base_error")
   }
   ```

2. As/AsType

   ```go
   type NotFoundError struct {
   	Whom string
   }
   func (nfErr *NotFoundError) Error() {
   	return fmt.Errorf("not found:%s", nfErr.Whom)
   }

   func lookupUser(name string) err {
   	return fmt.Errorf("lookup failed:%w", &NotFoundError{Whom: name})
   }
   ...
   err3 := lookupUser("Bob")
   ```

   > Метод Error() вызовется во время любого формата объекта &NotFoundError{Whom: name}. То есть во время `fmt.Errorf("lookup failed:%w”)`.
   >
   > Он работает со значением по указателю, так как так принято. Две ошибки с одним и тем же содержанием будут различаться.

   Старый способ:

   ```go

   var nfErr *NotFoundError // = nil
   if errors.As(err3, &nfErr) {
   	fmt.Printf("err3 is actually a NotFoundError, the name was:%s\n", nfErr.Name) // можем обратиться к полю ошибки!
   }
   ```

   Новый способ:

   ```go
   if nfErr, ok := errors.AsType*NotFoundError; ok {
   	fmt.Printf("err3 is actually a NotFoundError, the name was:%s\n", nfErr.Name)
   }
   ```

   - не создаём переменную nfErr, она локальная и удалится после цикла

3. Unwrap:

   ```go
   err3_insides := errors.Unwrap(err3)
   fmt.Println("err3_insides:", err3_insides)
   ```

Есть определённый набор дефолтных ошибок — **sentinel (”страж”) errors**:

- `_, err := os.Open("missing.txt")` — может выдать `os.ErrNotExist`
- `_, err := os.Create("/root/locked")` — может выдать `os.ErrPermission` (алиас на `fs.ErrPermission`)
- `f.Close(); f.Close()` (повторное закрытие) — может выдать `os.ErrClosed` (алиас на `fs.ErrClosed`)
- `os.Mkdir("existing_dir", 0755)` — может выдать `os.ErrExist` (алиас на `fs.ErrExist`)
