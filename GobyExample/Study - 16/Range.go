package main

import "fmt"

func main() {
    // Criamos uma slice (fatia) de inteiros com 3 valores: 2, 3 e 4
    nums := []int{2, 3, 4}

    // Inicializamos a variável "sum" com 0
    sum := 0

    // Usamos "range" para percorrer cada valor da slice
    // O "_" ignora o índice; só usamos o valor "num"
    for _, num := range nums {
        sum += num // somamos os valores um por um
    }

    // Imprime o resultado da soma
    fmt.Println("sum:", sum)

    // Agora vamos usar "range" novamente, mas dessa vez usando o índice também
    for i, num := range nums {
        // Se o número atual for 3, imprime o índice onde ele está
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // Criamos um mapa onde as chaves são strings e os valores também
    kvs := map[string]string{"a": "apple", "b": "banana"}

    // Usamos "range" para percorrer o mapa
    // Cada iteração retorna a chave "k" e o valor "v"
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v) // imprime: chave -> valor
    }

    // Podemos também iterar apenas pelas chaves do mapa
    for k := range kvs {
        fmt.Println("key:", k) // imprime só a chave
    }

    // Finalmente, usamos "range" para percorrer uma string
    // Ele retorna o índice do byte e o valor da "rune" (caractere Unicode)
    for i, c := range "go" {
        fmt.Println(i, c) // imprime o índice e o valor Unicode de cada caractere
    }
}
