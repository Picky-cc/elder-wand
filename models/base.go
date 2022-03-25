package models

import (
	"time"

	"elder-wand/utils/dbUtils"
)

type BaseModel struct {
	ID      dbUtils.SFID `gorm:"primary_key;" faker:"-"`
	Created time.Time    `gorm:"type:datetime(6);index;not null;" faker:"-"`
	Updated time.Time    `gorm:"type:datetime(6);index;not null;" faker:"-"`
	Deleted *time.Time   `gorm:"type:datetime(6)" faker:"-"`
}

// BeforeCreate gorm interface, call func before model create
func (m *BaseModel) BeforeCreate() (err error) {
	if !m.ID.IsValid() {
		m.ID = dbUtils.GenerateID()
	}
	nowTime := time.Now()
	m.Created = nowTime
	m.Updated = nowTime

	return
}

func (m *BaseModel) InitBaseModel() {
	m.ID = dbUtils.GenerateID()
	nowTime := time.Now()
	m.Created = nowTime
	m.Updated = nowTime
}

func (m BaseModel) GetID() dbUtils.SFID {
	return m.ID
}

// IsValid 大致判断 model instance 是否有效的方法
func (m BaseModel) IsValid() bool {
	return m.ID.IsValid()
}
