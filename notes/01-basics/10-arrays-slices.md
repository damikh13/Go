В Go есть статические массивы:

```go
var myArr1 [3]int = [3]int{1, 2, 3}
fmt.Println(myArr1) // [1 2 3]
```

```go
myArr2 := [3]int{4, 5, 6}
```

```go
myArr3 := [...]int{7, 8, 9}
```

Работа с ними всегда происходит по значению. Если присваивать (`newArr := oldArr`) или передавать в функцию, то элементы будут перекопироваться.

Их можно сравнивать между собой:

```go
[3]int{1,2,3} == [3]int{1,2,3} // true
```

Но чаще всего разработчики работают с динамическими массивами — slice’ами:

```go
type sliceHeader struct {
	ptr *T
	len int
	cap int
}
```

Создать их можно несколькими способами:

```go
var s1 []int
```

> Нулевой. len = 0, cap = 0.

```go
s2 := []int{1,2,3}
```

```go
s3 := make([]int, 5)
```

> len = 5. Этот параметр обязательно указывается! В результате получается массив из нулей длины 5.
>
> Capacity = 5!

```go
s4 := make([]int, 5, 10)
```

> len = 5, cap = 10

Помимо того, можем взять срез от другого слайса или массива:

```go
s5 := []int{0,1,2,3,4,5,6}
s6 := s5[1:4] // [1 2 3], len = 3, cap = 6
s6[0] = 1000 // изменит и оригинальный s5 (он станет [0 1000 2 3 ... 6])
s7 := append(s6, 42) // снова изменит оригинальный s5:
// s5 = [0 1000 2 3 42 5 6]
```

`capNew = capOld - low`

```go
s5 := arr[1:3]
```

Функция append принимает то, что мы хотим добавить, так:

- sl = append(sl, 1, 2, 3)
- sl = append(sl, []int{1,2,3}…) ← необходим оператор многоточия

Как можно заметить, следует быть аккуратным с созданием слайсов из других слайсов, так как это может привести к переписыванию памяти друг друга.

Чтобы этого избежать, можно:

1. Копировать слайсы:

   Современный подход:

   ```go
   b := slices.Clone(a[1:10])
   ```

   Старый подход:

   ```go
   src := []int{0, 1, 2, 3}
   dst := make([]int, len(src))
   copy(dst, src)
   ```

2. Указывать 3-й индекс при копировании — s6 := s5[1:4:5] — взять элементы с 1-го индекса (вкл) по 4-й (невкл.) и ограничить cap: 5 - 1 = 4. Если его не указать, то cap будет: cap(s5) - low(=1).

   Если так сделать, то append увидит, что места не хватает, и создаст новый массив.

   Но, если не делать append, то создания нового массива не произойдет и `sliceB[0] = 1000` изменит и `sliceA[0]`.

Не всякий “пустой” слайс сравнивается с nil’ом:

```go
var varSlice []int // == nil
bracketSlice := []int{} // != nil
makeSlice := make([]int, 0) // != nil
```

Чтобы создать многомерный массив, есть 2 способа:

```go
grid1 := make([][]int, rows)
for i := range grid1 {
	grid1[i] = make([]int, cols)
}
```

> Но такой массив будет разбросан по памяти. Первый кусочек — адреса с 1-го по 10-й, второй — с 50-го по 60-й. Это мешает локальности памяти и может работать медленнее.

```go
type FlatGrid[T any] struct {
	Data       []T
	Rows, Cols int
}
func NewGridT any *FlatGrid[T] {
	return &FlatGrid[T]{
		Data: make([]T, rows*cols),
		Rows: rows,
		Cols: cols,
	}
}
func (fg *FlatGrid[T]) At(i, j int) T     { return fg.Data[i*fg.Cols+j] }
func (fg *FlatGrid[T]) Set(i, j int, v T) { fg.Data[i*fg.Cols+j] = v }
func (fg *FlatGrid[T]) String() (output string) {
	for i := range fg.Rows {
		for j := range fg.Cols {
			output += fmt.Sprintf("[%v]", fg.At(i, j))
		}
		output += "\n"
	}
	return output
}
```

Помимо slices.Clone, есть и другие полезные методы:

- slices.Index(sl, “b”) → 1
- slices.Equal([]int{1, 2, 3}, []int{1, 2, 3})
- slices.contain(sl, val) → true/false
- slices.Sort(sl) → in-place sorting
- slices.Reverse(sl) → in-place
- slices.Max(sl)/Min(sl)
- slices.Delete(sl, from, to)
- slices.Insert(sl, at, vals…)
