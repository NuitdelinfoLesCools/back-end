package store

import (
	"github.com/NuitdelinfoLesCools/back-end/pkg/store/model"
)

type alertStore interface {
	CreateAlert(int64, string, int8) error
	GetAlerts(int64) (*[]model.Alert, error)
}

func (s *Store) CreateAlert(userid int64, message string, sev int8) error {
	alert := model.Alert{
		UserId:   userid,
		Message:  message,
		Severity: sev,
	}

	// validate user
	v, err := s.Valid(alert)
	if err != nil || !v {
		return err
	}

	// insert user
	if _, err := s.Insert(&alert); err != nil {
		return err
	}

	return nil
}

func (s *Store) GetAlerts(userid int64) (*[]model.Alert, error) {
	alerts := []model.Alert{}
	err := s.Where("user_id = ?", userid).Find(&alerts)
	if err != nil {
		return nil, err
	}

	return &alerts, nil
}
