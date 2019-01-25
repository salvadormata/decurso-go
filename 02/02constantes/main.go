//Comentario de linea
/*
  comentario de bloque
*/

package main

import "fmt"

func main() {
	//const nombre = "Salvador"
	//fmt.Println(nombre)

	var a int
	var b int8
	n := "shava"
	p := "mata"

	a = 121212
	b = 5
	//casting
	c := a + int(b)
	fmt.Println(c)
	fmt.Printf("Hola %s %s como estafas??\n", n, p)
	fmt.Printf("b es de tipo: %T\n", b)
	//prioridad aritmetica
	//+ - * /
	// () % * / + -
	d := 6 + 6*(6-6)
	fmt.Println(d)

}
