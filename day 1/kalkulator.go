package main

import "fmt"

func inputPengguna() (int, int) {
	var a, b int
	fmt.Println("Masukkan angka pertama:")
	fmt.Scan(&a)
	fmt.Println("Masukkan angka kedua:")
	fmt.Scan(&b)
	return a, b
}

func Tambah(a, b int) int {
	return a + b
}

func Kurang(a, b int) int {
	return a - b
}

func Kali(a, b int) int {
	return a * b
}

func Bagi(a, b int) int {
	if b == 0 {
		fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan!")
		return 0
	}
	return a / b
}

func Modulus(a, b int) int {
	if b == 0 {
		fmt.Println("Error: Modulus dengan nol tidak diperbolehkan!")
		return 0
	}
	return a % b
}

func main() {
	var userInput int
	fmt.Println("---WELCOME---")
	fmt.Println("1. Tambah")
	fmt.Println("2. Kurang")
	fmt.Println("3. Kali")
	fmt.Println("4. Bagi")
	fmt.Println("5. Modulus")
	fmt.Println("Masukkan pilihan Anda:")
	fmt.Scan(&userInput)
	

	switch userInput {
	case 1:
		a, b := inputPengguna() 
		fmt.Println("Hasilnya adalah", Tambah(a, b))
	case 2:
		a, b := inputPengguna() 
		fmt.Println("Hasilnya adalah", Kurang(a, b))
	case 3:
		a, b := inputPengguna() 
		fmt.Println("Hasilnya adalah", Kali(a, b))
	case 4:
		a, b := inputPengguna() 
		fmt.Println("Hasilnya adalah", Bagi(a, b))
	case 5:
		a, b := inputPengguna() 
		fmt.Println("Hasilnya adalah", Modulus(a, b))
	default:
		fmt.Println("Anda memasukkan angka yang salah")
	}
}
