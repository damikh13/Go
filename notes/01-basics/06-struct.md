Структурка определяется таким образом:

```go
type Person struct {
	Name string
	Age int
	IsAdmin bool
}
```

В Go нет `public`/`private`. Если назвать поле с большой буквы, оно будет доступно для других пакетов, которые будут импортировать наш. В рамках самого go-файлика (package x) поле всегда будет видно.

И далее есть 3 способа проинициализировать объект:

```go
p1 := Person{}
p2 := Person{"John", 30}
```

```go
p3 := new(Person) // type = *Person
p3.Age = 35
```

```go
var p4 Person
p4.Age = 25
p5.IsAdmin = false
```

Также есть несколько способов сделать функцию, которая будет взаимодействовать со структуркой:

Обычные функции:

```go
func printAge(p Person) {
	fmt.Println(p.Age)
}
```

```go
func changeAge(p *Person, num int) {
	p.Age += num
}
```

Методы:

```go
func (p Person) PrintAge() {
	fmt.Println(p.Age)
}
```

```go
func (p *Person) ChangeAge(num int) {
	p.Age += num
}
```

- когда структура слишком большая (и её дорого копировать)
- когда хотим менять поля
- если хоть один из методов принимает по указателю (то все остальные лучше тоже сделать такими)

Прикольный факт: методы можно определять для любого кастомного типа данных. Но нельзя для стандартных. Поэтому можем сделать так:

```go
type Celsius float64
func (c Celsius) ToFarenheit() float64 {
	return float64(c) * 9/5 + 32
}

с := Celsius(5.6)
fmt.Printf("%v\n", c) // 5.6
fmt.Println(c.ToFarenheit()) //

```

Так как в Go нет наследования, активно используются композиция:

```go
type Person struct {
	Name string
	Age int
	IsAdmin bool
}
type Address struct {
	City, State string
}
type Employee struct {
	Person
	Address
	Role string
}

e := Employee{
	Person{"John", 30, true},
	Address{"Miami", "Florida"},
	"Admin"
}
```

Внутренняя структура вообще не имеет понятия о том, что она является частью чего-то!

!git_remote-3.svg

В зависимости от того, указываем ли мы имя внутренней коробке или нет, будет разное поведение!

```go
type Base struct {
	field1 string
}

type Composite struct {
	Base
}

cmp := Composite{}
```

```go
cmp.field1 = "123"
```

- свойства как бы перетекли к внешней, в JSON’е они будут у объекта Composite

```go
type Base struct {
	field1 string
}

type Composite struct {
	**Base** Base
}
```

```go
cmp.**Base**.field1 = "123"
```

- свойства остались внутри внутренней, в JSON’е они будут обёрнуты в свой объект

У полей структур могут быть теги. К примеру, можем сказать, как следует методу encoding/json.Marshal(struct) записать каждое поле:

```go
struct Person {
	Name string     `json:"name"`
	Age int         `json:"age,omitempty"` // if = 0, will be excluded
	Email string    `json:"email"`
	Password string `json:"-"` // will be excluded
}
```
