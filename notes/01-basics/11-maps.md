Создать map можно несколькими способами, но лучше всего использовать make или создание через инстанс:

```go
mp1 := make(map[string]int, 100)
```

> cap = 100 необязателен:
>
> ```go
> mp2 := make(map[string]int)
> ```

```go
var mp3 map[string]int // == nil
```

> Можно читать, но нельзя писать!
>
> ```go
> mp3["any key"] // 0
> mp3["x"] = 10 // panic!
> ```

```go
mp4 := map[string]int{"John": 30, "Alice": 25}
```

Основные операции — чтение/запись, <u>**удаление**</u> и длина:

```go
ages["Alice"] = 10 // добавить или обновить
JohnAge := ages["John"] // прочитать
missingPersonsAge := ages["missing key"] // 0
delete(ages, "John")
fmt.Println(ages, len(ages)) // map[Alice:10] 1
```

Так как Go выдаёт дефолтное (нулевое) значение на отсутствующий ключ, то непонятно, как проверить, была ли вообще искомая запись. Для этого используется особый приём — **comma ok idiom**:

```go
val, ok := ages["some key"]
if !ok {
	fmt.Println("key was missing!")
}
```

Итерация по map’у будет выдавать значения в случайном порядке. Это умышленный дизайн, призванный напомнить разработчикам, что на порядок в этой структуре данных рассчитывать не стоит.

```go
for k, v := range mp {
	fmt.Printf("mp[%q] = %v\n", k, v)
}
/* в этот раз:
	mp["3st"] = 30
	mp["1st"] = 10
	mp["2st"] = 20
	mp["4st"] = 40
*/
```

Если всё-таки очень хочется вывести в нужном порядке, то можно достать ключи, отсортировать их и пройтись по этому отсортированному слайсу, выводя соответствующие значения для ключей:

```go
	mp5 := map[string]int{"1st": 10, "2nd": 20, "3rd": 30, "4th": 40}
	mp5Keys := make([]string, **0**, len(mp5))
	for k := range mp5 {
		mp5Keys = append(mp5Keys, k)
	}
	fmt.Println(mp5Keys)  // 4th 1st 3rd 2nd
	sort.Strings(mp5Keys) // 1st 2nd 3rd 4th
	for _, v := range mp5Keys {
		fmt.Printf("mp5[%q] = %v\n", v, mp5[v])
	}
```

Мы всегда должны уметь сравнить ключи (==):

**Можно** выбрать в качестве ключей:

- строки, числа, булевые значения
- указатели
- интерфейсы и структуры (если в них содержатся значения, которые сами можно сравнить)
- статические массивы

**Нельзя**:

- функции
- slice’ы
- map’ы

Чтобы не использовать map map’ов для составного ключа (`map[”first name”][”second name”]`), так как с ними возникают некоторые трудности (нужно проинициализировать каждый внутренний map, чтобы он перестал быть = nil, иначе код при попытке получить значение по ключу от nil’а будет выдавать panic), придумали класть структурку:

```go
type CompositeKey struct {
	Path, Country string
}
mp := make(map[CompositeKey]int)
```

Map’ы передаются по указателям:

```go
func addNew(mp map[string]int) {
	mp["new"] = 1000
}

fmt.Println("mp5:", mp5) // mp5: map[1st:10 2nd:20 3rd:30 4th:40]
addNew(mp5)
fmt.Println("mp5:", mp5) // mp5: map[1st:10 2nd:20 3rd:30 4th:40 new:1000]
```

Если мы храним структурки, то нельзя просто так изменить значение: нужно или копировать, менять копию и записывать её обратно, или работать со словарём указателей:

```go
personsInfo := map[string]Person{"1st": {"John", 30}}
personsInfo["1st"].Age = 31 // нельзя!

// нужно так:
p1 := personsInfo["1st"]
p1.Age = 31
personsInfo["1st"] = p1
```

```go
personsInfoPointers := map[string]*Person{"1st": {"John", 30}}
personsInfoPointer["1st"].Age = 31 // можно!
```

Прикольный факт: необязательно указывать `map[string]Person{"1st": <u>**Person**</u>{"John", 30}}`. Go сам понимает, что, раз после [string] указан тип Person, то его и нужно ожидать после двоеточия.

Map’ы плохо подходят для многопоточности. Если нужно и писать, и читать, то нужно или использовать mutex, или особый Map из пакета `sync`:

```go
var mu mutex.Mutex
mu.Lock()
mp["x"] = 1
mu.Unlock()
```

```go
var sm sync.Map
sm.Store("x", 1)
v, ok := sm.Load("x")
```

Так как в Go нет встроенного Set’а, то Map с пустой структурой выполняет его функциональность:

```go
set := make(map["string"]struct{})
set["first item"] = struct{}{}
_, ok := set["second item"]
if !ok {
	fmt.Println("second item's not in the map")
}
```

Есть несколько стандартных методов пакета `maps`:

- `maps.Keys(m)` ← возвращает итератор! Или проходимся по нему через range, или `keys := slices.Collect(maps.Keys(mp))`
- `maps.Values(m)`
- `maps.Clone(m)`
- `maps.Equal(m1, m2)`
- `maps.DeleteFunc(m, func(k string, v int) bool { return v < 0 })`
