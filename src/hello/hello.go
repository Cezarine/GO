package main

import (
	"fmt"
	"os"
)

func main() {
	Introduction()
	var vComando = displayMenu()
	switch vComando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Exiting")
		os.Exit(0)
	default:
		fmt.Println("Comando não encontrado")
	}

}

func Introduction() {
	const vVersao float32 = 1.1
	var vNome string
	fmt.Println("Qual seu nome?")
	fmt.Scan(&vNome)
	fmt.Println("Olá, Sr/Sra.", vNome)
	fmt.Println("Este programa esta na versao", vVersao)
}

func displayMenu() int {
	var vComando int
	fmt.Println("1- Iniciar Monitoriamento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")
	fmt.Scan(&vComando)

	return vComando
}
