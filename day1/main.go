package main

import (
	"fmt"
	"hari_tiga/soal"
)

func main() {
	// var allStudents = soal.Student{
	// 	Name:  []string{"andre", "johan", "cliff"},
	// 	Score: []int{80, 29, 78},
	// }
	// Average, namamin, namamax := allStudents.Average()

	// //fmt.Println(Average, nama)
	// fmt.Println("Rata-rata adalah ", Average)
	// fmt.Println("nilai minimum adalah ", namamin)
	// fmt.Println("nilai Maximum Adalah ", namamax)

	var Result soal.Chiper
	Result = &soal.Siswa{
		Name:       "BAlDE",
		NameEncode: "",
		Score:      70,
	}
	fmt.Println("hasil endkrip ", Result.Encode())

	fmt.Println("hasil Dekrip ", Result.Decode())

}
