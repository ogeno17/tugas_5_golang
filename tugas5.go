package main

import (
	"fmt"
	"math/rand"
)

type pelajar struct {
	nama string
	umur int
}

func main() {
	var n1, n2, n3 pelajar
	n1.nama = "Ahmad"
	n2.nama = "Abi"
	n3.nama = "Zaka"

	n1.umur = rand.Intn(100)
	n2.umur = rand.Intn(100)
	n3.umur = rand.Intn(100)

	n1.tampilPelajar()
	n2.tampilPelajar()
	n3.tampilPelajar()
}

func (p pelajar) tampilPelajar() {
	fmt.Println(p.nama, p.umur)
}
