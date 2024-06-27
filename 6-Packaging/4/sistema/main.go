package main

import (
	"github.com/google/uuid"
	"github.com/lucastg/pos-go-expert/6-Packaging/3/math"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
	println(uuid.New().String())
}
