//Fatorial de um Número//
/*
o fatorial de um número (representado por n! na matemática), basta multiplicar o número por
todos os inteiros positivos menores que ele até chegar a 1.
Por exemplo, 5! = 5 * 4 * 3 * 2 * 1 = 120. Sendo assim:
Escreva uma função que calcule o fatorial de um número.
*/

package main

import "fmt"

func fatorialDeUmNumero(n1 int) int{
	resultado := 1
	expressao := ""

	for i := n1; i >= 1; i--{
		resultado *= i
	if i == 1{
		expressao += "1"
	 }else {
		expressao += fmt.Sprintf("%d * ", i)
	 }
	}
   
   fmt.Printf("O Fatorial de %d! = %s = %d\n", n1, expressao, resultado)
   return resultado
}

func main(){
	var numero int
	fmt.Println("Fatorial de Um Número")
	fmt.Println("Digite um Número")
	fmt.Scanln(&numero)

	if numero < 0 {
		fmt.Println("Fatorial nõa está definido para números negativos")
	}else{
		fatorialDeUmNumero(numero)
	}

	//Sintaxe do código desenvolvida dessa forma para que ao imprimir sair dessa forma: //
	//Digite um número: 5 -> O fatorial de 5! = 5 * 4 * 3 * 2 * 1 = 120//

}



/*Explicando o código em detalhes:

 O objetivo dele é calcular o fatorial de um número e imprimir também a 
expressão da multiplicação que levou ao resultado.

resultado := 1 -> vai acumular o valor do fatorial.

expressao := "" -> é uma string que vai montar o texto da multiplicação (ex: "5 * 4 * 3 * 2 * 1").

Loop que calcula o fatorial

for i := n1; i >= 1; i-- {
		resultado *= i

Laço que vai de n1 até 1, decrementando (i--).

A cada passo, multiplica resultado pelo valor de i		


if i == 1 {
		expressao += "1"
	} else {
		expressao += fmt.Sprintf("%d * ", i)
	}

Monta a expressão como string:

Se for o último número (1), apenas adiciona "1".

Caso contrário, adiciona "i * ".
*/