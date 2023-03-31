package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const Monitoriamento = 3
const Delay = 5
const ArquivoSites = "Sites.txt"
const LogStatus = "LogStatus.txt"

func main() {
	var vComando int
	Introduction()
	fmt.Println("")

	for {
		displayMenu()
		fmt.Scan(&vComando)

		switch vComando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
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
	/*	vSites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br",
		"https://caelum.com.br",
		"https://gesclinica.com.br"}
	*/
	vSites := leSitesArquivo()
	for vTempo := 0; vTempo < Monitoriamento; vTempo++ {
		for vPosicao, site := range vSites {
			fmt.Println("Testando o Site:", vPosicao, "Site:", site)
			testaSite(site)

			fmt.Println("")
		}
		time.Sleep(Delay * time.Second)
		fmt.Println("")
	}
}

func testaSite(pSite string) {
	resp, error := http.Get(pSite)

	if error != nil {
		fmt.Println("Un error occurred in get url", error)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O Site:", pSite, " -> foi carregado")
		gravaLog(pSite, true)
	} else {
		fmt.Println("O Site:", pSite, " -> não foi carregado. Status Code:", resp.StatusCode)
		gravaLog(pSite, false)
	}
}

func leSitesArquivo() []string {
	var vSites []string
	vArquivo, error := os.Open(ArquivoSites)
	//vArquivo, error := ioutil.ReadFile(ArquivoSites)
	if error != nil {
		fmt.Println("Un error occurred", error)
	} else {
		leitor := bufio.NewReader(vArquivo)

		for {
			linha, error := leitor.ReadString('\n')

			linha = strings.TrimSpace(linha)
			vSites = append(vSites, linha)

			if error == io.EOF {
				break
			}
		}
	}
	vArquivo.Close()
	return vSites
}

func gravaLog(pSite string, pStatus bool) {
	file, err := os.OpenFile(LogStatus, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	} else {
		if pStatus {
			file.WriteString(time.Now().Format("02/01/2006 15:05:05") + " - " + pSite + " - online: " + strconv.FormatBool(pStatus) + "\n")
		} else {
			file.WriteString(time.Now().Format("02/01/2006 15:05:05") + " - " + pSite + " - offline: " + strconv.FormatBool(pStatus) + "\n")
		}
	}

	file.Close()
}

func imprimeLogs() {
	file, err := ioutil.ReadFile(LogStatus)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(file))
	}
}
