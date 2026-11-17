Тесты в Go: оканчиваются на _test.go. Они не компилируются в бинарник на входе `go build`.

Unit-тест:

```go
package main

import "testing"

func TestGreet(t *testing.T) {
	got := greet()
	want := "Hello, World!"
	if got != want {
		t.Errorf("grret() = %q, want = %q", got, want)
	}
}
```

- t.Errorf помечает тест упавшим

Запуск: go test ./...

Пример теста из домашки:

```go
package main

import "testing"

func TestGreeting(t *testing.T) {
	for _, tc := range []struct { // говорим, какие поля должны быть у анонимной структуры
		name string
		in   string
		want string
	}{ // инициализируем её
		{name: "empty name", in: "", want: "Hello, World!"},
		{name: "plain name", in: "Alice", want: "Hello, Alice!"},
		{name: "unicode spaces trimmed", in: " \tАнна ", want: "Hello, Анна!"},
		{name: "only whitespace", in: " \n", want: "Hello, World!"},
		{name: "inner spaces kept", in: " A  B ", want: "Hello, A  B!"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := greet(tc.in)

			if got != tc.want {
				t.Errorf("greet(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

```
