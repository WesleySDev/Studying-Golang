// Declara o pacote principal
package main

// Importa pacotes necessários:
// fmt para imprimir na tela
// unicode/utf8 para trabalhar com codificações UTF-8
import (
	"fmt"
	"unicode/utf8"
)

// Função principal: ponto de entrada do programa
func main() {

    // Declara uma constante string com texto em tailandês
    const s = "สวัสดี"

    // Imprime o comprimento da string em bytes
    fmt.Println("Len:", len(s))
    // len conta quantos bytes existem na string.
    // Atenção: caracteres Unicode podem ter mais de 1 byte!

    // Itera byte a byte pela string, imprimindo o valor hexadecimal de cada byte
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()
    // Isso mostra a codificação real em bytes da string UTF-8

    // Conta e imprime quantas runas (caracteres Unicode) existem na string
    fmt.Println("Rune count:", utf8.RuneCountInString(s))
    // utf8.RuneCountInString conta corretamente os caracteres, ignorando quantos bytes cada um ocupa

    // Itera sobre a string, usando o range que já interpreta corretamente como runas
    for idx, runeValue := range s {
        // Imprime a runa no formato Unicode e sua posição (em bytes)
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

    // Mostra outra maneira de percorrer a string, usando DecodeRuneInString
    fmt.Println("\nUsing DecodeRuneInString")
    // i é a posição atual, w será a largura (em bytes) da runa lida
    for i, w := 0, 0; i < len(s); i += w {
        // DecodeRuneInString lê a próxima runa a partir da posição i
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        // Imprime a runa no formato Unicode e a posição onde ela começa
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        // Atualiza w com a quantidade de bytes que a runa ocupa
        w = width

        // Chama uma função para examinar a runa
        examineRune(runeValue)
    }
}

// Função que examina uma runa e imprime se encontrar certos caracteres
func examineRune(r rune) {

    // Se a runa for igual ao caractere 't'
    if r == 't' {
        fmt.Println("found tee")
    // Se a runa for igual ao caractere tailandês 'ส'
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
    // Caso contrário, não faz nada
}

