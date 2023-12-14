package db

import (
	"testing"

	"github.com/dasagho/htmx-test/models"
)

func TestSaveContact(t *testing.T) {
	// Setup
	db := NewMyContactData() // Suponiendo que esta función crea una nueva instancia de la implementación de ContactData
	contact := models.Contact{Id: 1, Name: "Test Contact", Telefono: 123456789}

	// Ejecución
	err := db.Save(contact)

	// Verificación
	if err != nil {
		t.Errorf("Se esperaba que no hubiera error, pero se obtuvo: %v", err)
	}
}
