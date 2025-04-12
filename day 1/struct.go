package main

import "fmt"

type Pelanggan struct {
	Nama,Alamat string
	Umur int
}

func main()  {
	var habib Pelanggan
	habib.Nama = "Muhammad Habib"
	habib.Alamat = "Medan"
	habib.Umur = 20

	fmt.Println(habib)
}