package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){

	file, err := os.Open("./calorias.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	
	defer file.Close()

	// Cria um leitor de arquivo linha por linha
	scanner := bufio.NewScanner(file)

	// Leitura e impressão do arquivo .txt
	var elfo1, elfo2, elfo3 int
	var fulanoElfos int = 1
	var maior1 int = 0
	var maior2 int = 0
	var maior3 int = 0
	var soma, somaCalElfos int
	for scanner.Scan(){
		line := scanner.Text()
		if line == "" {
			soma = 0
			fulanoElfos++
		} else{
			num, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Erro ao converter linha em número:", err)
				continue
			}
			soma += num 
			if soma > maior3 { // soma = 2000
				maior3 = soma // maior3 = 2000 
				elfo3 = fulanoElfos // elfo3 = 54
				if maior3 > maior2 { // 
					maior3, maior2 = maior2, maior3 // maior3 = 0 ; maior2 = 2000
					elfo3, elfo2 = elfo2, elfo3 // elfo3 = 0; elfo2 = 54
					if maior2 > maior1 { //
						maior2, maior1 = maior1, maior2 // maior2 = 0 ; maior1 = 2000
						elfo2, elfo1 = elfo1, elfo2 // elfo2 = 0 ; elfo1 = 54
					}
				}
			}
		}	
	}
	somaCalElfos = maior1 + maior2 + maior3
	fmt.Printf("TOP3 elfos em ordem: %d:%d, %d:%d, %d:%d\n", elfo1, maior1, elfo2, maior2, elfo3, maior3)
	fmt.Printf("Soma dos TOP3 elfo: %d", somaCalElfos)
}


