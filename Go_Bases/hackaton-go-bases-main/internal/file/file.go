package file

import (
	"fmt"
	"os"
	"strings"

	"github.comfzunda/backpack-bcgow6-agustin-zunda/Go_Bases/hackaton-go-bases-main/internal/service"
)

type File struct {
	path string
}

func (f *File) Read() ([]service.Ticket, error) {
	//var sliceTickets = []service.Ticket{}
	data, err := os.ReadFile("./tickets.csv")
	if err != nil {
		return nil, err
	}

	file := string(data)
	rows := strings.Split(file, "\n")

	var new_rows string = fmt.Sprintf("\nEjemplo: \n\n")
	new_rows += fmt.Sprintf("ID\tNombre\tApellido\tDni\tSalario\n")

	for _, value := range rows {
		row := strings.Replace(value, ",", "\t", -1)
		new_rows += fmt.Sprintf("%v\n", row)
	}
	fmt.Println(new_rows)

	return nil, nil
}

func (f *File) Write(service.Ticket) error {
	return nil
}
