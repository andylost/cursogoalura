package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	for {

		exibeIntroducao()       //função que exibe a funcionalidade de introdução
		exibirMenu()            //exibir o menu
		comando := lerComando() //ler o comando de escolha do cliente.

		switch comando {
		case 1:
			iniciarMonitoramento()
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

}

func exibeIntroducao() {
	fmt.Println("")
	nome := "Dantas"
	idade := 34
	//versao := 1.1
	fmt.Println("Olá sr.", nome, "sua idade é", idade, "anos.")
	//fmt.Println("Este programa está na versão", versao)
	fmt.Println("")
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	// fmt.Println("O endereço da minha váriavel comando é ", &comando)
	fmt.Println("O comando escolhido foi ", comando)
	fmt.Println("")

	return comando
}

func exibirMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println("")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	fmt.Println(" ")

	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br",
		"https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testeSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func testeSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site está no ar", site, "Status Code: ", resp.StatusCode)
	} else {
		fmt.Println("Site está fora ar", site, "Status Code: ", resp.StatusCode)
	}

}
