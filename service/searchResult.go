package service

import (
	"strings"

	"github.com/dasagho/htmx-test/models"
)

type SearchResultSerice struct {
}

func (s *SearchResultSerice) GetResults(query string) (*models.Result, error) {
	var lista []string
	valorSinEspacios := strings.TrimSpace(query)
	if valorSinEspacios == "" {
		lista = []string{}
	} else {
		lista = strings.Split(valorSinEspacios, " ")
	}

	data := models.Result{List: lista}
	return &data, nil
}
