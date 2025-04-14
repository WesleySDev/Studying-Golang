package main

import (
	"fmt"  // Importa o pacote fmt para imprimir valores no terminal
	"maps" // Importa o pacote maps (Go 1.21+), que fornece funções utilitárias para trabalhar com mapas
)

func main() {

    // Cria um mapa com chave do tipo string e valor do tipo int
    m := make(map[string]int)

    // Atribui o valor 7 à chave "k1"
    m["k1"] = 7
    // Atribui o valor 13 à chave "k2"
    m["k2"] = 13

    // Imprime o mapa completo
    fmt.Println("map:", m)

    // Acessa o valor associado à chave "k1" e armazena na variável v1
    v1 := m["k1"]
    fmt.Println("v1:", v1) // Imprime o valor de "k1"

    // Tenta acessar uma chave que não existe ("k3") — o valor retornado será 0 (valor zero de int)
    v3 := m["k3"]
    fmt.Println("v3:", v3) // Imprime 0

    // Imprime o tamanho atual do mapa (número de pares chave-valor)
    fmt.Println("len:", len(m))

    // Remove a chave "k2" do mapa
    delete(m, "k2")
    fmt.Println("map:", m) // Imprime o mapa após a remoção

    // Limpa completamente o mapa (remove todos os elementos)
    clear(m)
    fmt.Println("map:", m) // Mapa agora está vazio

    // Verifica se a chave "k2" ainda existe no mapa (mesmo depois de limpar)
    _, prs := m["k2"]
    fmt.Println("prs:", prs) // Imprime false, pois a chave não existe

    // Declara e inicializa um mapa com dois pares chave-valor
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n) // Imprime o mapa n

    // Declara outro mapa com os mesmos valores que o anterior
    n2 := map[string]int{"foo": 1, "bar": 2}

    // Compara os mapas usando a função maps.Equal
    if maps.Equal(n, n2) {
        fmt.Println("n == n2") // Se os mapas forem iguais, imprime a mensagem
    }
}
