package file

import "github.com/fzunda/backpack-bcgow6-agustin-zunda/Go_Bases/hackaton-go-bases-main/internal/service"

type File struct {
	path string
}

func (f *File) Read() ([]service.Ticket, error) {
	return nil, nil
}

func (f *File) Write(service.Ticket) error {
	return nil
}
