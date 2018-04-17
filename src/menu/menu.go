package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
			lerSitesdoArquivo()
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

	sites := lerSitesdoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testeSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func testeSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro o acesso ao site: ", err)
		return
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site está no ar", site, "Status Code: ", resp.StatusCode)
	} else {
		fmt.Println("Site está fora ar", site, "Status Code: ", resp.StatusCode)
	}

}

func lerSitesdoArquivo() []string {
	arquivo, err := os.Open("sites.txt")

	var sites []string

	if err != nil {
		fmt.Println("Erro ao Obter o arquivo ", err)
		return nil
	}

	leitor := bufio.NewReader(arquivo)

	// linha, err := leitor.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Erro ao ler a linha do arquivo ", err)
	// 	return nil
	// }
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()

	return sites
}
