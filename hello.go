package main

import (
	"fmt"
	// "time"
	"strings"
)

// "rsc.io/quote"

func main() {

	//error handling
	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	// message, err := greetings.Hello("")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//error handling end

	// var firstName string = "Roy"
	// var lastName string
	// lastName = "XPR"
	// age := 100
	// var value = (((2+6)%3)*4 - 2) / 3
	// var hasil = (value != 2)
	// nilai := 85
	// point := 75

	// fmt.Println("Hello, World!")
	// fmt.Println(quote.Go())
	// message := greetings.Hello("Gladys")
	// fmt.Println(message)

	// fmt.Printf("Hello, %s %s! Umurmu %d\n", firstName, lastName, age)
	// fmt.Println("Coba Coba")
	// fmt.Print("Coba Coba \n")
	// fmt.Printf("Hasil dari 2 + 6 mod 3 * 4 - 2 / 3 adalah %d\n", value)
	// fmt.Printf("Apakah hasilnya sama dengan 2? %t\n", hasil)
	// if nilai >= 90 {
	// 	fmt.Printf("Nilaimu %d Sangat Baik \n", nilai)
	// } else if nilai >= 80 {
	// 	fmt.Printf("Nilaimu %d Baik \n", nilai)
	// } else if nilai >= 70 {
	// 	fmt.Printf("Nilaimu %d Cukup \n", nilai)
	// } else {
	// 	fmt.Printf("Nilaimu %d Kurang \n", nilai)
	// }

	// if poinFinal := point / 10; poinFinal >= 10 {
	// 	fmt.Printf("Poin finalmu adalah %d kamu dinyatakan lulus\n", poinFinal)
	// } else if poinFinal >= 7 {
	// 	fmt.Printf("Poin finalmu adalah %d kamu harus mengikuti remedi\n", poinFinal)
	// } else {
	// 	fmt.Printf("Poin finalmu adalah %d silahkan mencoba lagi lain kali \n", poinFinal)
	// }

	// var ys = [5]int{10, 20, 30, 40, 50} // array
	// for _, v := range ys {
	// 	fmt.Println("Value=", v)
	// }

	// var zs = ys[0:2] // slice
	// for _, v := range zs {
	// 	fmt.Println("Value=", v)
	// }
	// var fruits = []string{"apple", "grape", "banana", "melon", "orange"}
	// var bFruits = fruits[1:5:5]

	// fmt.Println(bFruits)
	// fmt.Println(len(bFruits))
	// fmt.Println(cap(bFruits))

	//buat variabel map
	// var data map[string]int
	// // buat default value map
	// data = map[string]int{}
	// data["one"] = 1
	// fmt.Println(data["one"])
	// fmt.Println(data["two"])

	// hitung umurr (belajar function)
	// var tahunLahir int
	// tahunSekarang := time.Now().Year()
	// fmt.Println("Masukkan tahun lahir :")
	// fmt.Scanln(&tahunLahir)
	// if tahunLahir > tahunSekarang {
	// 	fmt.Println("Tahun lahir tidak boleh melebihi tahun sekarang, silahkan input ulang tahun lahir anda")
	// 	fmt.Scanln(&tahunLahir)
	// }
	// umur := hitungUmur(tahunLahir, tahunSekarang)
	// fmt.Printf("Sekarang Umurmu %d tahun", umur)

	//belajarvariadic function
	var hobbies = []string{"coding", "gaming", "traveling"}
	yourHobbies("Roy", hobbies...)
}

func selisih(nilaiAwal, nilaiAkhir int) int {
	return nilaiAkhir - nilaiAwal
}

func hitungUmur(tahunLahir, tahunSekarang int) int {
	return tahunSekarang - tahunLahir
}

// variadic function
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}
