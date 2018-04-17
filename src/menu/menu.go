package main

import "fmt"
import "os"

func main() {

	exibeIntroducao()       //função que exibe a funcionalidade de introdução
	exibirMenu()            //exibir o menu
	comando := lerComando() //ler o comando de escolha do cliente.

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando desconhecido....")
		os.Exit(-1)
	}

}

func exibeIntroducao() {
	nome := "Dantas"
	idade := 34
	versao := 1.1
	fmt.Println("Olá sr.", nome, "sua idade é", idade, "anos.")
	fmt.Println("Este programa está na versão", versao)
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha váriavel comando é ", &comando)
	fmt.Println("O comando escolhido foi ", comando)

	return comando
}

func exibirMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}
