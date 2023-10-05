/*Latihan soal camp 3:
1. Buatlah beberapa data menggunakan MAP
2. Buatlah beberapa struct dengan methodnya dibebaskan
3. Buatlah manipulasi data di struct dan pointer*/

package main

import "fmt"

func main() {
	println("**data menggunakan MAP**")
	Map()
	println("--------------------------------")
	println("**struct**")
	var tgl lahir
	tgl.bulan = "juni"
	tgl.tanggal = 30
	tgl.tahun = 2003

	fmt.Println("tanggal lahir: ", tgl.tanggal, ", ", tgl.bulan, tgl.tahun)
	println("--------------------------------")
	println("**Modifikasi Struct dan pointer**")
	// Membuat instansiasi struct Mahasiswa
	mahasiswa1 := Mahasiswa{
		Nama:     "Azhar",
		NIM:      "H1A021077",
		Semester: 3,
	}

	// Menampilkan data mahasiswa tanpa pointer
	fmt.Println("Data Mahasiswa Sebelum Manipulasi:")
	fmt.Println("Nama:", mahasiswa1.Nama)
	fmt.Println("NIM:", mahasiswa1.NIM)
	fmt.Println("Semester:", mahasiswa1.Semester)

	// Mendapatkan pointer ke struct Mahasiswa
	pointerMahasiswa := &mahasiswa1

	// Mengubah data menggunakan pointer
	pointerMahasiswa.Nama = "Azhar Fauzan"
	pointerMahasiswa.Semester = 5

	// Menampilkan data mahasiswa setelah manipulasi
	fmt.Println("\nData Mahasiswa Setelah Manipulasi:")
	fmt.Println("Nama:", mahasiswa1.Nama)
	fmt.Println("NIM:", mahasiswa1.NIM)
	fmt.Println("Semester:", mahasiswa1.Semester)
}

//No 1
func Map() {
	fmt.Println("Peggunaan Map")
	hari := map[string]int{
		"senin":  1,
		"selasa": 2,
		"rabu":   3,
		"kamis":  4,
		"jumat":  5,
		"sabtu":  6,
		"minggu": 7,
	}
	for key, val := range hari {
		fmt.Println(key, "  \t:", val)
	}

}

//No 2
type lahir struct {
	bulan   string
	tanggal int
	tahun   int
}

//No 3
type Mahasiswa struct {
	Nama     string
	NIM      string
	Semester int
}
