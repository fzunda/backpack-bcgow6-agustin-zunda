package main

import (
	"github.com/fzunda/backpack-bcgow6-agustin-zunda/Go_Bases/hackaton-go-bases-main/internal/file"
	"github.com/fzunda/backpack-bcgow6-agustin-zunda/Go_Bases/hackaton-go-bases-main/internal/service"
)

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
	tickets, err := file.Read()
}
