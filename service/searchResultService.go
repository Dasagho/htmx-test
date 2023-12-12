package service

import (
	"strings"

	"github.com/dasagho/htmx-test/db"
	"github.com/dasagho/htmx-test/models"
)

type SearchResultSerice struct {
	SearchRepo *db.ResultRepository
}

func (s *SearchResultSerice) GetResults(query string) (*models.Result, error) {
	// result, err := s.GetResults(query)
	// if err != nil {
	// return &models.Result{List: []string{}}, err
	// }

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
