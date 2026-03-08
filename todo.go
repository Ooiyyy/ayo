package main

import "fmt"

type Mahasiswa struct {
	Nama, Prodi, Alamat string
	Umur                int
}

func inputData() Mahasiswa {
	var m Mahasiswa
	fmt.Println("Masukkan Nama :")
	fmt.Scanln(&m.Nama)
	fmt.Println("Masukkan Prodi :")
	fmt.Scanln(&m.Prodi)
	fmt.Println("Masukkan Alamat :")
	fmt.Scanln(&m.Alamat)
	fmt.Println("Masukkan Umur :")
	fmt.Scanln(&m.Umur)

	return m
}

func tampilData(dataMahasiswa Mahasiswa) {
	fmt.Println("Nama :", dataMahasiswa.Nama)
	fmt.Println("Prodi :", dataMahasiswa.Prodi)
	fmt.Println("Alamat :", dataMahasiswa.Alamat)
	fmt.Println("Umur :", dataMahasiswa.Umur)
}

func main() {
	data := inputData()
	tampilData(data)
}
