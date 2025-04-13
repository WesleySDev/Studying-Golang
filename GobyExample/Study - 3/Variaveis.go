package main

import "fmt"

func main() {
	//vardeclara 1 ou mais variáveis.
	var a = "initial"
	fmt.Println(a)
	//Você pode declarar várias variáveis ​​de uma só vez.
	var b, c int = 1, 2
	fmt.Println(b, c)
    //Variáveis ​​declaradas sem uma inicialização correspondente são consideradas como valor zero.
	//  Por exemplo, o valor zero para um inté 0
	var d, e, f int
	fmt.Println(d, e, f)
     
	//A :=sintaxe é uma abreviação para declarar e inicializar uma variável, 
	// por exemplo, para var f string = "apple"neste caso. 
	// Esta sintaxe só está disponível dentro de funções.
	g:= "apple"
    fmt.Println(g)


	//Go inferirá o tipo de variáveis ​​inicializadas.
	var h = true 
	fmt.Println(h)
}