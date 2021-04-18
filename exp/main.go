package main

import (
	"html/template"
	"os"
)

type Dog struct {
	Name string
}
type User struct {
	Name string
	Age  int
	Dog  Dog
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name: "John Smith",
		Age:  45,
	}

	data.Dog.Name = "Morty"
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	data.Name = "Andrea C"
	data.Age = 2
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	data.Name = "Sarah E"
	data.Age = 34
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
