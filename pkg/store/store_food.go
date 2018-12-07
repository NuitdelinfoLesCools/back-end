package store

import (
	"github.com/NuitdelinfoLesCools/back-end/pkg/store/model"
)

type foodStore interface {
	CreateFood(int64, string, bool, int64, string) error
	FoodStock(int64) (*[]model.Food, error)
}

// CreateFood .
func (s *Store) CreateFood(userid int64, name string, stock bool, quantity int64, unit string) error {
	food := model.Food{
		UserId:   userid,
		Name:     name,
		InStock:  stock,
		Quantity: quantity,
		Unit:     unit,
	}

	// validate user
	v, err := s.Valid(food)
	if err != nil || !v {
		return err
	}

	// insert user
	if _, err := s.Insert(&food); err != nil {
		return err
	}

	return nil
}

func (s *Store) FoodStock(userid int64) (*[]model.Food, error) {
	foods := []model.Food{}
	err := s.Where("user_id = ?", userid).Find(&foods)
	if err != nil {
		return nil, err
	}

	return &foods, nil
}
