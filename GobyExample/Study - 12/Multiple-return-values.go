package main

import "fmt"

// A função vals retorna dois valores inteiros: 3 e 7
func vals() (int, int) {
    return 3, 7
}

func main() {

    // Aqui, os dois valores retornados por vals() são atribuídos a 'a' e 'b'
    a, b := vals()
    fmt.Println(a) // Imprime o valor de 'a' (3)
    fmt.Println(b) // Imprime o valor de 'b' (7)

    // Aqui, o primeiro valor retornado é ignorado usando o '_' e o segundo é atribuído a 'c'
    _, c := vals()
    fmt.Println(c) // Imprime o valor de 'c' (7)
}
