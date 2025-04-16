package main

import "fmt"

// Função recursiva que calcula o fatorial de um número
func fact(n int) int {
    // Caso base: fatorial de 0 é 1
    if n == 0 {
        return 1
    }
    // Chamada recursiva: n * fatorial de (n - 1)
    return n * fact(n-1)
}

func main() {
    // Imprime o fatorial de 7 (7 * 6 * 5 * ... * 1 = 5040)
    fmt.Println(fact(7))

    // Declara uma variável de função chamada 'fib' que recebe int e retorna int
    var fib func(n int) int

    // Define a função de Fibonacci recursiva
    fib = func(n int) int {
        // Caso base: se n for 0 ou 1, retorna o próprio n
        if n < 2 {
            return n
        }
        // Chamada recursiva: soma dos dois anteriores na sequência
        return fib(n-1) + fib(n-2)
    }

    // Imprime o 7º número da sequência de Fibonacci (0, 1, 1, 2, 3, 5, 8, 13...)
    fmt.Println(fib(7)) // Resultado: 13
}
