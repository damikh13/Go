Базовый синтаксис таков:

```go
if A && !B { // !B начнёт вычисляться только в том случае, если A == true!
	...
} else if C {
	...
} else {
	...
}
```

`else` должен начинаться на той же строке, что и закрывающая скобка!

Фишка Go: если переменная используется только внутри блока if, можем проинициализировать её исключительно внутри него:

```go
if weight := 10; weight > 10 {
	fmt.Println("obese")
} else {
	fmt.Println("scrawny")
}

fmt.Println(weight) // ошибка! нет такой

```

В switch тоже есть локальная инициализация:

```go
switch x := 10; {
case x > 10:
	fmt.Println("big")
default:
	fmt.Println("small")
}
```

- switch без условия. Применяют, чтобы не писать много-много if else … { … } if else …

```go
switch day := "alskdjasld"; day {
case "Mon", "Tue", "Wed", "Thu", "Fri":
	fmt.Println("work day :(")
case "Sat", "Sun":
	fmt.Println("weekend day :)")
default:
	fmt.Println("unknown")
}
```
