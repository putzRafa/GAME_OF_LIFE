package main

import (
	"fmt"
	"math/rand"
	"time"
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

func condicoes(matriz [10][10]int) [10][10]int {
	var novaMatriz [10][10]int
	var i, ni int
	var j, nj int
	var dx int
	var dy int
	var vizinhosVivos int

	// Percorrer cada célula da matriz
	for i = 0; i < 10; i++ {
		for j = 0; j < 10; j++ {
			vizinhosVivos = 0

			// Verifica todos os vizinhos ao redor
			for dx = -1; dx <= 1; dx++ {
				for dy = -1; dy <= 1; dy++ {
					if dx == 0 && dy == 0 {
						continue // ignora a própria célula
					}

					// Coordenadas do vizinho
					ni = i + dx
					nj = j + dy

					if ni >= 0 && ni < 10 && nj >= 0 && nj < 10 {
						if matriz[ni][nj] == 1 {
							vizinhosVivos++
						}
					}
				}
			}

			// Aplica as regras
			if matriz[i][j] == 1 {
				if vizinhosVivos < 2 || vizinhosVivos > 3 {
					novaMatriz[i][j] = 0 // morre
				} else {
					novaMatriz[i][j] = 1 // sobrevive
				}
			} else {
				if vizinhosVivos == 3 {
					novaMatriz[i][j] = 1 // nasce
				} else {
					novaMatriz[i][j] = 0 // continua morta
				}
			}
		}
	}
	return novaMatriz
}

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

	for {
		grade = condicoes(grade)
		fmt.Println("\nNOVA GERAÇÃO")
		imprimeMatriz(grade)
		time.Sleep(time.Second)
	}
}
