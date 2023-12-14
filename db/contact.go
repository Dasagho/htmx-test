package db

import (
	"database/sql"

	"github.com/dasagho/htmx-test/models"
)

type ContactData interface {
	Save(contact models.Contact) error
	FindByID(contactID int) (models.Contact, error)
	Update(contactID int, contact models.Contact) error
	Delete(contactID int) error
}

type myContactData struct {
	db *sql.DB
}

func NewMyContactData() ContactData {
	return &myContactData{
		db: db,
	}
}

// Delete implements ContactData.
func (cd *myContactData) Delete(contactID int) error {
	// res, err := cd.db.Exec("DELETE FROM database_test WHERE id = %s", contactID)
	// if err != nil {
	// 	return err
	// }

	// rows, err := res.RowsAffected()
	// if err != nil {
	// 	return err
	// }

	// logging.Debug(fmt.Sprintf("Eliminado contacto: %d, numero de filas afectadas: %d", contactID, rows))
	return nil
}

// FindByID implements ContactData.
func (cd *myContactData) FindByID(contactID int) (models.Contact, error) {
	panic("unimplemented")
}

// Save implements ContactData.
func (cd *myContactData) Save(contact models.Contact) error {
	panic("unimplemented")
}

// Update implements ContactData.
func (cd *myContactData) Update(contactID int, contact models.Contact) error {
	panic("unimplemented")
}
