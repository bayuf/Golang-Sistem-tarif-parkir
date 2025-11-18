package main

import "fmt"

type Parkir struct {
	CarId     string //Id Mobil
	Duration  int    // jam / waktu
	IsMember  bool   // status member
	IsHolyday bool   //hari libur / kerja
}

func main() {
	// Mobil 1 parkir 4jam, pada hari kerja dan bukan member
	car1 := Parkir{
		CarId:     "W 1090 T",
		Duration:  4,
		IsMember:  false,
		IsHolyday: true,
	}
	// Mobil 2 parkir 2jam, pada hari libur dan seorang member
	car2 := Parkir{
		CarId:     "AG 2025 G",
		Duration:  2,
		IsMember:  true,
		IsHolyday: true,
	}

	// Hitung Biaya Parkir
	hitungBiaya(car1)
	hitungBiaya(car2)

}

func hitungBiaya(car Parkir) {
	// Deklarasi Tarif
	tarifDasar := 5000.0
	tarifHolyday := 3000.0
	total := 0.0

	// Menghitung biaya perjam
	for jam := 1; jam <= car.Duration; jam++ {

		// jika melebihi 2jam maka biaya dasar menjadi 2000/jam
		if jam >= 3 {
			tarifDasar = 2000
		}

		// tambah biaya dasar ke total
		total += tarifDasar

		//jika member maka hitung diskon
		if car.IsMember && jam == car.Duration {
			if car.Duration <= 5 {
				total = total * 50 / 100
			} else {
				total = total * 30 / 100
			}
		}

		// jika hari libur maka ada biaya tambahan
		if car.IsHolyday && jam == car.Duration {
			total += tarifHolyday
		}
	}

	fmt.Println("--------System Tarif Parkir--------")
	fmt.Println("Plat Mobil\t:", car.CarId)
	fmt.Println("Lama Parkir\t:", car.Duration, "Jam")
	fmt.Println("Member Status\t:", car.IsMember)
	fmt.Println("Hari Libur\t:", car.IsHolyday)
	fmt.Printf("Total Tarif\t: Rp. %.2f\n", total)
	fmt.Println("___________________________________")
}
