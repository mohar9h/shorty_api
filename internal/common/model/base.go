package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	IsDeleted bool          `json:"is_deleted" gorm:"default:false"`
	IsActive  bool          `json:"is_active" gorm:"default:true"`
	CreatedBy sql.NullInt64 `json:"created_by" gorm:"null"`
	UpdatedBy sql.NullInt64 `json:"updated_by" gorm:"null"`
	DeletedBy sql.NullInt64 `json:"deleted_by" gorm:"null"`
	CreatedAt time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt time.Time     `json:"deleted_at" gorm:"null"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	userId := tx.Statement.Context.Value("UserId")

	if userId != nil {
		base.CreatedBy = sql.NullInt64{
			Int64: userId.(int64),
			Valid: true,
		}
	}
	base.CreatedAt = time.Now().UTC()

	return
}

func (base *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	userId := tx.Statement.Context.Value("UserId")
	if userId != nil {
		base.UpdatedBy = sql.NullInt64{
			Int64: userId.(int64),
			Valid: true,
		}
	}

	base.UpdatedAt = time.Now().UTC()

	return
}

func (base *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	userId := tx.Statement.Context.Value("UserId")
	base.DeletedBy = sql.NullInt64{
		Int64: userId.(int64),
		Valid: true,
	}

	base.DeletedAt = time.Now().UTC()

	return
}
