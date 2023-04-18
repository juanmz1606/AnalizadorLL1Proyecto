package main

import "github.com/juanmz1606/AnalizadorLL1Proyecto/modelo"
import "github.com/juanmz1606/AnalizadorLL1Proyecto/controlador"
import "fmt"

/*import ("fmt"
"reflect"
"strconv"
"strings"
"github.com/juanmz1606/AnalizadorLL1Proyecto/modelo")*/

func main() {

	gramatica:= []modelo.Produccion{
		{Simbolo: "E", Valores: []string{"E", "+", "T", "|", "E", "-", "T", "|", "T"}},
	}

	fact:= []modelo.Produccion{
		{Simbolo: "E", Valores: []string{"E'","|", "T"}},
		{Simbolo: "E", Valores: []string{"+","T", "|","-","T"}},
	}

	factorizado := controlador.FactorizarGramatica(gramatica)

	for i := 0; i< len(factorizado);i++{
		fmt.Println(factorizado[i])
	}

	fmt.Println()

	for i := 0; i< len(fact);i++{
		fmt.Println(fact[i])
	}

	/*
	E = E' | T
	E' = +T | -T
	*/
}
/* gramatica := []modelo.Produccion{
		{Simbolo: "S", Valores: []string{"E"}},
		{Simbolo: "E", Valores: []string{"E", "+", "T", "|", "E", "-", "T", "|", "T"}},
		{Simbolo: "T", Valores: []string{"T", "*", "F", "|", "T", "/", "F", "|", "F"}},
		{Simbolo: "F", Valores: []string{"(", "E", ")", "|", "num"}},
	} */


	
