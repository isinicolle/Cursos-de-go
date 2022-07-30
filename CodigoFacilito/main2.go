package main

import "fmt"

//arreglos
func main() {

	var numeros[5] int

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50

	fmt.Println(numeros)

	numero := [...] int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println(numero)

	monedas := [...] string {"Euro", "Dolar", "Peso", "Libra", "Dolar"}
	fmt.Println(monedas)

	moneda := [...] string {0: "Euro", 1: "Dolar", 2: "Peso", 3: "Libra", 4: "Dolar"}
	fmt.Println(moneda)

	submoneda := moneda[0:2] //slice
	fmt.Println(submoneda)
}


//slice acceso a subsecuencia de elementos para un arreglo
//slice se toma en cuuenta 3 aspectos
//puntero
//longitud
//capacidad

func main() {

	numeros := [] int {1,2,3,4,5,6,7,8,9,10}
	

	//los slice son dinamicos

	numeros = append(numeros, 5)
	numeros = append(numeros, 6)
	numeros = append(numeros, 7)

	nuevoSlice := numeros[0:5]
	otroSlice := numeros[0:3]

	numeros[0] = 100

	fmt.Println(numeros)
	fmt.Println(nuevoSlice)
	fmt.Println(otroSlice)
}

//slice
func main() {
	
	meses:= [] string {"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	longitud := len(meses)
	capacidad := cap(meses)
	fmt.Println(longitud)
	fmt.Println(capacidad)

	//si la estructura s eencuentra en capacidad maxima se genera un nuevo arreglo
	meses = append(meses, "Noviembre")

	fmt.Printf("La longitud es %T\n La capacidad %T\n", longitud, capacidad,meses )

	
}


//Resumen slice: los slice son una referencia , un acceso a una subsecuencia de elementos de un arreglo.

func main() {
	
	numeros := [] int {1,2,3,4,5,6,7,8,9,10}
	
	numeros[0] = 100
	numeros[5] = 600

	inicio := numeros[0:3]
	final := numeros[3:6]

	fmt.Println(numeros)
	fmt.Println(inicio)
	fmt.Println(final)
}



//Make
func main() {

	slice:= make([]int, 3, 5)

	slice[0] = 10
	slice[1] = 20
	slice[2] = 30

	slice = append(slice, 40)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}


//mapas
func main() {

	usuarios := make(map[string] []int)

	usuarios["Isis"] = []init {10,20,30,40,50}

	fmt.Println(usuarios)
}


//Iterar sobre mapas
func main() {
	
	usuarios:= map[int] string{}

	usuarios[1] = "Juan"
	usuarios[2] = "Pedro"
	usuarios[3] = "Maria"
	usuarios[4] = "Juana"

	for id,username:= range usuarios {
		fmt.Println(id,username)
	}
}
