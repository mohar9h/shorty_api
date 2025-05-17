package model

import "time"

type Link struct {
	ID        uint      `json:"id" gorm:"primary_key;auto_increment"`
	Code      string    `json:"code" gorm:"unique"`
	Url       string    `json:"url"`
	IPAddress string    `json:"ip_address" gorm:"index"`
	ExpireAt  time.Time `json:"expire_at" gorm:"index"`
	BaseModel
}

func (Link) TableName() string {
	return "links"
}
