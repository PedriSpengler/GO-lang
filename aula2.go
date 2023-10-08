package main

import(
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func main(){
	file, err := os.Open("./calorias.txt")
	if err != nil{
		fmt.Println("Erro ao abrir o arquivo:!", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var elfoFulano int = 0
	var maiorCaloria int = 0 ; var menorCaloria int = 0

	for scanner.Scan(){
		line := scanner.Text()
		if line == "" {
			elfoFulano++
			menorCaloria = 0
		} else{
			num, err := strconv.Atoi(line)
			if err != nil{
				fmt.Println("Erro ao converter string para inteiro:", err)
				continue
			}
			menorCaloria += num
			if menorCaloria > maiorCaloria{
				maiorCaloria = menorCaloria
			}
		}
	}
	fmt.Println("Elfo com maior número de calorias:", elfoFulano)
	fmt.Println("Maior número de calorias:", maiorCaloria)
}
