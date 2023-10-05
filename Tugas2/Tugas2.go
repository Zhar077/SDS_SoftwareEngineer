/*
Soal:
1. Buatlah algoritma pengurutan bilangan dari yang terbesar sampai yang terkecil!
2. Buatlah slice dengan jumlah numbers di dalam slice yaitu ada 5 buah dan tambahkan numbers ke dalam slice sebanyak 3!
*/
package main

import "fmt"

func main() {
	fmt.Println("1. Algoritma pengurutan bilangan dari yang terbesar sampai yang terkecil")
	fmt.Print("Array sebelum diurutkan: ")
	numbers := []int{77, 44, 33, 22, 11, 66, 55}
	fmt.Println(numbers)
	Sorter(numbers)
	fmt.Print("Array setelah diurutkan: ")
	fmt.Println(numbers)
	fmt.Println("2. slice dengan jumlah numbers di dalam slice yaitu ada 5 buah dan tambahkan numbers ke dalam slice sebanyak 3")
	slice()
}

func Sorter(numbers []int) {

	n := len(numbers)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] < numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}

func slice() {
	var NUM = []int{19, 20, 21, 22, 23}
	fmt.Print("Slice Awal: ")
	fmt.Println(NUM)
	NUM = append(NUM, 24, 25, 26)
	fmt.Print("Slice Awal: ")
	fmt.Println(NUM)
}
