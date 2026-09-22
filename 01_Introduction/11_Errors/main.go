package main

import (
	"errors"
	"fmt"
)

func doSomething(fail bool) (res int, err error) {
	// some logic
	if fail {
		return 0, errors.New("doSomething failed")
	}
	return 42, nil
}

type NotFoundError struct {
	Name string
}

func (nfe *NotFoundError) Error() string {
	return fmt.Sprintf("not found: %v", nfe.Name)
}

func lookupUser(name string) error {
	return fmt.Errorf("lookup failed for:%w", &NotFoundError{Name: name})
}

func main() {
	res, err := doSomething(false)
	if err != nil {
		fmt.Println("error happened:", err)
		return
	}
	fmt.Println("res:", res)

	base_error := errors.New("base level error")
	fmt.Println("err1:", base_error)
	err2 := fmt.Errorf("wrapping base error: %w", base_error)
	fmt.Println("err2:", err2)

	if errors.Is(err2, base_error) {
		fmt.Println("err2 is actually a base_error")
	}

	err3 := lookupUser("Bob")
	fmt.Println("err3:", err3)

	var nfe *NotFoundError
	if errors.As(err3, &nfe) {
		fmt.Printf("err3 is actually a NotFoundError, the name was:%s\n", nfe.Name)
	}

	if nfErr, ok := errors.AsType[*NotFoundError](err3); ok {
		fmt.Printf("err3 is actually a NotFoundError, the name was:%s\n", nfErr.Name)
	}

	err3_insides := errors.Unwrap(err3)
	fmt.Println("err3_insides:", err3_insides)
}
