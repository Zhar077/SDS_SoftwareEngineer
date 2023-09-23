/*
Soal:
1. Buatlah fungsi untuk :
a. menghitung luas persegi
b. menghitung luas segitiga
c. menghitung luas lingkaran
d. menghitung energi potensial, kinetik
e. menghitung frekuensi atau getaran
f. konversi untuk semua satuan suhu
*/
package main

import "fmt"

func main() {
	LuasPersegi()
	LuasSegitiga()
	LuasLingkaran()
	Energi()
	Frekuensi()
	suhu()
}

//real basic
//a. Fungsi Menghitung Luas Persegi
func LuasPersegi() {
	var (
		sisi float64
	)
	fmt.Print("Masukkan Sisi Persegi: ")
	fmt.Scanf("%f", &sisi)
	L := sisi * sisi //Rumus Luas Persegi
	fmt.Println(L)
}

//b. Fungsi luas segitiga
func LuasSegitiga() {
	var (
		alas, tinggi float64
	)
	fmt.Println("Masukkan alas segitiga: ")
	fmt.Scanf("%f", &alas)
	fmt.Println("Masukkan tinggi segitiga: ")
	fmt.Scanf("%f", &tinggi)
	L := 0.5 * alas * tinggi
	fmt.Print("Luas Segitiga yaitu= ")
	fmt.Println(L)
}

//c. Fungsi Luas Lingkaran
func LuasLingkaran() {
	var (
		JariJari float32
	)
	fmt.Print("Masukkan Jari jari lingkaran: ")
	fmt.Scanf("%f", &JariJari)
	L := 3.14 * JariJari * JariJari
	fmt.Print("Luas Lingkaran yaitu= ")
	fmt.Println(L)
}

//Fungsi Energi Potensial dan Kinetik
func Energi() {
	var massa, tinggi, kecepatan float64

	Ep := massa * 10 * tinggi
	Ek := 0.5 * massa * kecepatan * kecepatan
	Em := Ek + Ep

	fmt.Print("Energi potensial= ")
	fmt.Println(Ep)
	fmt.Print("Energi Kinetik= ")
	fmt.Println(Ek)
	fmt.Print("Energi Mekanik= ")
	fmt.Println(Em)
}

//Fungsi Frekuensi
func Frekuensi() {
	var n_getaran float32 = 50
	F := n_getaran / 1
	T := 1 / F
	fmt.Print("Frekuensi= ")
	fmt.Println(F)
	fmt.Print("Periode= ")
	fmt.Println(T)
}

//Fungsi Konfersi Suhu
func suhu() {
	var (
		c, r, f, k float64
	)
	fmt.Println("Masukkan suhu dalam satuan celcius :")
	fmt.Scanf("%f", &c)
	reamur := (c * 4) / 5
	fmt.Println(c, "C =", reamur, "R")
	kelvin := c + 273
	fmt.Println(c, "C =", kelvin, "K")
	fahrenheit := (c*9)/5 + 32
	fmt.Println(c, "C =", fahrenheit, "F")

	fmt.Println("Masukkan suhu dalam satuan reamur :")
	fmt.Scanf("%f", &r)
	celcius1 := (r * 5) / 4
	fmt.Println(r, "R =", celcius1, "C")
	kelvin1 := (r/4)/5 + 273.15
	fmt.Println(r, "R =", kelvin1, "K")
	fahrenheit1 := (r*9)/4 + 32
	fmt.Println(r, "R =", fahrenheit1, "F")

	fmt.Println("Masukkan suhu dalam satuan fahrenheit :")
	fmt.Scanf("%f", &f)
	celcius2 := (f - 32) * 1.8
	fmt.Println(f, "F =", celcius2, "C")
	kelvin2 := (f - 459.67) / 1.8
	fmt.Println(f, "F =", kelvin2, "K")
	reamur2 := (f - 32) / 2.25
	fmt.Println(f, "F =", reamur2, "R")

	fmt.Println("Masukkan suhu dalam satuan kelvin :")
	fmt.Scanf("%f", &k)
	celcius3 := (k - 273.5) * 1.8
	fmt.Println(k, "K =", celcius3, "C")
	reamur3 := (k - 273.15) * 0.8
	fmt.Println(k, "K =", reamur3, "R")
	fahrenheit3 := (k * 1.8) - 459.67
	fmt.Println(k, "K =", fahrenheit3, "F")
}
