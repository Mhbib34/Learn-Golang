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

func (pelanggan Pelanggan) sayHello()  {
	fmt.Println("Halo nama saya",pelanggan.Nama)
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

	dian := Pelanggan{
		Nama: "Dian",
		Alamat: "Medan",
		Umur: 29,
	}
	fmt.Println(dian)

	joko := Pelanggan {"Joko","Medan",20}
	fmt.Println(joko)

	chris := Pelanggan{Nama: "Chris"}
	chris.sayHello()
}