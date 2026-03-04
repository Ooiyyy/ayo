package main

import (
	"fmt"
	// "rsc.io/quote"
)

func main() {

	//error handling
	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	// message, err := greetings.Hello("")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//error handling end

	var firstName string = "Roy"
	var lastName string
	lastName = "XPR"
	age := 100
	var value = (((2+6)%3)*4 - 2) / 3
	var hasil = (value != 2)
	nilai := 85
	point := 75

	// fmt.Println("Hello, World!")
	// fmt.Println(quote.Go())
	// message := greetings.Hello("Gladys")
	// fmt.Println(message)

	fmt.Printf("Hello, %s %s! Umurmu %d\n", firstName, lastName, age)
	fmt.Println("Coba Coba")
	fmt.Print("Coba Coba \n")
	fmt.Printf("Hasil dari 2 + 6 mod 3 * 4 - 2 / 3 adalah %d\n", value)
	fmt.Printf("Apakah hasilnya sama dengan 2? %t\n", hasil)
	if nilai >= 90 {
		fmt.Printf("Nilaimu %d Sangat Baik \n", nilai)
	} else if nilai >= 80 {
		fmt.Printf("Nilaimu %d Baik \n", nilai)
	} else if nilai >= 70 {
		fmt.Printf("Nilaimu %d Cukup \n", nilai)
	} else {
		fmt.Printf("Nilaimu %d Kurang \n", nilai)
	}

	if poinFinal := point / 10; poinFinal >= 10 {
		fmt.Printf("Poin finalmu adalah %d kamu dinyatakan lulus\n", poinFinal)
	} else if poinFinal >= 7 {
		fmt.Printf("Poin finalmu adalah %d kamu harus mengikuti remedi\n", poinFinal)
	} else {
		fmt.Printf("Poin finalmu adalah %d silahkan mencoba lagi lain kali \n", poinFinal)
	}
}
