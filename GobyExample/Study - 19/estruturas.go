// Declara o pacote principal
package main

// Importa o pacote fmt para imprimir na tela
import "fmt"

// Define uma estrutura (struct) chamada 'person' com dois campos: name e age
type person struct {
    name string
    age  int
}

// Função que cria e retorna um ponteiro para uma 'person'
func newPerson(name string) *person {
    // Cria uma variável do tipo person, inicializando o campo name
    p := person{name: name}
    // Define o campo age como 42
    p.age = 42
    // Retorna o endereço da variável p (um ponteiro para person)
    return &p
}

func main() {
    // Cria e imprime uma variável do tipo person, inicializada com name = "Bob" e age = 20
    fmt.Println(person{"Bob", 20})

    // Cria e imprime uma variável do tipo person, inicializada usando campos nomeados
    fmt.Println(person{name: "Alice", age: 30})

    // Cria e imprime uma variável do tipo person, inicializando apenas o campo name
    // O campo age recebe o valor zero padrão (0)
    fmt.Println(person{name: "Fred"})

    // Cria uma variável do tipo person com campos nomeados e imprime o ponteiro para ela
    fmt.Println(&person{name: "Ann", age: 40})

    // Chama a função newPerson e imprime o ponteiro retornado
    fmt.Println(newPerson("Jon"))

    // Cria uma variável do tipo person
    s := person{name: "Sean", age: 50}
    // Acessa e imprime o campo name da variável
    fmt.Println(s.name)

    // Cria um ponteiro para a variável s
    sp := &s
    // Acessa e imprime o campo age através do ponteiro (Go automaticamente desreferencia)
    fmt.Println(sp.age)

    // Modifica o campo age da struct através do ponteiro
    sp.age = 51
    // Imprime o campo age após a modificação
    fmt.Println(sp.age)

    // Cria e inicializa uma struct anônima (sem nome), com dois campos: name e isGood
    dog := struct {
        name   string
        isGood bool
    }{
        "Rex",  // Inicializa o campo name com "Rex"
        true,   // Inicializa o campo isGood com true
    }
    // Imprime a struct anônima
    fmt.Println(dog)
}
