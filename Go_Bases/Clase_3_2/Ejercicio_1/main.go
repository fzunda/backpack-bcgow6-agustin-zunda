package main

import "fmt"

/*
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. Para optimizar
y ahorrar memoria requieren que la estructura de usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones.
La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:
Cambiar nombre: me permite cambiar el nombre y apellido.
Cambiar edad: me permite cambiar la edad.
Cambiar correo: me permite cambiar el correo.
Cambiar contraseña: me permite cambiar la contraseña.

*/

type UserData struct {
	nombre     string
	apellido   string
	Edad       int
	correo     string
	contraseña string
}

// formas de hacer un setter --> se usan punteros
func (u *UserData) rename(name *string) {
	u.nombre = *name
}

func (u *UserData) changeAge(age *int) {
	u.Edad = *age
}

func (u *UserData) changeEmail(email *string) {
	u.correo = *email
}

func (u *UserData) changePassword(pass *string) {
	u.contraseña = *pass
}

func main() {
	//punteros
	var name string = "Jose"
	var namePointer *string = &name
	//fmt.Println(namePointer)
	var surname string = "Gomez"
	var surnamePointer *string = &surname
	//fmt.Println(surnamePointer)
	var age int = 45
	var agePointer *int = &age
	//fmt.Println(agePointer)
	var email string = "jgomez@gmail.com"
	var emailPointer *string = &email
	//fmt.Println(emailPointer)
	var password string = "12345Gomez"
	var passwordPointer *string = &password
	//fmt.Println(passwordPointer)

	var user = UserData{nombre: *namePointer, apellido: *surnamePointer, Edad: *agePointer, correo: *emailPointer, contraseña: *passwordPointer}
	fmt.Println(user)
	var newName string = "Alberto"
	user.rename(&newName)
	fmt.Println(user)
	var newAge int = 50
	user.changeAge(&newAge)
	fmt.Println(user)
	var newPassword string = "555GomezA"
	user.changePassword(&newPassword)
	fmt.Println(user)
	var newEmail string = "Agomez@gmail.com"
	user.changeEmail(&newEmail)
	fmt.Println(user)
}
