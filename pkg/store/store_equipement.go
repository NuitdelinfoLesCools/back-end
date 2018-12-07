package store

import (
	"github.com/NuitdelinfoLesCools/back-end/pkg/store/model"
)

type equipementStore interface {
	CreateEquipement(int64, string, bool) error
	EquipementStock(int64) (*[]model.Equipement, error)
}

// CreateEquipement .
func (s *Store) CreateEquipement(userid int64, name string, stock bool) error {
	equipement := model.Equipement{
		UserId:  userid,
		Name:    name,
		InStock: stock,
	}

	// validate user
	v, err := s.Valid(equipement)
	if err != nil || !v {
		return err
	}

	// insert user
	if _, err := s.Insert(&equipement); err != nil {
		return err
	}

	return nil
}

// EquipementStock .
func (s *Store) EquipementStock(userid int64) (*[]model.Equipement, error) {
	equipements := []model.Equipement{}
	err := s.Where("user_id = ?", userid).Find(&equipements)
	if err != nil {
		return nil, err
	}

	return &equipements, nil
}
