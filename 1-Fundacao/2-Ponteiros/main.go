package main

func main() {
	a := 10
	var ponteiro *int = &a
	println(ponteiro)

	*ponteiro = 20
	b := &a
	println(*b)

	*b = 30
	println(a)
}
