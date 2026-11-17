Можем создать проект на Go, прописав команду `go mod init api` в пустой папке. Она создаст файлик `go.mod`:

```go
module api

go 1.27.1
```

Далее можем создать несколько файлов:

```
~/Documents/Programming/Go/
|-- go.mod
|-- main.go <- package main
|-- handlers/
    |--user.go <- package handlers
    |--another.go <- package handlers
```

Чтобы использовать внутри main’а функции из user.go, нужно:

```go
package main

import "api/handlers"

func main() {
	handlers.SomeFunc()
}
```

Но для удалённой разработки и правильного подключения модулей (допустим, мы захотим использовать handlers в другом проекте) нужно указывать название модуля через проект на Github:

```go
module ~~api~~ github.com/yourname/api

go 1.27.1
```

```go
package main

import ~~"api/handlers"~~
import "github.com/yourname/api/handlers"

func main() {
	handlers.SomeFunc()
}
```

Далее его следует выложить на Github. И тогда другой человек (или мы сами) сможет его импортировать.

- Если не выкладывать и хочется импортировать локальный модуль в другой, то следует добавить в `go.mod`:

  ```go
  require github.com/yourname/api v0.0.0
  replace github.com/yourname/api => ../Go/api
  ```

  Или можно использовать **workspace**:

  ```go
  cd ~/Documents/Programming/Go
  go work init ./api ./other-module
  ```

Если мы пропишем `go get github.com/google/uuid[@v1.5.0]/[@latest]/[@commit-hash]`, то произойдут следующие вещи:

1. Код скачается из proxy.golang.org (если его ещё никогда никто не скачивал, то внутри proxy-сервера исходный код найдётся на Github)
2. Сохранится в $GOPATH/pkg/mod (обычно ~/go/pkg/mod).
3. В go.mod добавится строчка `require`.
4. Обновится файлик `go.sum`.

Далее его можно будет спокойно импортировать:

```go
import "github.com/google/uuid"
```

- `go mod tidy`: добавляет require, если мы что-то импортировали, но забыли сделать go get, и удаляет require того, что нигде не используется
- `go get -u ./...`: upgrade dependencies to newer minor/patch versions.
- `go list -m all`: show every module in your build.
- `go mod download`: pre-fetch everything into the cache.
- `go mod why <module>`: explain why a module is in your build.

Иногда вместо скачивания зависимостей из github’а используют папку `vendor/`.

Есть особый пакет — **internal**. Только тот модуль, в котором он и лежит, может импортировать из него. А другой модуль уже не сможет.
