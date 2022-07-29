package main

import "fmt"
import "reflect"


//imprimir una cadena de caracteres
func main() {
	fmt.Println("Hello World")
}


//usar go run main.go para ejecutar el programa
//go build main.go para devuelve un archivo ejecutable


//declaracion de variables sin omitir tipo de dato
func main() {
	var nombre string = "Isis Zapata"
	var edad int = 21
	var peso float32 = 45.5
	var casado bool = false
	fmt.Println(nombre, edad, peso, casado)
}


//declaracion de variables omitiendo tipo de dato
//el compilador intuye el tipo de dato al asignarle un valor
func main() {
	nombre := "Isis Zapata"
	edad := 21
	peso := 45.5
	var casado = false
	fmt.Println("El nombre es:", nombre,"La edad es:", edad, "El peso es:", peso, casado)
}


//declarar multiples variables
func main() {
	var nombre, apellido string = "Isis", "Zapata"
	edad, peso, casado := 21, 45.5, false
	fmt.Println("El nombre es:", nombre, "El apellido es:", apellido, "La edad es:", edad, "El peso es:", peso, "El casado es:", casado)
}

//constantes en Go se declara fuera de una funcion principal
const Curso string = "Codigo Facilito"

func main() {

 fmt.Println(Curso)

}



//operadores relacionales
func main() {


	>
	<
	>=
	<=
	==
	!=
	

	var numero1, numero2 int = 10, 20
	var resultado bool
	resultado = numero1 > numero2
	fmt.Println(resultado)
}



//operadores logicos
func main() {
	
	
	&&
	||
	!
	

	//&& todas deben  ser verdaderas para devolver true
	resultado := 20 == 20 && 30 == 30 && 20>40
	fmt.Println(resultado)

	//|| una debe ser verdadera para devolver true
	resultado := 20 == 20 || 30 == 30 || 20>40
	fmt.Println(resultado)
}


//secuencia de valores


const Domingo = 0
const Lunes = 1
const Martes = 2
const Miercoles = 3
const Jueves = 4
const Viernes = 5
const Sabado = 6



const (
	Domingo = 0
	Lunes = 1
	Martes = 2
	Miercoles = 3
	Jueves = 4
	Viernes = 5
	Sabado = 6
)



const (
	Domingo = iota
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {

	fmt.Println(Domingo)
	fmt.Println(Lunes)
	fmt.Println(Martes)
	fmt.Println(Miercoles)
	fmt.Println(Jueves)
	fmt.Println(Viernes)
	fmt.Println(Sabado)
}


//strings
func main() {
	//declarando variable tipo string
	var nombre string = "Isis Zapata"
	var apellido string = "Zapata"
	var nombreCompleto string = nombre + " " + apellido
	fmt.Println(nombreCompleto)

	//tipo de dato haciendo que el compilador lo intuya
	name:= "Isis Zapata"
	fmt.Println(name)

	longitud := len(name)
	fmt.Println(longitud)

	primerCaracter := name[0]
	fmt.Println(primerCaracter)

	//tipo de dato de la variable
	fmt.Println(reflect.TypeOf(name))

	//ver el primer caractr en letra
	fmt.Printf("%c\n",primerCaracter)

	//tener ultimo caracter
	ultimoCaracter := name[longitud-1]
	fmt.Printf("%c\n",ultimoCaracter)
}


//lectura de valores
func main() {

	var nombre string
	var edad int
	var peso float32

	fmt.Print("Ingrese su nombre: ")
	fmt.Scan("%s",&nombre)

	fmt.Print("Ingrese su edad: ")
	fmt.Scan("%d",&edad)

	fmt.Print("Ingrese su peso: ")
	fmt.Scan("%f",&peso)

	fmt.Printf("Hola %s\n tu edad es %d\n con tu peso: %f\n",nombre, edad,peso)
}

