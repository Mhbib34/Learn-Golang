package main

import "fmt"

type HasName interface  {
	GetName() string
	GetAge() int
}

func sayHello(hasName HasName)()  {
	fmt.Println("Hello",hasName.GetName())
	fmt.Println("My Age Is",hasName.GetAge())
}

type Person struct {
	Name string
	Age int
}

func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetAge() int {
	return person.Age
}


func Nama() any  {
	return "Habib"
}
func main()  {
	Habib := Person{"Habib",20}
	Daniel := &Habib
	Daniel.Name = "Daniel"
	Daniel.Age = 11

	var nama any = Nama()
	sayHello(Habib)
	sayHello(Daniel)
	fmt.Println(nama)
}