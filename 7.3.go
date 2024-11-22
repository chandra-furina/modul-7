package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var klubA, klubB string
	var pemenang []string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input nama Klub A: ")
	klubA, _ = reader.ReadString('\n')
	klubA = strings.TrimSpace(klubA)

	fmt.Print("input nama Klub B: ")
	klubB, _ = reader.ReadString('\n')
	klubB = strings.TrimSpace(klubB)

	fmt.Println("input skor pertandingan. Masukkan skor negatif untuk berhenti.")

	// Input skor pertandingan
	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d: ", i)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		scores := strings.Split(input, " ")

		if len(scores) != 2 {
			fmt.Println("Format input tidak valid. input dua skor dipisahkan spasi.")
			i--
			continue
		}

		skorA, errA := strconv.Atoi(scores[0])
		skorB, errB := strconv.Atoi(scores[1])

		if errA != nil || errB != nil {
			fmt.Println("Skor harus berupa angka.")
			i--
			continue
		}

		if skorA < 0 || skorB < 0 {
			fmt.Println("Pertandingan selesai.")
			break
		}

		if skorA > skorB {
			fmt.Printf("// %s = %d, %s = %d\n", klubA, skorA, klubB, skorB)
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			fmt.Printf("// %s = %d, %s = %d\n", klubA, skorA, klubB, skorB)
			pemenang = append(pemenang, klubB)
		} else {
			fmt.Printf("// %s = %d, %s = %d\n", klubA, skorA, klubB, skorB)
			pemenang = append(pemenang, "Draw")
		}
	}

	fmt.Println("\nHasil Pertandingan:")
	for i, hasil := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, hasil)
	}
}
