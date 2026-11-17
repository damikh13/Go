Основа такая:

```go
func sum(int a, int b) int {
	return a + b
}
```

Далее можем обозначить тип нескольких переменных всего один раз (если он у них один и тот же):

```go
func sub(a, b int) int {
	return a + b
}
```

Также Go нативно поддерживает несколько возвращаемых значений, указанных через запятую:

```go
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("can't divide by 0")
	}
	return x / y, nil
}
```

```go
div_res1, err1 := divide(5, 0)
div_res2, err2 := divide(8, 2)
fmt.Printf("div(%v, %v) = %v, %v\n", 5, 0, div_res1, err1) //
fmt.Printf("div(%v, %v) = %v, %v\n", 8, 2, div_res2, err2)
```

Есть интересная особенность: в маленьких функциях можно вообще не указывать, что мы возвращаем — **naked return**. Go сам подставит рядом с return то, что мы укажем возвращаемых значениях сигнатуры функции:

```go
func DoSmt(val int) (x, y int) { // не просто (int, int)
	x = val * 2
	y = val * 4
	return
}

DoSmt(10) // 20 40
```

Чтобы работать с переменным числом параметров, используется многоточие:

```go
func PrintValues(nums ...int) {
	for _, v := range nums {
		fmt.Println(v)
	}
}

PrintValues(1,2,3)
nums := []int{1, 2, 3}
Prinvalues(nums...)
```

В Go функции все функции 1-го класса. То есть с ними можно работать как с обычными переменными — передавать в качестве аргументов, присваивать переменным, возвращать и т.д:

```go
add := func(a, b int) int {
	return a + b
}
fmt.Println(add(5, 6)) // 11
```

```go
func(a, b int) {
	fmt.Println(a, b) // 5 6
}(5, 6)
```

**замыкания**

```go
func outer() func() int {
	count := 0 // внешняя переменная
	return func() int {
		count++
		return count
	}
}

out := outer()
fmt.Println(out())
fmt.Println(out())
```

Также можно закрепить определённый тип функций:

```go
type type_name func(a, b int) int
func f(another_func type_name, x int) { ... }
```

Чтобы определять метод у структуры, используется несколько иной синтаксис — перед именем добавляется указание структурки:

```go
type Person struct {
	name string
}

func **(p Person)** PrintName() {
	fmt.Prinln(p.name)
}
```

Большинство примитивных типов данных + кастомных структур передаются в функцию по значению, то есть просто копируются. Изменения внутри функции не приводят к изменению переданного в них объекта.

Чтобы изменения внутри функции всё-таки влияли на исходный объект, нужно передавать его по указателю:

```go
func changeValue(x *int) {
	x = 50
}

x := 10
changeValue(&x)
```

Для slice’ов/map’ов/каналов это правило не работает. В них функция копирует указатели на начало структурки данных. А так как скопировался указатель, то изменения будут проходить!

В Go есть возможность отложить несколько функций (причём исполняться они будут в обратном порядке (2, 1, close)):

```go
func readFile(filePath string) {
	f, _ := os.Open("file.txt")
	defer f.Close()
	defer func() {
		fmt.Println(1)
	}**()**
	defer func() {
		fmt.Println(2)
	}**()**
}
```

Значения переменных, которые будут использоваться внутри defer, определяются в момент определения самого defer (а не на этапе исполнения):

```go
x := 1
defer fmt.Println(x) // в конце исполнения функции напишет 1
x = 2
```

Через defer + recover используется обработка совсем лютых ошибок (panic). Если они случаются, то исполняется соответствующий блок (а выполнение всей функции прекращается):

```go
func f() err error {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()

	panic("smt broke!")

	// no further code will be executed
}
```

В функциях Go нет следующих стандартных для других языков фич:

1. `func add(a, b int = 0)` — дефолтные значения
2. перегрузок функций
3. `divide(a: 10, b: 2)` — передачи по имени
