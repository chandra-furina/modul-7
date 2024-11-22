package main

import (
	"fmt"
	"math"
)

func rataRata(array []int) float64 {
	total := 0
	for _, val := range array {
		total += val
	}
	return float64(total) / float64(len(array))
}

func standarDeviasi(array []int) float64 {
	mean := rataRata(array)
	var variance float64
	for _, val := range array {
		variance += math.Pow(float64(val)-mean, 2)
	}
	variance /= float64(len(array))
	return math.Sqrt(variance)
}

func frekuensi(array []int, target int) int {
	count := 0
	for _, val := range array {
		if val == target {
			count++
		}
	}
	return count
}

func main() {
	var n int

	fmt.Print("jumlah elemen array (N): ")
	fmt.Scan(&n)

	array := make([]int, n)
	fmt.Println("elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen %d: ", i)
		fmt.Scan(&array[i])
	}

	fmt.Println("a. Keseluruhan isi array:")
	fmt.Println(array)

	fmt.Println("b. Elemen dengan indeks ganjil:")
	for i := 1; i < len(array); i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	fmt.Println("c. Elemen dengan indeks genap:")
	for i := 0; i < len(array); i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("input nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Println("d. Elemen dengan indeks kelipatan x:")
	for i := 0; i < len(array); i++ {
		if i%x == 0 {
			fmt.Print(array[i], " ")
		}
	}
	fmt.Println()

	var delIndex int
	fmt.Print("indeks yang ingin dihapus: ")
	fmt.Scan(&delIndex)
	if delIndex >= 0 && delIndex < len(array) {
		array = append(array[:delIndex], array[delIndex+1:]...)
		fmt.Println("e. Array setelah elemen dihapus:")
		fmt.Println(array)
	} else {
		fmt.Println("Indeks tidak valid!")
	}

	fmt.Printf("f. Rata-rata elemen array: %.2f\n", rataRata(array))

	fmt.Printf("g. Standar deviasi elemen array: %.2f\n", standarDeviasi(array))

	var target int
	fmt.Print("inputbilangan untuk menghitung frekuensi: ")
	fmt.Scan(&target)
	fmt.Printf("h. Frekuensi bilangan %d dalam array: %d\n", target, frekuensi(array, target))
}
