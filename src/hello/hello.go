package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const Monitoriamento = 5
const Delay = 5

func main() {
	var vComando int
	Introduction()
	for {
		displayMenu()
		fmt.Scan(&vComando)
		switch vComando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Exiting")
			os.Exit(0)
		default:
			fmt.Println("Comando não encontrado")
		}
	}

}

func Introduction() {
	const vVersao float32 = 1.1
	var vNome string = "Guilherme"
	fmt.Println("Qual seu nome", vNome)
	fmt.Println("Este programa esta na versao", vVersao)
}

func displayMenu() {

	fmt.Println("1- Iniciar Monitoriamento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")
}

func iniciarMonitoramento() {
	vSites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br",
		"https://caelum.com.br",
		"https://gesclinica.com.br"}

	for vTempo := 0; vTempo < Monitoriamento; vTempo++ {
		for vPosicao, site := range vSites {
			fmt.Println("Testando o Site:", vPosicao, "Site:", site)
			testaSite(site)
		}
		time.Sleep(Delay * time.Second)
		fmt.Println("")
	}
}

func testaSite(pSite string) {
	resp, _ := http.Get(pSite)
	if resp.StatusCode == 200 {
		fmt.Println("O Site:", pSite, " -> foi carregado")
	} else {
		fmt.Println("O Site:", pSite, " -> não foi carregado. Status Code:", resp.StatusCode)
	}
}
