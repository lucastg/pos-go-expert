package main

type myNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T Number](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Lucas": 1000, "Joao": 2000, "Maria": 3000}
	m2 := map[string]float64{"Lucas": 1000.2, "Joao": 2000.3, "Maria": 3000.4}
	m3 := map[string]myNumber{"Lucas": 1000, "Joao": 2000, "Maria": 3000}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10.0))
}
