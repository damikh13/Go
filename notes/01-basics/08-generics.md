Раньше, чтобы определить функцию для нескольких типов, приходилось делать несколько разных функций:

```go
func SumInt(nums []int) int {
	var total int
	for _, val := range nums {
		total += val
	}
	return total
}
```

```go
numsInt := []int{1, 2, 3, 4, 5}
fmt.Println("SumInt(numsInt) =", SumInt(numsInt)) // 15
```

```go
func SumFloat(nums []float64) float64 {
	var total float64
	for _, val := range nums {
		total += val
	}
	return total
}
```

```go
numsFloat := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
fmt.Println("SumFloat(numsFloat) =", SumFloat(numsFloat)) // 16.5
```

Начиная с версии 1.18, в Go появилась возможность писать общие функции:

```go
func SumNumsT int | float64 T {
	var total T
	for _, val := range nums {
		total += val
	}
	return total
}
```

```go
fmt.Println("SumNums(numsInt) =", SumNums(numsInt))
fmt.Println("SumNums(numsFloat) =", SumNums(numsFloat))

```

На самом деле, часть внутри квадратных скобок можно вынести в отдельный интерфейс:

```go
type Number interface {
	~int | int8 | int16 | ~float32 | ~float64
}
```

- или int8, или любой кастомный тип, содержащий в себе int (`type CustomInt int`), или int16, или …

```go
func SumT Number T { ... }
```

Пример реально использующегося интерфейса — `comparable`, в который входят примитивные типы данных, указатели, статические массивы но не входят динамические массивы, хэш-таблицы и функции.

Кроме обобщения функций, есть обобщение структур:

```go
type FlatGrid**[T any]** struct {
	Data       []**T**
	Rows, Cols int
}
func NewFlatGrid**[T any]**(rows, cols int) *FlatGrid**[T]** {
	return &FlatGrid**[T]**{
		Data: make([]**T**, rows*cols),
		Rows: rows,
		Cols: cols,
	}
}
func (fg *FlatGrid**[T]**) At(y, x int) **T** { return fg.Data[y * fg.Cols + x] }
func (fg *FlatGrid**[T]**) Set(y, x int, val **T**) { fg.Data[y * fg.Cols + x] = val }
func (fg *FlatGrid**[T]**) String() (output string) {
	for i := range fg.Rows {
		for j := range fg.Cols {
			output += fmt.Sprintf("[%v]", fg.At(i, j))
		}
		output += "\n"
	}
	return output
}

fg := NewFlatGrid(5, 3)
fmt.Println(fg)
```

Особое место — пустой интерфейс.
