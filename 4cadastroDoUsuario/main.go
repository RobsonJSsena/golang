         // Cadastro do Usúario//

/*Pelo fato do Scanln só captura a primeira palavra até o espaço (ou seja, "Robson") e 
considera "Silva" como resto da entrada, o que pode atrapalhar as próximas leituras (Email e Idade). 
Isso ocorre porque Scanln lê apenas até o primeiro espaço ou nova linha para cada argumento.
Nesse caso se deseja continuar com o Scanln pode usar o camel case onde as palavras não unidas
sem espaço dessa forma (RobsonNatan) ou caso queira usar os espaços podemos usar o sintaxe
bufio.NewReader com ReadString('\n') (recomendado para strings com espaços)!
*/		 

//Usnado sintaxe bufio.NewReader com ReadString('\n')//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Usuario struct {
	Nome string
	Email string
	Idade int
}

func main(){
	var user Usuario
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("---Cadastro do Usúario---")

	fmt.Println("Digite seu nome:")
	nome, _ := reader.ReadString('\n')
	user.Nome = strings.TrimSpace(nome)
	
	fmt.Println("Digite seu E-mail:")
	email, _ := reader.ReadString('\n')
	user.Email = strings.TrimSpace(email)
	
	fmt.Println("Digite sua idade: ")
	fmt.Scanln(&user.Idade)

fmt.Println("---Usúario Cadastrado---")
fmt.Println("Nome:", user.Nome)	
fmt.Println("E-mail:", user.Email)	
fmt.Println("sua idade:", user.Idade)	
fmt.Println("------------------------------")
}

//--------------------------------------------//
//usamos a sintaxe Scanln
//package main
//
//import (
//  "fmt"
//)
//
//type Usuario struct{
//	Nome string
//	Email string
//	Idade int
//}
//    
//func main()	{
//	var user Usuario
//
//fmt.Println("---Cadastro do Usúario---")
//
//fmt.Println("NameUser: ")
//fmt.Scanln(&user.Nome)
//
//fmt.Println("E-mail: ")
//fmt.Scanln(&user.Email)
//
//fmt.Println("Your age")
//fmt.Scanln(&user.Idade)
//
//fmt.Println("----------------------")
//
//fmt.Println("\nUsúario Cadastrado: ")
//fmt.Println("Nome:", user.Nome)
//fmt.Println("E-mail:", user.Email)
//fmt.Println("Sua Idade:", user.Idade)
//
//} 

