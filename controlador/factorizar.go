package controlador

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/juanmz1606/AnalizadorLL1Proyecto/modelo"
)

func factorizar(gramatica []modelo.Gramatica) []modelo.Gramatica {
	for i := 0; i < len(gramatica); i++ {
		for j := 0; j < len(gramatica[i].Produccion); j++ {
			for k := j + 1; k < len(gramatica[i].Produccion); k++ {
				if gramatica[i].Produccion[j][0] == gramatica[i].Produccion[k][0] {
					temp := PrefijosComunes(gramatica[i].Produccion[j], gramatica[i].Produccion[k])
					//condiccion para aplicar la factorizacion
					if temp != "" {
						noTerminal := fmt.Sprintf("%s'", temp)
						gramatica = agregarProduccion(gramatica, noTerminal, []string{temp + string(noTerminal[1])})
						gramatica[i].Produccion[j] = strings.Replace(gramatica[i].Produccion[j], temp, noTerminal, 1)
						gramatica[i].Produccion[k] = strings.Replace(gramatica[i].Produccion[k], temp, noTerminal, 1)
						// este remover no funciona por ahora
						//gramatica = removerProduccion(gramatica, modelo.Gramatica{Simbolo: gramatica[i].Simbolo, Produccion: []string{}})
						i = 0
						break
					}
				}
			}
		}
	}
	return gramatica
}

func PrefijosComunes(s1, s2 string) string {
	// Obtiene la longitud mínima entre los dos strings
	minLen := len(s1)
	if len(s2) < minLen {
		minLen = len(s2)
	}

	// Busca el índice del primer carácter que difiere entre los dos strings
	i := 0
	for i < minLen && s1[i] == s2[i] {
		i++
	}

	// Devuelve el prefijo común más largo
	return s1[:i]
}

// agrega una produccion a la gramatica
func agregarProduccion(gramatica []modelo.Gramatica, Simbolo string, produccion []string) []modelo.Gramatica {
	for i := range gramatica {
		if gramatica[i].Simbolo == Simbolo {
			gramatica[i].Produccion = append(gramatica[i].Produccion, produccion...)
			return gramatica
		}
	}
	gramatica = append(gramatica, modelo.Gramatica{Simbolo: Simbolo, Produccion: produccion})
	return gramatica
}

// remueve una produccion
func removerProduccion(gramatica []modelo.Gramatica, Simbolo string, produccion []string) []modelo.Gramatica {
	for i, g := range gramatica {
		if g.Simbolo == Simbolo {
			for j, p := range g.Produccion {
				if reflect.DeepEqual(p, produccion) {
					gramatica[i].Produccion = append(g.Produccion[:j], g.Produccion[j+1:]...)
					if len(gramatica[i].Produccion) == 0 {
						gramatica = append(gramatica[:i], gramatica[i+1:]...)
					}
					break
				}
			}
			break
		}
	}
	return gramatica
}
