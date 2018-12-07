package store

import (
	"github.com/NuitdelinfoLesCools/back-end/pkg/store/model"
	"github.com/medtune/go-utils/errors"
)

// Sync store models
func (s *Store) Sync() error {
	models := getModels()
	errs := []error{}
	for _, m := range models {
		if err := s.Sync2(m); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return errors.NewList(errs...)
	}

	return nil
}

// Return a map of empty model instances
func getModels() map[string]interface{} {
	m := make(map[string]interface{}, 3)
	m["User"] = &model.User{}
	m["Alert"] = &model.Alert{}
	m["Position"] = &model.Position{}
	m["Mission"] = &model.Mission{}
	m["Equipement"] = &model.Equipement{}
	m["Food"] = &model.Food{}
	return m
}
