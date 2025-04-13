package main // Declara o pacote principal, necessário para que o programa seja executável

import "fmt" // Importa o pacote fmt, utilizado para formatação de entrada e saída

func main() { // Função principal, ponto de entrada do programa

    var a [5]int // Declara um array 'a' com 5 elementos do tipo int, inicializados com zero
    fmt.Println("emp:", a) // Imprime o array 'a' vazio (todos os elementos são 0)

    a[4] = 100 // Atribui o valor 100 ao quinto elemento do array 'a' (índice 4)
    fmt.Println("set:", a) // Imprime o array 'a' após a modificação, mostrando o último elemento como 100
    fmt.Println("get:", a[4]) // Acessa e imprime o valor do quinto elemento do array 'a', que é 100

    fmt.Println("len:", len(a)) // Imprime o comprimento do array 'a', que é 5

    b := [5]int{1, 2, 3, 4, 5} // Declara e inicializa o array 'b' com os valores 1, 2, 3, 4 e 5
    fmt.Println("dcl:", b) // Imprime o conteúdo do array 'b'

    b = [...]int{1, 2, 3, 4, 5} // Outra forma de declarar e inicializar o array 'b', permitindo que o compilador determine o tamanho
    fmt.Println("dcl:", b) // Imprime o conteúdo do array 'b'

    b = [...]int{100, 3: 400, 500} // Inicializa o array 'b' com 5 elementos: o primeiro com 100, o quarto com 400 e o último com 500
    fmt.Println("idx:", b) // Imprime o conteúdo do array 'b'

    var twoD [2][3]int // Declara uma matriz bidimensional 'twoD' com 2 linhas e 3 colunas, inicializada com zeros
    for i := 0; i < 2; i++ { // Inicia um loop para percorrer as linhas da matriz
        for j := 0; j < 3; j++ { // Inicia um loop para percorrer as colunas da matriz
            twoD[i][j] = i + j // Atribui à posição [i][j] da matriz a soma dos índices i e j
        }
    }
    fmt.Println("2d: ", twoD) // Imprime o conteúdo da matriz 'twoD'

    twoD = [2][3]int{ // Re-inicializa a matriz 'twoD' com valores específicos
        {1, 2, 3}, // Primeira linha da matriz
        {4, 5, 6}, // Segunda linha da matriz
    }
    fmt.Println("2d: ", twoD) // Imprime o conteúdo atualizado da matriz 'twoD'
}
