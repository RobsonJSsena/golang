// Números Primos //

/*
Imprima todos os números primos de 1 a 100.
*/

package main

import "fmt"

func Eprimo(n1 int) bool{
	if n1 < 2 {
		return false
	}
	for i :=2; i*i <= n1; i++ {
		if n1%i == 0{
			return false
		}	
	}
	return true
}

func main(){
	fmt.Println("Números Primos entre 1 a 100")
	for i := 1; i <= 100; i++ {
		if Eprimo(i){
			fmt.Println(i, "")
		}
	}
	fmt.Println()
}