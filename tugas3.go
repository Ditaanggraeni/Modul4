package main

import (
	"fmt"
)

func main() {
	var uang, numMenu int

	fmt.Println("=========================")
	fmt.Print("Masukkan Jumlah Uang Rupiah : ")
	fmt.Scan(&uang)
	fmt.Println("=========================")
	fmt.Println("Sub Menu : ")
	fmt.Println("1. Soal1")
	fmt.Println("2. Soal2")
	fmt.Println("2. Soal3")
	fmt.Print("Pilih menu angka diatas : ")
	fmt.Scan(&numMenu)

}
