package store

import (
	"time"

	"github.com/NuitdelinfoLesCools/back-end/pkg/store/model"
)

type missionStore interface {
	CreateMission(int64, string, string, time.Time, bool) error
	GetMissions(int64) (*[]model.Mission, error)
}

func (s *Store) CreateMission(userid int64, name string, description string, deadline time.Time, completed bool) error {
	mission := model.Mission{
		UserId:      userid,
		Name:        name,
		Description: description,
		Deadline:    deadline,
		Completed:   completed,
	}

	// validate user
	v, err := s.Valid(mission)
	if err != nil || !v {
		return err
	}

	// insert user
	if _, err := s.Insert(&mission); err != nil {
		return err
	}

	return nil
}

func (s *Store) GetMissions(userid int64) (*[]model.Mission, error) {
	missions := []model.Mission{}
	err := s.Where("user_id = ?", userid).Find(&missions)
	if err != nil {
		return nil, err
	}

	return &missions, nil
}
