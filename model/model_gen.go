// Code generated by bit. DO NOT EDIT.

package model

import (
	"database/sql/driver"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
	"time"
)

type Array []interface{}

func (x *Array) Scan(input interface{}) error {
	return jsoniter.Unmarshal(input.([]byte), x)
}

func (x Array) Value() (driver.Value, error) {
	return jsoniter.Marshal(x)
}

type Object map[string]interface{}

func (x *Object) Scan(input interface{}) error {
	return jsoniter.Unmarshal(input.([]byte), x)
}

func (x Object) Value() (driver.Value, error) {
	return jsoniter.Marshal(x)
}

func True() *bool {
	value := true
	return &value
}

func False() *bool {
	return new(bool)
}

type Role struct {
	ID          int64     `json:"id"`
	Status      *bool     `gorm:"default:true" json:"status"`
	CreateTime  time.Time `gorm:"autoCreateTime;default:current_timestamp" json:"create_time"`
	UpdateTime  time.Time `gorm:"autoUpdateTime;default:current_timestamp" json:"update_time"`
	Key         string    `gorm:"type:varchar;not null;unique" json:"key"`
	Name        string    `gorm:"type:varchar;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Routers     Array     `gorm:"type:jsonb;default:'[]'" json:"routers"`
	Permissions Array     `gorm:"type:jsonb;default:'[]'" json:"permissions"`
}

type Admin struct {
	ID          int64     `json:"id"`
	Status      *bool     `gorm:"default:true" json:"status"`
	CreateTime  time.Time `gorm:"autoCreateTime;default:current_timestamp" json:"create_time"`
	UpdateTime  time.Time `gorm:"autoUpdateTime;default:current_timestamp" json:"update_time"`
	Uuid        uuid.UUID `gorm:"type:uuid;not null;unique;default:uuid_generate_v4()" json:"-"`
	Username    string    `gorm:"type:varchar;not null;unique" json:"username"`
	Password    string    `gorm:"type:varchar;not null" json:"password"`
	Roles       Array     `gorm:"type:jsonb;not null;default:'[]'" json:"roles"`
	Name        string    `gorm:"type:varchar" json:"name"`
	Email       string    `gorm:"type:varchar" json:"email"`
	Phone       string    `gorm:"type:varchar" json:"phone"`
	Avatar      Array     `gorm:"type:jsonb;default:'[]'" json:"avatar"`
	Routers     Array     `gorm:"type:jsonb;default:'[]'" json:"routers"`
	Permissions Array     `gorm:"type:jsonb;default:'[]'" json:"permissions"`
}

func AutoMigrate(tx *gorm.DB, models ...string) {
	mapper := map[string]interface{}{
		"role": &Role{}, "admin": &Admin{},
	}

	for _, model := range models {
		if mapper[model] != nil {
			tx.AutoMigrate(mapper[model])
		}
	}
}
