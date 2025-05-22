// Área do retângulo //

/*
Crie uma função que recebe base e altura e retorna a área do retângulo, e outra com função 
que retorne o perímetro.
*/

package main

import "fmt"

func areaRetangulo(base, altura int) (int){
    return base * altura
	
}

func perimetroRetangulo(base, altura int)(int){
	return 2 * (base + altura)
}

func main(){
	var base, altura int

	fmt.Println("Digite o tamanho da base em cm: ")
	fmt.Scanln(&base)

	fmt.Println("Digite o tamanho da altura cm: ")
	fmt.Scanln(&altura)

	medidaTotal := areaRetangulo(base, altura)
	fmt.Println("Medida do Retângulo ->",medidaTotal,"cm")

	medidaDoPerimetro := perimetroRetangulo(base, altura)
	fmt.Println("Medida do Perímetro de um Retângulo:", medidaDoPerimetro,"cm")



}