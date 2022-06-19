package main

import (
	"errors"
	"fmt"
	"log"

	// "storage/internal/storage"
	"github.com/ftikhonin/storage/internal/storage"
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
	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("it works", file)

	RestoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it is restored", RestoredFile)
	fmt.Println("file content: ", string(file.Data))

}
