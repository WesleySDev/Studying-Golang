package main

import "fmt" // Importa o pacote fmt para imprimir no console

// Define a função 'plus' que recebe dois inteiros e retorna a soma deles
func plus(a int, b int) int {
    return a + b // Retorna a soma de a e b
}

// Define a função 'plusPlus' que recebe três inteiros e retorna a soma deles
func plusPlus(a, b, c int) int {
    return a + b + c // Retorna a soma de a, b e c
}

func main() {

    // Chama a função plus com os valores 1 e 2, e armazena o resultado em 'res'
    res := plus(1, 2)
    fmt.Println("1+2 =", res) // Imprime o resultado: "1+2 = 3"

    // Chama a função plusPlus com os valores 1, 2 e 3, e armazena o resultado em 'res'
    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res) // Imprime o resultado: "1+2+3 = 6"
}
