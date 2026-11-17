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
