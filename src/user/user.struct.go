package user

import "time"

type User struct {
	Id        uint      `json:"id" gorm:"column:id,primaryKey,autoIncrement,<-:create"`
	Username  string    `json:"username" gorm:"column:username,index,not null,unique,<-:create"`
	Name      string    `json:"name" gorm:"column:name,not null"`
	Phone     string    `json:"phone" gorm:"column:phone"`
	AvatarUrl string    `json:"avatar_url" gorm:"column:avatar_url"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at,autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at,autoCreateTime,autoUpdateTime"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at,index"`
}
