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
		Duration:  5,
		IsMember:  false,
		IsHolyday: false,
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
	tarifHoliday := 3000.0
	total := tarifDasar

	if car.Duration > 2 {
		total += (float64(car.Duration) - 2) * 2000
	}

	if car.IsMember && car.Duration <= 5 {
		total = total * 50 / 100
	} else if car.IsMember {
		total = total * 30 / 100
	}

	if car.IsHolyday {
		total += tarifHoliday
	}

	fmt.Println("--------System Tarif Parkir--------")
	fmt.Println("Plat Mobil\t:", car.CarId)
	fmt.Println("Lama Parkir\t:", car.Duration, "Jam")
	fmt.Println("Member Status\t:", car.IsMember)
	fmt.Println("Hari Libur\t:", car.IsHolyday)
	fmt.Printf("Total Tarif\t: Rp. %.2f\n", total)
	fmt.Println("___________________________________")
}
