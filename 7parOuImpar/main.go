// É par ou Impar //

/*
Crie uma função que recebe dois números inteiros e retorna a soma deles,
é informe se o resultado é par ou ímpar!
*/

package main

import "fmt"

func NumeroParOuImpar(number1, number2 int)(int, string){
	somaDosNumeros := number1 + number2
    if(somaDosNumeros % 2 == 0){
		return somaDosNumeros, "PAR"
	}else {
		return somaDosNumeros, "IMPAR"
	}
}

func main(){
	var number1, number2 int

	fmt.Println("Digite o Primeiro número")
	fmt.Scanln(&number1)

	fmt.Println("Digite o Segundo número")
	fmt.Scanln(&number2)

	somaDosNumeros, ParOuImpar := NumeroParOuImpar(number1, number2)
	 fmt.Println("Resultado da Soma foi:",somaDosNumeros, "\nEsse valor é um número:",ParOuImpar)
}