package main

import "fmt"

// zeroval recebe um inteiro como valor (ou seja, uma cópia do valor original).
func zeroval(ival int) {
    ival = 0 // Altera a cópia local de ival. Isso NÃO afeta a variável original.
}

// zeroptr recebe um ponteiro para int.
// Como é um ponteiro, ele consegue acessar e modificar o valor original.
func zeroptr(iptr *int) {
    *iptr = 0 // *iptr acessa o valor no endereço de memória e o altera para 0.
}

func main() {
    i := 1 // Cria uma variável inteira 'i' com valor inicial 1.
    fmt.Println("initial:", i) // Imprime: initial: 1

    zeroval(i) // Chama a função zeroval passando uma cópia de 'i'.
    fmt.Println("zeroval:", i) // Imprime: zeroval: 1 — o valor de 'i' não mudou.

    zeroptr(&i) // Chama a função zeroptr passando o endereço da variável 'i'.
    fmt.Println("zeroptr:", i) // Imprime: zeroptr: 0 — o valor de 'i' foi alterado pela função.

    fmt.Println("pointer:", &i) // Imprime o endereço de memória de 'i'.
}
