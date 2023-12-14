package service

import (
	"fmt"

	"github.com/dasagho/htmx-test/db"
	logging "github.com/dasagho/htmx-test/log"
	"github.com/dasagho/htmx-test/models"
)

type ContactService interface {
	AddContact(contact models.ContactInfo) (models.Contact, error)
	GetContact(contactID int) (models.Contact, error)
	UpdateContact(contactID int, updatedInfo models.ContactInfo) (models.Contact, error)
	DeleteContact(contactID int) error
}

type myContactService struct {
	data db.ContactData
}

// AddContact implements ContactService.
func (cs *myContactService) AddContact(contact models.ContactInfo) (models.Contact, error) {
	panic("unimplemented")
}

// DeleteContact implements ContactService.
func (cs *myContactService) DeleteContact(contactID int) error {
	err := cs.data.Delete(contactID)
	if err != nil {
		logging.Error(fmt.Sprintf("Error al eliminar el contacto: %d", contactID))
	}
	return err
}

// GetContact implements ContactService.
func (cs *myContactService) GetContact(contactID int) (models.Contact, error) {
	panic("unimplemented")
}

// UpdateContact implements ContactService.
func (cs *myContactService) UpdateContact(contactID int, updatedInfo models.ContactInfo) (models.Contact, error) {
	panic("unimplemented")
}

func NewMyContactService(data db.ContactData) ContactService {
	return &myContactService{
		data: data,
	}
}
