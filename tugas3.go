package main

import (
	"fmt"
	"strings"

	humanize "github.com/dustin/go-humanize"
)

var subMenu int

func main() {
	var idr, usd, euro float64
	var gbp int

	fmt.Println("=========================")
	fmt.Println("Sub Menu : ")
	fmt.Println("1. Soal1")
	fmt.Println("2. Soal2")
	fmt.Println("3. Soal3")
	fmt.Println("=========================")
	fmt.Print("Pilih menu angka diatas : ")
	fmt.Scan(&subMenu)
	fmt.Println("=========================")

	if subMenu == 1 {
		fmt.Print("Masukkan Jumlah Uang : Rp ")
		fmt.Scan(&idr)
		usd = idr / 14220.25
		fmt.Printf("IDR : %s,00 \n", FormatRupiah(idr))
		fmt.Printf("USD : $ %.2f \n", usd)
	} else if subMenu == 2 {
		fmt.Print("Masukkan Jumlah Uang : ")
		fmt.Scan(&idr)
		euro = idr / 16869.72
		fmt.Printf("IDR : %s,00 \n", FormatRupiah(idr))
		fmt.Printf("EURO : € %.2f \n", euro)
	} else {
		fmt.Print("Masukkan jumlah GB Pounds : ")
		fmt.Scan(&gbp)
		knut := gbp * 100
		sickle := knut / 29
		galleon := sickle / 17
		st := knut % galleon
		sisa := knut % sickle

		fmt.Println("Jumlah knut yang didapat = ", knut)
		fmt.Println("Hasil Penukaran Mendapatkan = ", galleon, "Galleon(s)")
		fmt.Println("Sisa ditukar menjadi = ", st, "Sickle(s)")
		fmt.Println("Keping knut yang tersisa = ", sisa, "Knut(s)")
	}
}

func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return "Rp " + stringValue
}
