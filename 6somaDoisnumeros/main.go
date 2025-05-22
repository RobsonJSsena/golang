//Somando 2 números//

/*
Crie uma função em golang que recebe dois números inteiros e retorna a soma deles.
*/

package main 

import "fmt"

func soma(n1, n2 int) int{
    return n1 + n2

}

func main(){
	somando := soma(250, 250)
	fmt.Println("O resultado:",somando)
   
	
   soma2 := 100
   soma3 := 300
   resultadoFinal := soma2 + soma3
   fmt.Println("O valor da soma:",resultadoFinal)
}






