package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dasagho/htmx-test/models"
)

// MockContactService es una implementación de prueba de ContactService
type MockContactService struct{}

// DeleteContact implements service.ContactService.
func (*MockContactService) DeleteContact(contactID int) error {
	panic("unimplemented")
}

// GetContact implements service.ContactService.
func (*MockContactService) GetContact(contactID int) (models.Contact, error) {
	panic("unimplemented")
}

// UpdateContact implements service.ContactService.
func (*MockContactService) UpdateContact(contactID int, updatedInfo models.ContactInfo) (models.Contact, error) {
	panic("unimplemented")
}

func (m *MockContactService) AddContact(contact models.ContactInfo) (models.Contact, error) {
	return models.Contact{
		Name:     contact.Name,
		Telefono: contact.Telefono,
	}, nil // Simula que el contacto se agrega sin errores
}

func TestContactHandler_ServeHTTP(t *testing.T) {
	mockService := new(MockContactService)
	handler := NewContactHandler(mockService)

	req, _ := http.NewRequest("POST", "/contact", nil) // TODO implement send body contact data
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler devolvió un código incorrecto: obtuvo %v esperaba %v",
			status, http.StatusOK)
	}
}
