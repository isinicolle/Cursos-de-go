package main

import "fmt"

//crear estructuras
type User struct {
	Name string
	Email string
}

//metodos
func (self *User) setName(name string) {
	self.Name = name
}

//pseudometodo
func (self *User) getName() string {
	return self.Name
}

func main() {

	/*
	Name := "Isis"
	Email := "isis@gmail.com"

	isis := User{Name, Email}

	fmt.Println(isis.Name)
	fmt.Println(isis.Email)
	*/

	cody := User{Name: "Isis", Email: "isis@gmail.com"}

	cody.setName("Cody")

	fmt.Println(cody)
}

//relacion 1 a 1
type User struct {
	Name string
	Email string
	Active bool
}

type Student struct {
	User User
	Id string
}

func main(){
	cody := User{Name: "Isis", Email: "isis@gmail.com", Active: true}

	uriel := User{Name: "Uriel", Email: "uriel@gmail.com", Active: false}

	codyStudent := Student{User: cody, Id: "123"}

	fmt.Println(uriel)
	fmt.Println(codyStudent.User.Active)
}

//relacion uno a muchos
type Curso struct {
	Titulo string
	Videos []Video
}

type Video struct {
	Titulo string
	Curso Curso
}

func main() {
	video1 := Video{Titulo: "Video 1"}
	video2 := Video{Titulo: "Video 2"}

	videos := []Video{video1, video2}

	curso := Curso{Titulo: "Curso 1", Videos: videos}

	video1.Curso = curso
	video2.Curso = curso

	for _, video := range curso.Videos {
		fmt.Println(video.Titulo)
	}


}


//Enums

type UserType int

const {
	Teacher UserType = 1
	Student UserType = 2
}

type User struct{
	Username string
	Type UserType
}

func main()
{
	isis := User{Username: "Isis", Type: Teacher}
	bryan := User{Username: "Bryan", Type: Student}

	fmt.Println(isis)
	fmt.Println(bryan)

	if isis.Type == Student{
		fmt.Println("El usuario", isis.Username, "es un estudiante")
	}

}

//interfaces
type Animal interface {
	Comer()
	Dormir()
	Cazar()
}

type Perro struct {
	Nombre string
}

func (self Perro) Comer(){
	fmt.Println("El perro", self.Nombre, "come")
}

func (self Perro) Dormir(){
	fmt.Println("El perro", self.Nombre, "duerme")
}

func (self Perro) Cazar(){
	fmt.Println("El perro", self.Nombre, "caza")
}

func ejecutarAcciones(animal Animal){
	animal.Comer()
	animal.Dormir()
	animal.Cazar()
}

func main(){
	
	loki := Perro{Nombre: "Loki"}

	fmt.Println(loki)
}

//Múltiples interfaces
type Animal interface {
	Dormir()
}

type Mascota interface {
	Comer()
}

type Gato struct {
	Nombre string
}

func (self Gato) Comer(){
	fmt.Println("El gato", self.Nombre, "come")
}

func (self Gato) Dormir(){
	fmt.Println("El gato", self.Nombre, "duerme")
}

func funcionParaUnAnimal(animal Animal){
	fmt.Println("El objeto es un Animal")
}

func funcionParaUnaMascota(mascota Mascota){
	fmt.Println("El objeto es una Mascota")
}

func main(){
	
	loki := Gato{Nombre: "Loki"}

	loki.Dormir()
	loki.Comer()
}

//Interfaces vacias
func mostrarVariable(objeto interface{}){
	fmt.Printf("El valor de la variable es: %v",objeto)
}

func suma(numero1, numero2 int) int{
	return numero1 + numero2
}

type User struct {
	Name string
}

func main(){
	usuario := User {Nombre: "Isis"}

	mostrarVariable(usuario)
}

