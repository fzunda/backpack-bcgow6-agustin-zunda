package main

import (
	"fmt"
)

/*Ejercicio 3 - Productos

Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y
retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
	Pequeño: El costo del producto (sin costo adicional)
	Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
	Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
	Crear una estructura “tienda” que guarde una lista de productos.
	Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
	Crear una interface “Producto” que tenga el método “CalcularCosto”
	Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
	Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva
	un Producto.
	Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
	El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
	El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda
*/

const (
	pequeño = "PEQUENO"
	mediano = "MEDIANO"
	grande  = "GRANDE"
)

// struct store and interface
type store struct {
	name          string
	ArrayProducts []typeProduct
}
type Ecommerce interface {
	Total() float64
	Agregar(p typeProduct) bool
}

// struct Product and interface
type Product interface {
	CalcularCosto(Tipe string, price float64) float64
	NewProduct(Tipe string, Name string, price float64) Product
}

type typeProduct struct {
	productTipe string
	name        string
	price       float64
}

func (p *typeProduct) CalcularCosto(Tipe string, price float64) float64 {
	var result float64

	switch Tipe {
	case "PEQUENO":
		result = price
	case "MEDIANO":
		result = price * 1.03
	case "GRANDE":
		result = (price * 1.06) + 2500
	}

	return result
}

func NewProduct(Tipe string, Name string, price float64) typeProduct {
	return typeProduct{productTipe: Tipe, name: Name, price: price}
}

func NewStore(name string, productList ...typeProduct) store {
	var _store = store{name: name}

	if productList != nil {
		for _, value := range productList {
			_store.ArrayProducts = append(_store.ArrayProducts, value)
		}
	}

	return _store
}

func (s *store) Total() float64 {
	var _totalPrice float64

	for _, value := range s.ArrayProducts {
		_totalPrice += value.CalcularCosto(value.productTipe, value.price)
	}

	return _totalPrice
}

func (s *store) Agregar(p typeProduct) {
	s.ArrayProducts = append(s.ArrayProducts, p)
}

func main() {
	_store := NewStore("NuevoCentro")
	fmt.Println("Store: ", _store.name)
	_newProduct := NewProduct(pequeño, "Ojotas", 150)
	_newProduct2 := NewProduct(grande, "Automotor", 250000)
	_store.Agregar(_newProduct)
	_store.Agregar(_newProduct2)
	fmt.Println(_store.ArrayProducts)
	fmt.Println("Precio total: ", _store.Total())
}
