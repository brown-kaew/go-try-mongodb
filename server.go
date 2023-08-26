package main

import (
	"fmt"

	"github.com/brown-kaew/go-try-mongodb/user"
)

func main() {
	fmt.Println("hello world")
	_, closeUserDb := user.TryAddAndFindData()
	defer closeUserDb()
}
