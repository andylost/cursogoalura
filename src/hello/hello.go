package main

import (
	"fmt"
	"math"
	"math/rand"
)

func calcular(x int, y int) int {
	return x + y
}
func swap(x string, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	var name string
	name = "dantas"

	fmt.Println("Randon", rand.Intn(10))
	fmt.Println("Randon", math.Sqrt(9))
	fmt.Println("Randon", math.Pi)
	fmt.Println("Olá Dantas", name)
	fmt.Println("Valor somado é de ", calcular(2, 33))

	a, b := swap("Hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(12223))

	var c, python, java bool
	var i int
	fmt.Println(i, c, python, java)

}
