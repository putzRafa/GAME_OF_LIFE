package main

import (
	"fmt"
	"math/rand"
)

func imprimeMatriz(matriz [10][10]int) {
	var i int
	var j int

	for i = 0; i < len(matriz); i++ {
		for j = 0; j < len(matriz[i]); j++ {
			fmt.Print(matriz[i][j], "  ")
		}
		fmt.Println()
	}
}

func criaGradeManual() [10][10]int {
	var matriz [10][10]int
	var linha int
	var coluna int

	for {
		fmt.Scan(&linha, &coluna)
		if linha == -1 && coluna == -1 {
			break
		}

		if linha >= 0 && linha < 10 && coluna >= 0 && coluna < 10 {
			matriz[linha][coluna] = 1
		}
	}
	return matriz
}

func criaGradeRand() [10][10]int {
	var matriz [10][10]int
	var i int
	var j int
	var randon int

	for i = 0; i < len(matriz); i++ {
		for j = 0; j < len(matriz[i]); j++ {
			randon = rand.Intn(100)
			if randon > 85 {
				matriz[i][j] = 1
			}
		}
	}
	return matriz
}

//EM DESENVOLVIMENTO
/*
func contVizinhos(matriz [10][10]int)  {
	var i, j int
	var linha, coluna int
	var vizinhoVivo int

	linha = len(matriz)
	coluna = len(matriz[0])

	for linha = i-1; linha <= i+1; i++{
		for coluna = j-1; coluna <= j+1; j++{
			if linha == i && coluna == j {
				continue
			}

		}
	}

}*/

func main() {
	var grade [10][10]int
	var opcao int

	fmt.Println("Selecione a opção que deseja:")
	fmt.Println("1 - SELEÇÃO MANUAL")
	fmt.Println("2 - CÉLULAS ALEATÓRIAS")

	fmt.Scanln(&opcao)
	if opcao == 1 {
		fmt.Println("-----------------------SELECÃO MANUAL-----------------------")
		fmt.Println("Escolha o local onde deseja colocar a célula. Ex: 1 1")
		fmt.Println("Digite -1-1 para confirmar")
		grade = criaGradeManual()
	} else {
		if opcao == 2 {
			grade = criaGradeRand()
		} else {
			fmt.Println("OPÇÃO INVÁLIDA")
			fmt.Scanln(&opcao)
		}
	}
	imprimeMatriz(grade)
}
