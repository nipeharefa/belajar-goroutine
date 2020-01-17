package main

import "fmt"

type Angka int

func fungsi(a int) Angka {
	return Angka(a + 1)
}
func main() {

	// Unbuffered channel
	var angka = make(chan Angka)

	// Anonym func
	var tambah1 = func(i int) {
		angka <- fungsi(i)
	}

	go tambah1(1)
	go tambah1(100)
	go tambah1(3)
	// go fungsi()
	// fmt.Println("S")

	fmt.Println(<-angka)
	fmt.Println(<-angka)
	fmt.Println(<-angka)

	close(angka)
}
