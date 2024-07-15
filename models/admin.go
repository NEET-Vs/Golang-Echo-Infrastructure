package models

import "time"

type Admin struct {
	ID        int       `json:"id" gorm:"type:int(11);column:ID;primaryKey;autoIncrement:true"`
	Name      string    `json:"name" gorm:"type:varchar(255);column:Name"`
	Password  string    `json:"password" gorm:"type:varchar(255);column:Password"`
	CreatedAt time.Time `json:"-" gorm:"column:CreateAt"`
}