package main

func main() {
	a := 10
	b := 2
	c := 3

	if a > b && c > a {
		println("a > b && c > a")
	}

	if a > b || c > a {
		println("a > b || c > a")
	}

	if a > b {
		println(a)
	} else {
		println(b)
	}

	switch a {
	case 1:
		println("A")
	case 2:
		println("B")
	case 3:
		println("C")
	default:
		println("D")
	}
}
