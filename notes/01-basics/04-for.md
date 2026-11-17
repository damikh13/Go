Обычное объявление цикла:

```go
for i := 0; i < 10; i++ {
	fmt.Println(i) // 0 1 2 ... 9
}
```

```go
for i := range 10 {
	fmt.Println(i) // 0 1 2 ... 9
}
```

Можем скипнуть любую часть:

```go
i := 0
for ; i < 10; i++ { ... }
fmt.Println(i)
```

```go
for i := 0; ; i++ {
	if (cond) break
}
```

```go
i := 0
for i < 10 {
	...
	i++
}
```

Получили собственный аналог while! Самого while в Go нет, и таким образом достигается его функциональность.

Можем пропустить вообще все части, получив аналог while true:

```go
cond := true
for { // while true
	...
	if cond {
		break
	}
}

```

Есть встроенная итерация по различным структуркам. Важно помнить, что итератор и значение — это копии, а не ссылки. Поменяв их внутри цикла, не получим изменения оригинального объекта:

```go
nums := []int{1,2,3,4}
for i, v := range nums {
	fmt.Println("i =", i, "v =", v)
	v = 100; // бесполезно!
}
```

```go
mp := map[string][int]{"a": 1, "b": 2}
for k, v := range mp {
	fmt.Printf("k = %q, v = %v", k, v)
}
```

```go
str := "hello"
for idx, r := range str {
	fmt.Printf("i = %v, c = %v (%c).\n", i, r, r)
}
```

Если скипаем последний итератор, то можем вовсе его не писать:

```go
for i, _ := range nums // можно
for i := range nums // но лучше
```

Если попробуем пройтись по динамическому массиву (slice’у) через range и внутри цикла увеличить его длину, итерироваться будем по до конца старой.

Можем прерывать внешний цикл, используя label’ы:

```go
outer:
for i := range 5 {
	for j := range 5 {
		if cond {
			continue outer // skip to next i, not j!
		}
	}
}
```
