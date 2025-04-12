package main

import "fmt"

type Pelanggan struct {
	Nama,Alamat string
	Umur int
}

type Barang struct {
	NamaBarang,KodeBarang string
	HargaBarang int
}

func main()  {
	var habib Pelanggan
	var rinso Barang

	rinso.NamaBarang = "Rinso"
	rinso.KodeBarang = "BG001"
	rinso.HargaBarang = 5000

	habib.Nama = "Muhammad Habib"
	habib.Alamat = "Medan"
	habib.Umur = 20

	fmt.Println(habib)
	fmt.Println(rinso)
}