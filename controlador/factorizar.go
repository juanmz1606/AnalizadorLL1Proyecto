package controlador

import (
	"fmt"
	"reflect"
	"strings"
	"github.com/juanmz1606/AnalizadorLL1Proyecto/modelo"
)

func FactorizarGramatica(gramatica []modelo.Produccion) []modelo.Produccion {
	for i := 0; i < len(gramatica); i++ {
		for j := 0; j < len(gramatica[i].Valores); j++ {
			for k := j + 1; k < len(gramatica[i].Valores); k++ {
				if gramatica[i].Valores[j][0] == gramatica[i].Valores[k][0] {
					temp := PrefijosComunes(gramatica[i].Valores[j], gramatica[i].Valores[k])
					fmt.Println(temp)
					//condiccion para aplicar la factorizacion
					if temp != "" && temp != "|"{
							noTerminal := fmt.Sprintf("%s'", temp)
							gramatica = agregarValores(gramatica, noTerminal, []string{temp + string(noTerminal[1])})
						gramatica[i].Valores[j] = strings.Replace(gramatica[i].Valores[j], temp, noTerminal, 1)
						gramatica[i].Valores[k] = strings.Replace(gramatica[i].Valores[k], temp, noTerminal, 1)
						// este remover no funciona por ahora
						//gramatica = removerValores(gramatica, modelo.Produccion{Simbolo: gramatica[i].Simbolo, Valores: []string{}})
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
	for i < minLen && s1 == s2 {
		i++
	}

	// Devuelve el prefijo común más largo
	return s1[:i]
}

// agrega una produccion a la gramatica
func agregarValores(gramatica []modelo.Produccion, Simbolo string, produccion []string) []modelo.Produccion {
	
	for i := range gramatica {
		
		if gramatica[i].Simbolo == Simbolo{
			gramatica[i].Valores = append(gramatica[i].Valores, produccion...)
			return gramatica
		}
	}
	gramatica = append(gramatica, modelo.Produccion{Simbolo: Simbolo, Valores: produccion})
	return gramatica
}

// remueve una produccion
func removerValores(gramatica []modelo.Produccion, Simbolo string, produccion []string) []modelo.Produccion {
	for i, g := range gramatica {
		if g.Simbolo == Simbolo {
			for j, p := range g.Valores {
				if reflect.DeepEqual(p, produccion) {
					gramatica[i].Valores = append(g.Valores[:j], g.Valores[j+1:]...)
					if len(gramatica[i].Valores) == 0 {
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
