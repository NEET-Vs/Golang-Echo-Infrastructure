package models

import "time"

type Staff struct {
	ID        int       `json:"id" gorm:"type:int(11);column:ID;primaryKey;autoIncrement:true"`
	Name      string    `json:"name" gorm:"type:varchar(255);column:Name"`
	AdminID   int       `json:"admin_id" gorm:"type:int(11);column:AdminID"`
	CreatedAt time.Time `json:"-" gorm:"column:CreateAt"`
}
