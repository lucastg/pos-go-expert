package main

func main() {

	println("------------------------")
	for i := 0; i < 10; i++ {
		println(i)
	}
	println("------------------------")
	numeros := []string{"um", "dois", "tres"}
	for k, v := range numeros {
		println(k, v)
	}
	println("------------------------")

	// OU (UTILIZANDO BLANK IDENTIFIER)

	// numeros2 := []string{"um", "dois", "tres"}
	// for _, v := range numeros2 {
	// 	println(v)
	// }

	// OU (UTILIZANDO BLANK IDENTIFIER)

	// numeros3 := []string{"um", "dois", "tres"}
	// for k, _ := range numeros3 {
	// 	println(k)
	// }

	// WHILE
	i := 0
	for i < 10 {
		println(i)
		i++
	}
	println("------------------------")

	// LOOP INFINITO
	// for {
	// 	println("Mitsubitch Lancer")
	// }
}
