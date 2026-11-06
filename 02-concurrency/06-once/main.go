package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	onceIncrementExample()
	onceIncrementDecrementExample()
	// onceDeadLockExample() // deadlock
	onceFuncExample()
	onceValueExample()
	onceValuesExample()
	oncePanicExample()
	onceNewPanicExample()
}

func onceIncrementExample() {
	var wg sync.WaitGroup
	var once sync.Once
	var counter int

	for i := range 100 {
		wg.Go(func() {
			once.Do(func() {
				counter++
				fmt.Printf("counter was increased by %v-th goroutine!\n", i)
			})
		})
	}

	wg.Wait()
	fmt.Printf("counter = %v.\n", counter)
}

func onceIncrementDecrementExample() {
	var once sync.Once
	var counter int

	increment := func() {
		counter++
	}
	decrement := func() {
		counter--
	}

	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("counter = %v.\n", counter)
}

func onceDeadLockExample() {
	var once sync.Once
	once.Do(func() {
		once.Do(func() {
			fmt.Println("here")
		})
	})
}

func onceFuncExample() {
	setup := sync.OnceFunc(func() {
		fmt.Println("setting up...")
	})

	setup()
	setup()
	setup()
}

func onceValueExample() {
	getNumber := sync.OnceValue(func() int {
		fmt.Println("calculating...")
		time.Sleep(1 * time.Second)
		return 42
	})

	fmt.Println(getNumber())
	fmt.Println(getNumber())
	fmt.Println(getNumber())
}

func onceValuesExample() {
	type config struct {
		highlight bool `json:"highlight"`
	}
	readConfig := sync.OnceValues(func() (*config, error) {
		configPath := "./02-concurrency/06-once/config.json"
		data, err := os.ReadFile(configPath)
		if err != nil {
			return nil, err
		}

		var c config
		return &c, json.Unmarshal(data, &c)
	})

	cfg, err := readConfig()
	cfg, err = readConfig()

	fmt.Printf("cfg: %v, err: %q\n", cfg, err)
}

func oncePanicExample() {
	var once sync.Once

	try := func() {
		defer func() {
			// recover()
			if r := recover(); r != nil {
				fmt.Printf("caught: %v.\n", r)
			}
		}()

		once.Do(func() {
			panic("boom")
		})
	}

	try() // caught: boom
	try() // nothing
	try() // nothing
}

func onceNewPanicExample() {
	get := sync.OnceValue(func() int {
		fmt.Println("here")
		panic("boom")
	})

	try := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("caught: %v.\n", r)
			}
		}()

		get()
	}

	try() // caught: boom
	try() // caught: boom
	try() // caught: boom
}
