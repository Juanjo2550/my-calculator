package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Creamos un struct lo cual es análogo a un objeto
type calc struct {
}

//Le agregamos una función al struct creado
func (calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	switch operador {
	case "+":
		return Suma(operador1, operador2)
	case "-":
		return Resta(operador1, operador2)
	case "*":
		return Multiplicacion(operador1, operador2)
	case "/":
		res, _ := Division(operador1, operador2)
		return res
	default:
		return 0
	}
}

func parsear(entrada string) int {
	operador1, _ := strconv.Atoi(entrada)
	return operador1
}

func leerEntrada(mensaje string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(mensaje)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	operacion := leerEntrada("Ingresa la operación")
	operador := leerEntrada("Ingresa el operador")
	c := calc{}
	res := c.operate(operacion, operador)
	fmt.Printf("%d", res)
}

//Suma Suma 2 números
func Suma(x int, y int) int {
	return (x + y)
}

//Resta Resta 2 números
func Resta(x int, y int) int {
	return x - y
}

//Multiplicacion Multiplica 2 números
func Multiplicacion(x int, y int) int {
	return x * y
}

//Division Divide 2 números
func Division(x int, y int) (int, error) {
	if y == 0 {
		return 0, error(fmt.Errorf("No existe división para cero"))
	}

	return x / y, nil
}
