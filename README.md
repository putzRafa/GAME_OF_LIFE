# GAME OF LIFE

Uma implementação simples do **Jogo da Vida de Conway**, desenvolvida em **Go** como um projeto pessoal para colocar em prática os conhecimentos adquiridos durante o primeiro período de Engenharia de Software.

O programa simula a evolução de uma população de células ao longo de diferentes gerações, aplicando as regras do Jogo da Vida a cada nova geração.

## 📖 Sobre o projeto

O **Jogo da Vida de Conway** é um autômato celular desenvolvido pelo matemático britânico John Horton Conway.

Apesar de possuir regras relativamente simples, o sistema é capaz de produzir diferentes comportamentos a partir da configuração inicial das células.

Neste projeto, a simulação utiliza uma **grade de 10 × 10 posições**, na qual cada posição pode representar uma célula viva ou morta.

A cada geração, o programa verifica os vizinhos de cada célula e determina seu próximo estado de acordo com as regras do jogo.

## ⚙️ Funcionamento

O programa possui duas formas de definir a configuração inicial da grade:

### ✏️ Seleção manual

O usuário pode escolher as posições onde deseja criar as células.

Para isso, são informados os valores correspondentes à **linha** e à **coluna**:

```text
Exemplo:
1 1
2 2
2 3
```

Para finalizar a configuração manual, o usuário deve informar:

```text
-1 -1
```

As posições válidas estão entre `0` e `9`, tanto para linhas quanto para colunas.

### 🎲 Geração aleatória

Também é possível gerar automaticamente uma configuração inicial.

Nesse modo, cada posição da matriz recebe um valor aleatório e existe uma probabilidade de aproximadamente **15%** de uma célula ser criada naquela posição.

## 🧬 Regras do Jogo da Vida

A cada geração, o programa analisa os **oito vizinhos** de cada posição da grade.

As regras utilizadas são:

### Célula viva

Uma célula viva:

* **Morre por solidão** quando possui menos de 2 vizinhos vivos;
* **Continua viva** quando possui 2 ou 3 vizinhos vivos;
* **Morre por superpopulação** quando possui mais de 3 vizinhos vivos.

### Célula morta

Uma célula morta:

* **Nasce** quando possui exatamente 3 vizinhos vivos;
* Permanece morta em qualquer outra situação.

Essas regras são aplicadas simultaneamente para produzir a próxima geração.

## 🔄 Evolução das gerações

Após criar a configuração inicial, o programa passa a calcular novas gerações.

A cada geração:

1. A grade atual é analisada;
2. Os vizinhos de cada posição são contabilizados;
3. As regras do Jogo da Vida são aplicadas;
4. Uma nova grade é criada;
5. A nova geração é exibida no terminal;
6. O programa aguarda 2 segundos antes de continuar.

A simulação termina quando **não existe nenhuma alteração entre uma geração e a seguinte**.

## 🖥️ Exemplo de execução

Uma configuração inicial pode ser apresentada no terminal como:

```text
0  0  0  0  0  0  0  0  0  0
0  0  1  0  0  0  0  0  0  0
0  0  0  1  0  0  0  0  0  0
0  1  1  1  0  0  0  0  0  0
0  0  0  0  0  0  0  0  0  0
...
```

A cada nova geração, a matriz é atualizada de acordo com o número de vizinhos vivos de cada célula.

## 🧩 Estrutura do código

O projeto foi dividido em funções responsáveis por diferentes partes da simulação:

| Função              | Responsabilidade                             |
| ------------------- | -------------------------------------------- |
| `imprimeMatriz()`   | Exibe a matriz no terminal                   |
| `criaGradeManual()` | Permite ao usuário criar células manualmente |
| `criaGradeRand()`   | Gera uma configuração inicial aleatória      |
| `condicoes()`       | Aplica as regras do Jogo da Vida             |
| `matrizesIguais()`  | Verifica se duas gerações são iguais         |
| `main()`            | Controla o fluxo principal da aplicação      |

Essa divisão permite separar as principais etapas da simulação e facilita a compreensão do funcionamento do algoritmo.

## 🛠️ Tecnologias utilizadas

* **Go**
* Matrizes bidimensionais
* Estruturas de repetição
* Estruturas condicionais
* Funções
* Geração de números aleatórios
* Manipulação de tempo
* Entrada e saída pelo terminal

## 📂 Estrutura do repositório

```text
GAME_OF_LIFE/
│
├── .vscode/
│
├── GAME_OF_LIFE.go
├── go.mod
├── README.md
└── __debug_bin.exe
```

O arquivo principal da aplicação é:

```text
GAME_OF_LIFE.go
```

## ▶️ Como executar

### Pré-requisito

É necessário possuir o **Go** instalado no computador.

### Executando pelo terminal

Clone o repositório:

```bash
git clone https://github.com/putzRafa/GAME_OF_LIFE.git
```

Entre na pasta:

```bash
cd GAME_OF_LIFE
```

Execute o programa:

```bash
go run GAME_OF_LIFE.go
```

Depois, escolha uma das opções apresentadas:

```text
1 - SELEÇÃO MANUAL
2 - CÉLULAS ALEATÓRIAS
```

## 🎯 Objetivo do projeto

Este projeto foi desenvolvido principalmente como uma forma de **praticar programação e consolidar os conhecimentos adquiridos no primeiro período de Engenharia de Software**.

Durante seu desenvolvimento, foram explorados conceitos importantes da programação, como:

* Manipulação de matrizes;
* Funções;
* Estruturas de repetição;
* Estruturas condicionais;
* Entrada e saída de dados;
* Geração de números aleatórios;
* Comparação de estruturas;
* Simulação de processos iterativos;
* Organização de um programa em diferentes funções.

Além de servir como exercício de programação, o projeto também apresenta um exemplo prático de como regras simples podem produzir diferentes comportamentos ao longo de sucessivas gerações.

## 👨‍💻 Autor

**Rafael Farias de Lima**

Projeto pessoal desenvolvido durante o primeiro período de **Engenharia de Software**, com o objetivo de praticar conceitos fundamentais de programação utilizando **Go**.

---

⭐ Obrigado por visitar o projeto!
