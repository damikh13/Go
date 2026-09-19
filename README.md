# Go

# 1. Введение

Go быстро компилируется и исполняется. Проект на Java может компилитися целый час. Go — максимум пару секунд.

## 1.1. Переменные

Основные типы не составных переменных:

- bool
- string
- int, uint ← в зависимости от разрядности системы
- int8, int16, int32, int64
- uint8, …, uint64
- byte = uint8
- rune = int32
- float32, float64 (default)
- complex64, complex128 (10.9 + 13.4i)
- указатели ()
- uintptr

Объявление их происходит двумя (третий для констант) способами:

```go
username := "admin"
```

> 📌 **Note:** Возможно только внутри функций, например, внутри `func main() {...}`!

> 📌 **Note:** Если объявим при помощи этого оператора внутри другой области видимости, то создастся новая переменная!

```go
{
	x := 10
	if true {
		x := 20
		fmt.Println(x) // 20
	}
	fmt.Println(x) // 10
}
```

```go
var num int = 123123123
var name string // = "" for now
name = "John"
```

> 📌 **Note:** Есть множественное объявление!

```go
var (
  name string = "John"
  age  int    = 30
)
```

Не указываем ни тип, ни `:=`:

```go
const pi = 3.14
const full_name = first_name + " " + last_name
```

> 📌 **Note:** Значения всех констант вычисляются во время компиляции, а не выполнения!

Есть интересный механизм для констант — автоматически увеличивающаяся переменная:

```go
const (
	Monday = iota // 0
	Tuesday
	Wednesday
)
fmt.Println(Wednesday) // 2
```

> 📌 **Note:** Если просто объявить переменную и не задать ей значение сразу, то оно по умолчанию станет = 0.

Определить тип переменной можно при помощи:

```go
fmt.Printf("type(%v) = %T\n", val, val) // type(12.34) = float64
```

Go поддерживает базовую конвертацию типов:

```go
var f float64 = -3.9
var i int = int(f) // -3
var u uint = uint(f) // 0

fmt.Println(string(65))       // "A"
fmt.Println(strconv.Itoa(65)) // "65"

```

> 📌 **Note:** Однако она может вести себя по-разному в зависимости от того, что мы пытаемся сконвертировать — переменную (в runtime’е) или константу (на этапе компиляции):

```go
int(3.9) // ошибка! cannot convert 3.9 (untyped float constant) to type int
```

```go
fl := 3.9
int(fl) // 3
```

Важный нюанс касающийся именования идентификаторов: если назвать с большой буквы, то он будет доступен при экспорте модуля, если с маленькой — то недоступен! К примеру, поле класса нельзя будет посмотреть или в целом проицинилизировать эксземпляр (`p := user.Profile{Name: "John", age: 30}`). Частая проблема у новичков — пропадающие поля в JSON. Геттеры принято называть Name(), сеттеры — SetName().

## 1.2. fmt.Printf

Внутри fmt.Printf можем указать placeholder, а далее его заполнить. Если не хотим выводить, а хотим лишь преобразовать строку, то используем `Sptintf()`:

```go
s1 := fmt.Sprintf("I am %v year old!", 21)
fmt.Println(s1) // I am 21 year old!
fmt.Printf("The lawyer's name is %v", "Saul Goodman.\n")
fmt.Printf("The foolowing sentence is either %v or %v.\n", true, false)
```

Основные типы:

- %v — для любых значений
- %s — для строк
- %d — для целых чисел
- %f (%.3f) — для чисел с плавающей точкой
- %t — для bool
- %q — для “” в выводе (`fmt.Printf("He said: %q.\n", "quote-quote-quote")`)
- %c — вывести символ по его коду (`fmt.Printf("%c\n", 65)` → A)
- %x — hex-представления для каждого байта чисел/строк (`fmt.Printf("%x\n", 255)` → ff)
- %+v, %#v — при выводе добавить ещё и поля (`{X:3 Y:4}`, `main.Point{X:3, Y:4}`)
- %p — вывести указатель (указывает туда же, куда и &int, для примитивных типов данных, но полезно для сложных структур):

  ```go
  	pointer_to_pt := &pt
  	fmt.Printf("%v\n", pointer_to_pt) // &{3 4}
  	fmt.Printf("%p\n", pointer_to_pt) // 0x605168c7c110
  ```

## 1.3. if, switch

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

Есть switch’и. В них тоже есть локальная инициализация:

```go
switch x := 10; {
case x > 10:
	fmt.Println("big")
default:
	fmt.Println("small")
}
```

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

## 1.4. for

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

Нет цикла while. Вместо этого можем скипать части `for`’а:

```go
i := 0
for i < 10 {
	...
	i++
}
```

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
for i, c := range str {
	fmt.Printf("i = %v, c = %v (%c).\n", i, c, c)
}
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

## 1.5. func

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
