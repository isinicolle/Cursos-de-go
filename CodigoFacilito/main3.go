package main

import "fmt"

func main() {
	fmt.Println("Hola mundo")

	//condicionales
	if true {
		fmt.Println("Si")
	} else {
		fmt.Println("No")
	}

	//switch
	switch "Isis" {
	case "Isis":
		fmt.Println("Isis")
	case "Juan":
		fmt.Println("Juan")
	default:
		fmt.Println("No se")
		}

	//ciclos
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//ciclo do-while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//ciclo while
	i = 0
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	//ciclo foreach
	var nombres []string
	nombres = []string{"Isis", "Juan", "Pedro"}
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}

	//continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	//break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	
	//funcion panic
	panic("Error en el programa")

	//obtener valores de un mapa
	 usuarios := make(map[string]string)

	usuarios["Isis"] = "isinicolle@hotmail.com"

	if correo, ok := usuarios["Isis"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No se encontro el correo")
	}

	
}