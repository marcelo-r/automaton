package main

import (
	"automaton/dfa"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var auto = dfa.NewDFA(0, false)
var infoFlag = flag.Bool("info", false, "Imprimir detalhes do funcionamento do automato")

const regras = `
alfabeto = {a, b, c}

1. palavras terminam apenas em 'b' ou 'c', nunca em 'a'
2. b não pode ser concatenado a outro 'b' antes que surja um 'c'
3. a aparição de qualquer 'c' permite concatenação de 'b's posteriores
4. nenhum 'a' é permitido após qualquer 'c'
5. qualquer palavra sempre termina com o simbolo de maior "valor" na palavra`

const gdef = "\n\u03a3 = {a, b, c}\nV = {S, X, Y, Z, T, W, V, K, L, M}\nS = S\nP: "

const gp = `
  S -> aS | K
  K ->  b | bX
  X -> aY | L
  L ->  c | cZ
  Y -> aY | K | L
  Z -> bT | L
  T -> bT | M
  M ->  c | cW
  W -> bT | M`

var INFO = *infoFlag

func main() {
	mainMenu()
}

func init() {
	flag.Parse()

	// adicionar todos os estados
	auto.AddState(1, true) // f
	auto.AddState(2, false)
	auto.AddState(3, true) // f
	auto.AddState(4, false)
	auto.AddState(5, true) // f
	auto.AddState(6, false)

	// o automato deve ter um alfabeto antes de poder ter alguma transição
	auto.AddToAlphabet('a')
	auto.AddToAlphabet('b')
	auto.AddToAlphabet('c')

	// transições no formato delta(estado, entrada) = estado_destino
	// 0
	auto.AddTransition(0, 'a', 0)
	auto.AddTransition(0, 'b', 1)
	auto.AddTransition(0, 'c', 6)

	// 1
	auto.AddTransition(1, 'a', 2)
	auto.AddTransition(1, 'b', 6)
	auto.AddTransition(1, 'c', 3)

	// 2
	auto.AddTransition(2, 'a', 2)
	auto.AddTransition(2, 'b', 1)
	auto.AddTransition(2, 'c', 3)

	// 3
	auto.AddTransition(3, 'a', 6)
	auto.AddTransition(3, 'b', 4)
	auto.AddTransition(3, 'c', 3)

	// 4
	auto.AddTransition(4, 'a', 6)
	auto.AddTransition(4, 'b', 4)
	auto.AddTransition(4, 'c', 5)

	// 5
	auto.AddTransition(5, 'a', 6)
	auto.AddTransition(5, 'b', 4)
	auto.AddTransition(5, 'c', 5)

	// 6
	auto.AddTransition(6, 'a', 6)
	auto.AddTransition(6, 'b', 6)
	auto.AddTransition(6, 'c', 6)
}

func mainMenu() {
	counter := make(map[string]int)
	counter["r"], counter["a"] = 0, 0
	fmt.Println("Bem vindo")
	listarRegras()
	leitor := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\n$ ")
		palavra, _ := leitor.ReadString('\n')
		palavra = strings.TrimSuffix(palavra, "\n")
		if palavra == "\\q" {
			fmt.Printf("\nresultado\na: %d \nr: %d\n", counter["a"], counter["r"])
			break
		}
		ok := auto.CheckWord(palavra)
		if ok {
			counter["a"]++
			fmt.Println("+ aceita +")
		} else {
			counter["r"]++
			fmt.Println("- rejeita -")
		}
	}

}

func listarRegras() {
	fmt.Println(regras)
	if *infoFlag {
		fmt.Println(gdef + gp)
	}
}
