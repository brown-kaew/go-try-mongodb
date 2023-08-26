package main

import (
	"fmt"

	"github.com/brown-kaew/go-try-mongodb/user"
)

func main() {
	fmt.Println("hello world")
	userDb := user.NewUserDB()
	defer userDb.Close()
	userDb.FindAll()
	userDb.FindById(2)
}
