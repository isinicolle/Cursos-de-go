package main

import "fmt"

type Operacion func(balance,cantidad int) int

//funciones

func saluda(username string){
	fmt.Println("Hola", username)

}

//suma de dos numeros
func suma(num1 int, num2 int) int {
	return num1 + num2
}

//multiplicacion de dos numeros
func multiplicacion(num1 int, num2 int) (int, string) {
	return num1 * num2, "La multiplicacion es: "
}

//funciones como argumento
func ejecutarOperacion(funcion Operacion, balance, cantidad int)  {
	
	fmt.Println("Antes de la operacion")

	resultado := funcion(balance, cantidad)
	fmt.Println("El resultado es: ", resultado)

	fmt.Println("Despues de la operacion")
}


func main() {

	saluda("Isis")
	fmt.Println(suma(2,3))
	
	//si no queremos guardar mensaje en ninguna variable
	//resultado, _ := multiplicacion(2,3)

	resultado, mensaje := multiplicacion(2,3)
	fmt.Println(resultado, mensaje)

	//funciones anonimas
	funcionAnonima := func (username string) string {
	
		message :=  fmt.Sprintf("Hola %s", username)

		return message
	}

	mensajeFinal := funcionAnonima("Isis")

	fmt.Println(mensajeFinal)

	ejecutarOperacion(suma, 100, 50)
}



type Operacion func(balance, cantidad int) int


func crearOperacion(tipo string) Operacion {
	if tipo == "suma" {
		return func(balance, cantidad int) int {
			return balance + cantidad
		}
	} else if tipo == "resta" {
		return func(balance, cantidad int) int {
			return balance - cantidad
		}
	} else if tipo == "multiplicacion" {
		return func(balance, cantidad int) int {
			return balance * cantidad
		}
	} else if tipo == "division" {
		return func(balance, cantidad int) int {
			return balance / cantidad
		}
	} else {
		return nil
	}
}

func main() {
	operacion := crearOperacion("suma")
	fmt.Println(operacion(100, 50))
    operacion(100, 50)
}


//bloques y funciones
//una variable vive en un bloque

//Variadic functions
//es una funcion que puede recibir una N cantidad de argumentos que nosotros deseemos
func promedio(calificaciones ...int) int{

	var promedio, suma int

	for _, calificacion := range calificaciones {
		suma += calificacion
	}

	promedio = suma / len(calificaciones)

	return promedio
}

func main() {

	resultado := promedio(10,20,30,40,50)
	fmt.Println("Promedio: ",resultado)
}


//punteros

func modificarValor(parametro *string) {
	*parametro = "Profesional de Go"
}

func main() {
	
	var curso = "Go"

	fmt.Println("Curso: ", curso)

	modificarValor(&curso) //referencia

	fmt.Println("Curso modificado: ", curso)
}


//funciones recursivas

func factorial(numero int) int{
	if numero == 0 {
		return 1
	} else {
		return numero * factorial(numero - 1)
	}
}

func main() {

	resultado := factorial(5)
	fmt.Println("Factorial: ", resultado)
}



//funciones como valor
func raizCuadrada(numero int) int {
	return numero * numero
}

func main() {
	var miFuncion func(n int) int
	
	if miFuncion == nil {
		miFuncion = raizCuadrada
	} 

	resultado := miFuncion(5)
	fmt.Println("Resultado: ", resultado)
}


import "errors"

func division(dividendo,divisor int) (int,error){
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por 0")
	} else {
		return dividendo / divisor, nil
	}
	//nil define la ausencia de un valor.
} 

func main(){

	if resultado, err := division(10,0); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Resultado: ", resultado)
	}
}

//funcio recover
func main() {
	defer func() {
		if err := recover(); err != nil { //recover() devuelve el error que se produjo en la funcion que lo llamo y lo guarda en la variable error
			fmt.Println("Error: ", err)
		}
	}()

	panic("Error en el programa")
} //panic es una funcion que se ejecuta en el momento en que se produce un error en el programa