package service

import (
	"testing"

	"github.com/dasagho/htmx-test/models"
)

// MockContactData es una implementación de prueba de ContactData
type MockContactData struct{}

// Delete implements db.ContactData.
func (*MockContactData) Delete(contactID int) error {
	panic("unimplemented")
}

// FindByID implements db.ContactData.
func (*MockContactData) FindByID(contactID int) (models.Contact, error) {
	panic("unimplemented")
}

// Update implements db.ContactData.
func (*MockContactData) Update(contactID int, contact models.Contact) error {
	panic("unimplemented")
}

func (m *MockContactData) Save(contact models.Contact) error {
	return nil // Simula que no hay errores
}

func TestAddContact(t *testing.T) {
	mockData := new(MockContactData)
	service := NewMyContactService(mockData)
	contact := models.ContactInfo{Name: "Test Contact", Telefono: 123456789}

	result, err := service.AddContact(contact)

	if err != nil {
		t.Errorf("Se esperaba que no hubiera error, pero se obtuvo: %v", err)
	}

	if result.Name != contact.Name || result.Telefono != contact.Telefono {
		t.Errorf("Se esperaba %v, pero se obtuvo %v", contact, result)
	}
}
