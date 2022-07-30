package main

import "os"
import "fmt"
import "bufio"
import "strings"

func crearUsuario() {
	fmt.Println("Usuario creado") 
}

func listarUsuarios() {
	fmt.Println("Lista de usuarios")
}

func actualizarUsuario() {
	fmt.Println("Usuario actualizado")
}

func eliminarUsuario() {
	fmt.Println("Usuario eliminado")
}

func readLine() string{
	
	if  option, err := reader.ReadString('\n'); err != nil {
		panic("Error al leer la opcion")
	}

	return strings.TrimSuffix(option, "\n")
}

func main() {

	var option string

	reader = bufio.NewReader(os.Stdin)


	for {
	
	fmt.Println("Bienvenido al sistema de usuarios")
	fmt.Println("1. Crear usuario")
	fmt.Println("2. Listar usuarios")
	fmt.Println("3. Actualizar usuario")
	fmt.Println("4. Eliminar usuario")

	fmt.Println("Ingrese una opcion VALIDA: ")

	option = readLine()

	switch option {
	case "1":
		crearUsuario()
	case "2":
		listarUsuarios()
	case "3":
		actualizarUsuario()
	case "4":
		eliminarUsuario()
	default:
		fmt.Println("Opcion no valida")
		}
	}
}


//Que depresion me dio seguir un proyecto con programacion estructurada ajajaj 
//este es el ultimo commit que hago al curso de codigofacilito mejor hare un proyecto real y luego hare el curso de platzi.