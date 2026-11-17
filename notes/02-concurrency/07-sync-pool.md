sync.Pool — набор временных **переиспользуемых** объектов. Применяется, когда мы не хотим каждый раз создавать объект заново.

- Get() — если объект есть, то он выдаётся, иначе — создаётся при помощи New()
- Put() — положить объект обратно

```go
pool := sync.Pool{
	New: func() any {
		fmt.Println("allocating new []byte slice")
		return make([]byte, 1024)
	}
}

obj := pool.Get().([]byte)
fmt.Printf("got object from pool of length: %v\n", len(obj))

pool.Put(obj)

reusedObj := pool.Get().([]byte)
fmt.Printf("got reused object from the pool of length: %v\n", len(reusedObj))

pool.Put(reusedObj)
```
