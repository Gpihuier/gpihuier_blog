package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID          uint64         `gorm:"primarykey;comment:主键ID"`     // 主键ID
	CreateTime  time.Time      `gorm:"autoCreateTime;comment:更新时间"` // 创建时间
	UpdateTime  time.Time      `gorm:"autoCreateTime;comment:创建时间"` // 更新时间
	DeletedTime gorm.DeletedAt `gorm:"index;comment:删除时间"`          // 删除时间
}
