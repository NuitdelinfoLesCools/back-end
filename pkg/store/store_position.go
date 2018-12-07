package store

import (
	"time"

	"github.com/NuitdelinfoLesCools/back-end/pkg/store/model"
)

type positionStore interface {
	CreatePosition(int64, float64, float64, time.Time) error
	GetPositions(int64) (*[]model.Position, error)
}

func (s *Store) CreatePosition(userid int64, lat, long float64, date time.Time) error {
	position := model.Position{
		UserId:    userid,
		Latitude:  lat,
		Longitude: long,
		Date:      date,
	}

	// validate user
	v, err := s.Valid(position)
	if err != nil || !v {
		return err
	}

	// insert user
	if _, err := s.Insert(&position); err != nil {
		return err
	}

	return nil
}

func (s *Store) GetPositions(userid int64) (*[]model.Position, error) {
	positions := []model.Position{}
	err := s.Where("user_id = ?", userid).Find(&positions)
	if err != nil {
		return nil, err
	}

	return &positions, nil
}
