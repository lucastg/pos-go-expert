package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lucastg/pos-go-expert/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Mitsubish"}

	fmt.Println("Resultado: ", s)
	fmt.Println(carro.Andar())
	fmt.Println(uuid.New())
}
