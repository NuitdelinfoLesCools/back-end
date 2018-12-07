package store

import (
	"github.com/NuitdelinfoLesCools/back-end/pkg/store/db"
	"github.com/asaskevich/govalidator"

	"github.com/go-xorm/xorm"
)

var _ userStore = &Store{}
var _ foodStore = &Store{}
var _ alertStore = &Store{}
var _ missionStore = &Store{}
var _ equipementStore = &Store{}

var (
	// Agent main object used by other packages
	Agent *Store
)

// Store type is the abstraction behind data interactions
// Database io & validation
type Store struct {
	*xorm.Engine
	Valid func(interface{}) (bool, error)
}

// New return a database engine
func New(config db.ConnStr) (*Store, error) {
	engine, err := db.NewPGEngine(config)
	if err != nil {
		return nil, err
	}

	s := &Store{
		Engine: engine,
		Valid:  govalidator.ValidateStruct,
	}

	return s, nil
}
