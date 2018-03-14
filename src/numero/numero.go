package main

import (
	"fmt"
	"reflect"
)

func main() {
	var precoLeite float32 = 2.99
	var precoOvo float32 = 1.99
	var precoPao float32 = 3.99

	fmt.Println("O Valor do Leite é de: R$", precoLeite)
	fmt.Println("O Valor do Ovo é de: R$", precoOvo)
	fmt.Println("O Valor do Pão é de: R$", precoPao)

	fmt.Println("Tipo de objeto", reflect.TypeOf(precoLeite))

}
