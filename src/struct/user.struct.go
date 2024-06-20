package customstruct

import "time"

type User struct {
	ID         uint      `json:"id" gorm:"column:id;not null;autoIncrement;primaryKey"`
	Username   string    `json:"username" gorm:"column:username;not null;uniqueIndex"`
	Name       string    `json:"name" gorm:"column:name;not null"`
	Password   string    `json:"-" gorm:"column:password;not null"`
	AvatarUrl  string    `json:"avatar_url" gorm:"column:avatar_url;default:null"`
	LastActive time.Time `json:"last_active" gorm:"column:last_active;autoCreateTime;autoUpdateTime"`
	Status     int       `json:"status" gorm:"column:status;default:1"`
	CreatedAt  time.Time `json:"-" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `json:"-" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt  time.Time `json:"-" gorm:"column:deleted_at;default:null"`
}
