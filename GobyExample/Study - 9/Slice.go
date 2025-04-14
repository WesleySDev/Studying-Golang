package main

import (
	"fmt"    // Importa o pacote fmt para saída formatada
	"slices" // Importa o pacote slices (disponível a partir do Go 1.21) para funções úteis com slices
)

func main() {

    // Declara uma slice de strings não inicializada (valor zero de slice é nil)
    var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0) // Imprime a slice (nil), se ela é nil, e se o comprimento é 0

    // Inicializa a slice com tamanho 3 usando make (capacidade e comprimento = 3)
    s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) // Imprime slice vazia, tamanho e capacidade

    // Atribui valores às posições da slice
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)       // Imprime a slice atualizada
    fmt.Println("get:", s[2])    // Imprime o valor do índice 2 (terceiro elemento)

    fmt.Println("len:", len(s))  // Imprime o comprimento da slice

    // Adiciona elementos à slice usando append (aumenta automaticamente o tamanho e a capacidade se necessário)
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s) // Imprime a slice após os appends

    // Cria uma nova slice com o mesmo comprimento que s
    c := make([]string, len(s))
    copy(c, s) // Copia os elementos de s para c
    fmt.Println("cpy:", c) // Imprime a slice copiada

    // Cria uma sub-slice de s, pegando elementos do índice 2 até 4 (não inclui 5)
    l := s[2:5]
    fmt.Println("sl1:", l) // Imprime sub-slice contendo os elementos dos índices 2, 3, 4

    // Sub-slice de s do início até o índice 4 (não inclui 5)
    l = s[:5]
    fmt.Println("sl2:", l) // Imprime elementos do índice 0 até 4

    // Sub-slice de s do índice 2 até o final
    l = s[2:]
    fmt.Println("sl3:", l) // Imprime do índice 2 até o fim

    // Declara e inicializa uma slice com 3 strings
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t) // Imprime a slice t

    // Cria outra slice com os mesmos elementos
    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) { // Compara se t e t2 têm os mesmos elementos na mesma ordem
        fmt.Println("t == t2") // Imprime se forem iguais
    }

    // Cria uma slice de slices de inteiros (matriz 2D com 3 linhas)
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1                  // Define o tamanho da linha interna como i + 1
        twoD[i] = make([]int, innerLen)   // Inicializa a linha interna
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j            // Atribui valores somando os índices i e j
        }
    }
    fmt.Println("2d: ", twoD) // Imprime a matriz 2D resultante
}
