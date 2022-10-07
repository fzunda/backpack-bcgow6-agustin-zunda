package main

import "fmt"

/*
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios
y Mantenimientos.
Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice
en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
	1. Productos: nombre, precio, cantidad.
	2. Servicios: nombre, precio, minutos trabajados.
	3. Mantenimiento: nombre, precio.

Se requieren 3 funciones:
	1. Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
	2. Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada,
	   si no llega a trabajar 30 minutos se le cobra
	como si hubiese trabajado media hora).
	3. Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).

*/

type product struct {
	name   string
	price  float64
	amount int
}

type service struct {
	name    string
	price   float64
	minutes int
}

type maintenance struct {
	name  string
	price float64
}

func addProducts(arrayProducts []product) float64 {
	var totalPriceProducts float64
	for _, prod := range arrayProducts {
		priceProduct := prod.price * float64(prod.amount)
		totalPriceProducts += priceProduct
	}
	return totalPriceProducts
}

func addServices(arrayServices []service) float64 {
	var totalPriceservices, priceService float64
	for _, ser := range arrayServices {
		if ser.minutes < 30 {
			priceService = ser.price
		} else {
			cantMinutes := (ser.minutes / 60) * 2
			priceService = ser.price * float64(cantMinutes)
		}
		totalPriceservices += priceService
	}
	return totalPriceservices
}

func addMaintenance(arrayMaintenance []maintenance) float64 {
	var totalMaintenance float64
	for _, val := range arrayMaintenance {
		totalMaintenance += val.price
	}
	return totalMaintenance
}

func main() {
	var listProduct = make([]product, 2)
	listProduct[0] = product{name: "Esponja", price: 250, amount: 10}
	listProduct[1] = product{name: "Lavandina", price: 150, amount: 15}
	fmt.Println("Suma Productos: ", addProducts(listProduct))

	var listServices = make([]service, 2)
	listServices[0] = service{name: "Limpieza de cuarto", price: 50, minutes: 30}
	listServices[1] = service{name: "Limpieza de Casa", price: 50, minutes: 120}
	fmt.Println("Suma Servicios: ", addServices(listServices))

	var listMaintenance = make([]maintenance, 2)
	listMaintenance[0] = maintenance{name: "tareas de carpinteria", price: 2000}
	listMaintenance[1] = maintenance{name: "tareas de alvañileria", price: 2500}
	fmt.Println("Suma Mantenimiento: ", addMaintenance(listMaintenance))
}
