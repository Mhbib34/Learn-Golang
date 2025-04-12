package main

import "fmt"

var hargaPerJam int = 100000

func hitungDiskon(jumlahDiskon,waktuMain int)int  {
	totalHarga := hargaPerJam * waktuMain
	hargaDiskon :=  float64(totalHarga) * (float64(jumlahDiskon) / 100)
	hargaFix := float64(totalHarga) - hargaDiskon
	return int(hargaFix)
}

func hitungHarga(waktuMain int) int {
	return hargaPerJam * waktuMain
}

func bayar(uang,totalHarga int) int {
	if totalHarga > uang {
		fmt.Println("Uangnya Kurang")
	}
	kembalian := uang - totalHarga
	return kembalian
}

func main()  {
	var userInput int

	for  {
			fmt.Println("---Kasir Futsal---")
			fmt.Println("1. Hitung Total")
			fmt.Println("2. Bayar")
			fmt.Println("3. Keluar")
			fmt.Println("Masukkan pilihan Anda:")
			fmt.Scanln(&userInput)

			var jumlahUang,waktuMain,diskon,totalHarga int

			switch userInput {
			case 1:
					fmt.Println("Masukkan Waktu Bermain ")
					fmt.Scanln(&waktuMain)
					fmt.Println("Masukkan Jumlah Diskon")
					fmt.Scanln(&diskon)
					if diskon != 0{
						fmt.Println("Total bayar adalah : ",hitungDiskon(diskon,waktuMain))
					}else{
						fmt.Println("Total bayar adalah : ",hitungHarga(waktuMain))
					}
			case 2:
					fmt.Println("Masukkan Total Bayar ")
					fmt.Scanln(&totalHarga)
					fmt.Println("Masukkan Jumlah Uang ")
					fmt.Scanln(&jumlahUang)
					fmt.Println("Kembaliannya adalah : ", bayar(jumlahUang,totalHarga))
			case 3:
					fmt.Println("Keluar dari aplikasi.")
					return
			default:
					fmt.Println("Pilihan tidak valid.")
			}
	}
}