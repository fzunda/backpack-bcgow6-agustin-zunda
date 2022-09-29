package main

import "fmt"

/*
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que tanto los usuarios
como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
	1.Usuario: Nombre, Apellido, Correo, Productos (array de productos).
	2.Producto: Nombre, precio, cantidad.
Se requieren las funciones:
	1.Nuevo producto: recibe nombre y precio, y retorna un producto.
	2.Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
	3.Borrar productos: recibe un usuario, borra los productos del usuario.

*/

type Product struct {
	name   string
	price  float64
	amount int
}

type User struct {
	name     string
	surname  string
	email    string
	products []Product
}

func newProduct(name *string, price *float64) *Product {
	var newProduct = Product{name: *name, price: *price}
	return &newProduct
}

func addNewProduct(user *User, product *Product, amount *int) {
	product.amount = *amount
	user.products = append(user.products, *product)
}

func deleteProduct(user *User) {
	user.products = []Product{}
}

func main() {
	var name1 string = "alfajor"
	var price1 float64 = 1500
	var amount1 int = 10

	/*var n2 string = "Dulce de Leche"
	var p2 float64 = 2000
	var a2 int = 15*/

	newProduct := newProduct(&name1, &price1)
	var newUser = User{name: "Alfredo", surname: "Acosta", email: "acostaA@gmail.com"}
	addNewProduct(&newUser, newProduct, &amount1)
	/*newProduct2 := newProduct(&n2, &p2)
	addNewProduct(&newUser, newProduct2, &a2)*/
	fmt.Println(newUser)

	deleteProduct(&newUser)
	fmt.Println(newUser)
}
