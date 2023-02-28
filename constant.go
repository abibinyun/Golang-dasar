package main

import "fmt"

func main() {
	const firstName string = "abi"
	const lastName = "bilal"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		namaDepan    string = "abi"
		namaBelakang        = "ganteng"
	)
	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
}
