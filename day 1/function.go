package main

import "fmt"

func sayHello()  {
	fmt.Println("Hallo ganteng")
}
func tambah(a int,b int) int {
	return a+b
}
func lingkaran () float32 {
	PI := float32(3.14)
	jariJari := 14

	luasLingkaran  := float32(PI * float32(jariJari) * float32(jariJari)) 
	return luasLingkaran
}

func sumAll(angka ...int) int {
	total := 0

	for _,number := range angka {
		total += number
		
	}
	return total
}

func main()  {
	sayHello()
	fmt.Println(tambah(10,20))
	fmt.Println(lingkaran())
	fmt.Println(sumAll(20,20,10,11,21))
}