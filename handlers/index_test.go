package handlers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	logging "github.com/dasagho/htmx-test/log"
	"github.com/dasagho/htmx-test/models"
	"github.com/dasagho/htmx-test/views"
)

func TestIndexHandler(t *testing.T) {
	// Cambio de ruta para correcto funcionamiento del paquete view
	os.Chdir("../")

	// Inicializar los logs
	logging.CreateLogs()

	// Crear peticion HTTP
	req, err := http.NewRequest("GET", "/", nil) // Tercer parametro para el cuerpo de la peticion http
	if err != nil {
		log.Fatal("Error al crear la peticion http /" + err.Error())
	}

	// Cargar en memoria los templates
	views.InitializeTemplate()

	// Crear ResponseRecorder
	rr := httptest.NewRecorder()

	// Llamar al handler
	http.HandlerFunc(IndexHandler).ServeHTTP(rr, req)

	// Comparar el contenido renderizado.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Crear datos mock
	data := models.Body{
		MouseDiv: "Mamala",
		Search:   "Search",
		List:     models.Result{List: []string{}},
	}

	// Renderizar el template en una cadena de texto con los datos esperados.
	var expected bytes.Buffer
	if err := views.GetTemplates().ExecuteTemplate(&expected, "index", data); err != nil {
		t.Fatal(err)
	}

	// Comparar el cuerpo de la respuesta con el contenido esperado.
	if rr.Body.String() != expected.String() {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected.String())
	}
}
