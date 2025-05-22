//Temperatura Celsius para fahrenheit //

/*
Faça um programa que receba uma temperatura em Celsius e converta para Fahrenheit.
*/

package main

import "fmt"

func celsiusFahrenheit(temp float64) float64{
	return (temp * 9 / 5) + 32
}

func main(){
	var temp float64
      fmt.Println("Digite a temperatura: ")
	   fmt.Scanln(&temp)

   temperaturaAtual :=celsiusFahrenheit(temp)
	 fmt.Printf("O termômetor mostra: %.2f°C é em fahrenheit fica: %.2f°F\n", temp, temperaturaAtual)

}
/*

Podemos usar de 2 formas para imprimir o resulato
Usamos da forma mais simples, com stings, e usando caracteres °C:

fmt.Println("Termômetro está medindo:",temp,"°C","e em fahrenheit fica:",temperaturaAtual

ou podemos usar: 

fmt.Printf("%.2f°c é igual a %.2f°f\n", celsius, temperatura//

Explicando a linha de comando:
fmt.Printf("%.2f°c é igual a %.2f°f\n", celsius, temperatura
	  
fmt.Printf

fmt: pacote da biblioteca padrão do Go que trata de formatação de entrada e saída (como imprimir na tela).

Printf: função que imprime uma string formatada.

2. "%.2f°C é igual a %.2f°F\n"

Essa é a string de formatação:

%.2f: significa que queremos imprimir um número de ponto flutuante (float) com 2 casas decimais.

°C e °F: são apenas caracteres literais (símbolos de graus Celsius e Fahrenheit).

\n: é um caractere de nova linha (enter).

Então, a string formatada espera dois valores numéricos, que aparecerão com 2 casas decimais, 
com as respectivas unidades.

Por que usamos +32 e não +212?
Quando convertemos Celsius para Fahrenheit, estamos tentando alinhar os 
pontos das duas escalas, e o ponto de congelamento da água é o nosso referencial base, 
não o ponto de ebulição.

Resumindo tudo o 32 é usado porque, é o valor em Fahrenheit que equivalente a 0 °C, 
que é o ponto de congelamento da água. Já o 212 só aparece quando 100 °C é convertido 
ele é o resultado final para o ponto de ebulição, não um valor que precisa ser somado.


*/

