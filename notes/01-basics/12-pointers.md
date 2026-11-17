Если у нас есть функция, которая работает со значением (например, strings.ReplaceAll(str, “x”, “y”)), то придётся раскрывать указатель:

```go
if pointerToMessage == nil {
	return errors.New("invalid input")
}
messageVal := *pointerToMessage
messageVal = strings.ReplaceAll(messageVal, "dang", "****")
messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
*pointerToMessage = messageVal
```

> У указателей может быть проблема: если попробовать достать значение нулевого, то программа запаникует.

Притом нам необязательно при создании объекта указывать `&` или `new`. Можем создавать их по-обычному, а Go будет сам подставлять `&` при обращении к методам и полям:

```go
type Person struct { Name string }
func (p *Person) changeName(newName string) { p.Name = newName }

p1 := Person{Name: "John"}
p1.changeName("Alice") // Go переведёт это в (&p1).changeName()
```

Стоит помнить о том, что при полиморфном создании использовании объектов нужно обращать внимание на то, как были определены методы:

```go
type MyError struct { msg string }
func (me *MyError) Error() string { return me.msg }

var my_err Error = MyError{"some error"} // нельзя!
var my_err Error = &MyError{"some error"} // только так!
```

Редко можно встретить двойные указатели:

```go
x := 10
p := &x
pp := &p
fmt.Println(**p) // 10
```
