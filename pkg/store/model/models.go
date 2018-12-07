package model

import "time"

// User .
type User struct {
	Id       int64  `yaml:"omitempty"` //`xorm:"pk autoincr"`
	Username string `xorm:"varchar(16) not null unique" valid:"alphanum,required" yaml:"username"`
	Email    string `xorm:"varchar(64) not null unique" valid:"email,required" yaml:"email"`
	Password string `xorm:"varchar(128) not null" valid:"required" yaml:"password"`
}

// Position .
type Position struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Date      time.Time `xorm:"created" json:"date"`
}

// Alert .
type Alert struct {
	Id       int64     `json:"id"`
	UserId   int64     `json:"user_id"`
	Message  string    `xorm:"varchar(1000) not null" json:"message"`
	Severity int8      `json:"severity"`
	Date     time.Time `xorm:"created" date:"date"`
}

type Food struct {
	Id       int64  `json:"id"`
	UserId   int64  `json:"user_id"`
	Name     string `xorm:"varchar(16) not null" json:"name"`
	InStock  bool   `json:"in_stock"`
	Quantity int64  `json:"quantity"`
	Unit     string `json:"unit"`
}

type Equipement struct {
	Id      int64  `json:"id"`
	UserId  int64  `json:"user_id"`
	Name    string `xorm:"varchar(16) not null" json="name"`
	InStock bool   `json:"in_stock"`
}

type Mission struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description" xorm:"varchar(1000)`
	Deadline    time.Time `json:"deadline"`
	Completed   bool      `json:"completed"`
}
