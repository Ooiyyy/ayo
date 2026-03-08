package main

import "fmt"

type Siswa struct {
	ID            int
	Nama, Jurusan string
	Umur          int
}

var dataSiswa []Siswa

func tambahSiswa(data *[]Siswa) {
	var s Siswa
	fmt.Print("Masukkan ID: ")
	fmt.Scanln(&s.ID)
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&s.Nama)
	fmt.Print("Masukkan Jurusan: ")
	fmt.Scanln(&s.Jurusan)
	fmt.Print("Masukkan Umur: ")
	fmt.Scanln(&s.Umur)

	*data = append(*data, s)

	fmt.Println("Data berhasil ditambahkan")
}

func tampilSiswa(data []Siswa) {
	for _, s := range data {
		fmt.Println("ID :", s.ID)
		fmt.Println("Nama :", s.Nama)
		fmt.Println("Jurusan :", s.Jurusan)
		fmt.Println("Umur :", s.Umur)
	}
}

func cariSiswa(data []Siswa, id int) *Siswa {
	for i := range data {
		if data[i].ID == id {
			return &data[i]
		}
	}
	return nil
}

func updateSiswa(data []Siswa) {
	var id int

	fmt.Println("Masukkan ID :")
	fmt.Scanln(&id)

	s := cariSiswa(data, id)

	if s == nil {
		fmt.Println("Data tidak ditemukan")
		return
	}

	fmt.Println("Nama Baru :")
	fmt.Scanln(&s.Nama)
	fmt.Println("Jurusan Baru :")
	fmt.Scanln(&s.Nama)
	fmt.Println("Umur Baru :")
	fmt.Scanln(&s.Nama)

	fmt.Println("Data Siswa Berhasil diubah")
}

func hapusSiswa(data *[]Siswa) {

	var id int

	fmt.Print("Masukkan ID:")
	fmt.Scanln(&id)

	for i, s := range *data {
		if s.ID == id {
			*data = append((*data)[:i], (*data)[i+1:]...)

			fmt.Println("Data Berhasil Dihapus")
			return
		}
	}

	fmt.Println("Data Tidak ditemukan")
}

func main() {
	for {
		fmt.Println("\n1. Tambah Siswa")
		fmt.Println("2, Tampilkan Siswa")
		fmt.Println("3, Update Siswa")
		fmt.Println("4, Hapus Siswa")
		fmt.Println("5, Keluar")

		var pilih int
		fmt.Print("Pilih :")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			tambahSiswa(&dataSiswa)
		case 2:
			tampilSiswa(dataSiswa)
		case 3:
			updateSiswa(dataSiswa)
		case 4:
			hapusSiswa(&dataSiswa)
		case 5:
			return
		}
	}
}
