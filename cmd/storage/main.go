package main

import (
	"errors"
	"fmt"

	//"github.com/ftikhonin/storage/internal/storage"
	"storage/internal/storage"
)

func Hello() (string, bool) {
	return "hello", true
}

func main() {
	// Just move it one line above: don't use a short-if
	value, ok := Hello()
	if ok {
		fmt.Println(value)
	}

	e := errors.New("NewError")
	if e != nil {
		fmt.Println(e)
	}
	st := storage.NewStorage()
	fmt.Println("it works", st)
}
