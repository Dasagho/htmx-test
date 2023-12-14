package models

type Contact struct {
	Id       int
	Name     string
	Telefono int
}

func NewContact(id, tel int, name string) Contact {
	return Contact{
		Id:       id,
		Name:     name,
		Telefono: tel,
	}
}

type ContactInfo struct {
	Name     string
	Telefono int
}
