package main

import (
	"bufio"
	"fmt"
	"os"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan teks diakhiri dengan titik: ")

	*n = 0
	for {
		ch, _, _ := reader.ReadRune()
		if ch == '.' || *n >= NMAX {
			break
		}
		(*t)[*n] = ch
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	isiArray(&tab, &n)

	fmt.Print("Teks: ")
	cetakArray(tab, n)

	if palindrom(tab, n) {
		fmt.Println("Palindrome? true")
	} else {
		fmt.Println("Palindrome? false")
	}
}
