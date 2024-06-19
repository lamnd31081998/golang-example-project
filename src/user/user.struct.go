package user

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"column:id;not null;autoIncrement;primaryKey"`
	Username  string    `json:"username" gorm:"column:username;not null;uniqueIndex"`
	Name      string    `json:"name" gorm:"column:name;not null"`
	Password  string    `json:"password" gorm:"column:password;not null"`
	AvatarUrl string    `json:"avatar_url" gorm:"column:avatar_url;default:null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at;default:null"`
}
