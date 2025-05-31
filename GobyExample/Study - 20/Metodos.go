// Declara o pacote principal
package main

// Importa o pacote fmt para imprimir na tela
import "fmt"

// Define uma estrutura chamada 'rect' com dois campos: width e height
type rect struct {
    width, height int
}

// Define um método chamado 'area' com um receptor ponteiro (*rect)
// O método calcula a área do retângulo
func (r *rect) area() int {
    return r.width * r.height
}

// Define um método chamado 'perim' com um receptor por valor (rect)
// O método calcula o perímetro do retângulo
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    // Cria uma variável do tipo rect com width = 10 e height = 5
    r := rect{width: 10, height: 5}

    // Chama o método area usando a variável r
    // Mesmo que 'area' tenha receptor ponteiro, o Go faz automaticamente a conversão (&r)
    fmt.Println("area: ", r.area())

    // Chama o método perim usando a variável r
    fmt.Println("perim:", r.perim())

    // Cria um ponteiro para r
    rp := &r

    // Chama o método area usando o ponteiro rp
    fmt.Println("area: ", rp.area())

    // Chama o método perim usando o ponteiro rp
    // Mesmo que perim tenha receptor por valor, o Go também converte automaticamente (*rp)
    fmt.Println("perim:", rp.perim())
}
