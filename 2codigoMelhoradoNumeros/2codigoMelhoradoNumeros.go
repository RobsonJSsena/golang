                 // Numero 2.0 //

/*
Usaremos como exemplo do arquivo numero.go,
alterando o cógido para melhor entendimento e sofisticação:

package main

import "fmt"

func main() {
	var a int = 10
    var b int = 20
	var c int = 30
	var d int = 40

	fmt.Println("O resultado do calculo:", a + b)
	fmt.Println("O resultado do calculo:", c * d)
	fmt.Println("O resultado do calculo:", b / a)
	fmt.Println("O resultado do calculo:", d - c)
}
*/

//usaremos função com retorno para alterar e trazer uma melhor sintaxe ao código//

package main

//import "fmt"

func calculo(n1, n2 int) (int, int, int){
	soma := n1 + n2
	subtracao := n1 - n2
	multiplica := n1 * n2
	return soma, subtracao, multiplica
}

//func main(){
//	soma, _, _ := calculo(10, 5)//usamos o underscore (_) para ignorar subracao, e multiplica//
//	_, subtracao, _ := calculo(20, 2)//usamos o underscore (_) para ignorar soma e multiplica//
//	_, _, multiplica := calculo(50, 10)//usamos o underscore (_) para ignorar soma e subtracao//
//    fmt.Println("Soma: ",soma)
//	fmt.Println("Subtração:",subtracao)
//	fmt.Println("Multiplicação:",multiplica)
//}







