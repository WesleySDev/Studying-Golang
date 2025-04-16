package main

import "fmt"

// A função sum recebe uma quantidade variável de números inteiros (usando ...int)
func sum(nums ...int) {
    // Imprime a lista de números recebida
    fmt.Print(nums, " ")
    
    total := 0

    // Percorre todos os números recebidos e soma eles
    for _, num := range nums {
        total += num
    }

    // Imprime o total da soma
    fmt.Println(total)
}

func main() {

    // Chama a função sum com dois números
    sum(1, 2)

    // Chama a função sum com três números
    sum(1, 2, 3)

    // Cria um slice de inteiros
    nums := []int{1, 2, 3, 4}

    // Chama a função sum passando o slice com o operador ... para "desempacotar" os valores
    sum(nums...)
}
