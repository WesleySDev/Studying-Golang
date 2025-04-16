package main

import "fmt"

// A função intSeq retorna outra função que retorna um int
func intSeq() func() int {
 
	i := 0 // Variável local que será "lembrada" pela função retornada


    // Retorna uma função anônima (closure) que incrementa e retorna 'i'
	return func() int{
		i++

	return i
	}
}

func main() {
    // Cria uma nova "sequência" com a função intSeq
	nextInt := intSeq()



  // Cada chamada de nextInt() incrementa e retorna o próximo número
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	 // Cria uma nova sequência independente da anterior
	newints := intSeq()
	fmt.Println(newints()) // 1 (começa do zero de novo)






}