package main

import "fmt"

func main () {
	if 7%2 == 0 {
		fmt.Println(" 7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//Você pode ter uma ifdeclaração sem um else.
   if 8%4 == 0 {
		fmt.Println("8 is divsible by 4")
	}

	//Operadores lógicos como && e || geralmente são úteis em condições.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	//Go permite que você declare variáveis dentro do próprio if, antes da condição.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")	
	}else {
		fmt.Println(num, "has multiple digits")
	}	
}