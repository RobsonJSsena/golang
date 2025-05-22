// Calculadora Simples //

package main

import "fmt"

func main(){

var num1, num2 float64
var operador string

//Entrada do Primerio Número//
fmt.Println("Digite o Primeiro Número: ")
fmt.Scanln(&num1)

//Entrada do Segundo Número//
fmt.Println("Digite o Segundo Número: ")
fmt.Scanln(&num2)

//Entrada do Operador//
fmt.Println("Digite o operador Matemático que Deseja (+,-,*,/): ")
fmt.Scanln(&operador)

switch operador{
case "+":
  fmt.Println("Resultado", num1+num2)
case "-":
  fmt.Println("Resultado", num1-num2)
case "*":
  fmt.Println("Resultado", num1*num2)
  case "/":
	if num2 != 0{
		fmt.Println("Resultado", num1/num2)
	}else{
		fmt.Println("Erro: divisão por zero")
	}  
	default:
		fmt.Println("Operação inválida")	 	  	
  }
}
