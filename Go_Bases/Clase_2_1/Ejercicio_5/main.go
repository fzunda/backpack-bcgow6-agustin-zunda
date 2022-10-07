package main

import "fmt"

/*
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el
momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos
más animales que refugiar.

perro necesitan 10 kg de alimento
gato 5 kg
Hamster 250 gramos.
Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal
especificado y que retorne una función y un mensaje (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de
animal especificado.
*/

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func funcDog(dogs int) float64 {
	var foodDog float64 = float64(dogs) * 10
	return foodDog
}

func funcCat(cats int) float64 {
	var foodCats float64 = float64(cats) * 5
	return foodCats
}

func funcHamster(hamsters int) float64 {
	var foodhamsters float64 = float64(hamsters) * 0.150
	return foodhamsters
}

func funcTarantula(tarantulas int) float64 {
	var foodTarantulas float64 = float64(tarantulas) * 0.250
	return foodTarantulas
}

func funcAnimal(animal string) (func(int) float64, string) {
	switch animal {
	case "dog":
		return funcDog, ""
	case "cat":
		return funcCat, ""
	case "hamster":
		return funcHamster, ""
	case "tarantula":
		return funcTarantula, ""
	}

	return nil, "El animal no se encuentra registrado.."
}

func main() {
	test1, err := funcAnimal("dog")
	if err == "" {
		fmt.Println("Comida Perros en Kg.: ", test1(5))
	} else {
		fmt.Println(err)
	}
	test2, err := funcAnimal("cat")
	if err == "" {
		fmt.Println("Comida Gatos en Kg.: ", test2(5))
	} else {
		fmt.Println(err)
	}
	test3, err := funcAnimal("hamster")
	if err == "" {
		fmt.Println("Comida Hamsters en Kg.: ", test3(5))
	} else {
		fmt.Println(err)
	}
	test4, err := funcAnimal("tarantula")
	if err == "" {
		fmt.Println("Comida Tarantulas en Kg.: ", test4(5))
	} else {
		fmt.Println(err)
	}
	test5, err := funcAnimal("rabbit")
	if err == "" {
		fmt.Println("Comida Conejos en Kg.: ", test5(5))
	} else {
		fmt.Println(err)
	}
}
