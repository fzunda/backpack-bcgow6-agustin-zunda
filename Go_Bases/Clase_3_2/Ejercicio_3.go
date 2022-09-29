package main

/*
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos.
Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
	1. Productos: nombre, precio, cantidad.
	2. Servicios: nombre, precio, minutos trabajados.
	3. Mantenimiento: nombre, precio.

Se requieren 3 funciones:
	1. Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
	2. Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra
	como si hubiese trabajado media hora).
	3. Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).

*/

type Product struct {
	name   string
	price  float64
	amount int
}

type Service struct {
	name    string
	price   float64
	minutes int
}

type maintenance struct {
	name  string
	price float64
}

func addProducts(arrayProducts ...Product) float64 {
	var totalPriceProducts float64
	for _, prod := range arrayProducts {
		priceProduct := prod.price * float64(prod.amount)
		totalPriceProducts += priceProduct
	}
	return totalPriceProducts
}

func addServices(arrayServices ...Service) float64 {
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

func main() {
	//var arrayProducts []Product = primes[:]
}
