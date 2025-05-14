package main

import "fmt"

type Address struct {
	Kota, Alamat string
}

func (address *Address) ChangeAddress() {
	address.Kota = "Jakarta"
}

type Man struct {
	Name string
}

func (man *Man) GetMarried()  {
	man.Name = "Mr. " + man.Name
}

func main() {
	alamat1 := Address{"Medan", "Jalan kesana"}
	fmt.Println(alamat1)
	alamat1.ChangeAddress()
	fmt.Println(alamat1)

	habib := Man{"Habib"}
	habib.GetMarried()
	fmt.Println(habib)
}