package main 
import (
	"fmt"
	"math"
)
//const declara um valor constante
const s string = "constant"

func main() {
	fmt.Println(s)
//Uma constdeclaração pode aparecer em qualquer lugar onde uma var declaração 
// possa aparecer.
	const n = 500000000

//Expressões constantes realizam operações aritméticas com precisão arbitrária.
	const d = 3e20 / n 
	fmt.Println(d)


//Uma constante numérica não tem tipo até que lhe seja atribuído um, como por meio de uma conversão explícita
	fmt.Println(int64(d))
	
//Um número pode receber um tipo usando-o em um contexto que o exija, como uma atribuição de variável ou chamada de função.
// Por exemplo, aqui math.Sinespera-se um float64
	fmt.Println(math.Sin(n))
}