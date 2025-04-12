package main

import (
	"fmt"
)

var todoList = []string{}

func input() string {
	var input string
	fmt.Println("Masukkan tugas: ")
	fmt.Scanln(&input)
	return input
}

func tampilkanTugas()  {
	fmt.Println("Daftar Tugas:")
	if len(todoList) == 0 {
		fmt.Println("Tugas Masih Kosong")
	}
    for i, tugas := range todoList {
    fmt.Printf("%d. %s\n", i+1, tugas)
	}
}

func tambahTugas(todo string) {
	todoList = append(todoList, todo)
	fmt.Println("Tugas ditambahkan:", todo)
}

func hapusTugas(index int) {
	if index < 0 || index >= len(todoList) {
			fmt.Println("Indeks tugas tidak valid.")
			return
	}
	todoList = append(todoList[:index], todoList[index+1:]...)
	fmt.Println("Tugas dihapus.")
}
func main() {
	var userinput int
	for {
			fmt.Println("---TO-DO LIST---")
			fmt.Println("1. Tambah Tugas")
			fmt.Println("2. Tampilkan Tugas")
			fmt.Println("3. Hapus Tugas")
			fmt.Println("4. Keluar")
			fmt.Println("Masukkan pilihan Anda:")

			fmt.Scanln(&userinput)

			switch userinput {
			case 1:
					tugas := input()
					tambahTugas(tugas)
			case 2:
					tampilkanTugas()
			case 3:
				tampilkanTugas() 
				var index int
				fmt.Println("Masukkan nomor tugas yang akan dihapus:")
				fmt.Scanln(&index)
				hapusTugas(index - 1)
			case 4:
					fmt.Println("Keluar dari aplikasi.")
					return
			default:
					fmt.Println("Pilihan tidak valid.")
			}
	}
}