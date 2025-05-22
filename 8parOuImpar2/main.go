// Par ou Impar 2, Usando apenas um número //

/*
Escreva uma função que receba um número e diga se ele é par ou ímpar.
*/

package main

import "fmt"

func EparOuImpar(n1 int)(int, string) {
	imparPar := n1
	if(imparPar % 2 == 0){
		return imparPar, "Par"
   }else{
	return imparPar, "Impar"
   }
}

func main(){
	var n1 int
	fmt.Println("Digito o Número que deseja")
	fmt.Scanln(&n1)

     imparPar, numeroEscolhido  := EparOuImpar(n1)
	 fmt.Println("O Número escolhido foi: ", imparPar," e ele é:", numeroEscolhido)
}