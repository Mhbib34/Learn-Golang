package main

import "fmt"

func main() {
	var values = []int{
		10, 20, 30,40,50,60,70,80,90,
	}

	for _,angka := range values {
		if angka == 20 {
			continue
		}else if angka == 70 {
			break
		}
		fmt.Println(angka)
	}

	person := map[string]string{
		"nama" : "Muhammad Habib",
		"alamat" : "Medan",
	}

	fmt.Println(person)
	delete(person,"nama")

 	baru := append(values, )

	panjangslice := cap(baru)

	fmt.Println(values)
	fmt.Println(baru)
	fmt.Println(panjangslice)
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])
	fmt.Println(person)
}