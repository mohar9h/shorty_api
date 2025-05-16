package model

type Link struct {
	ID   uint   `json:"id" gorm:"primary_key;auto_increment"`
	Code string `json:"code" gorm:"unique"`
	Url  string `json:"url"`
	BaseModel
}

func (Link) TableName() string {
	return "links"
}
