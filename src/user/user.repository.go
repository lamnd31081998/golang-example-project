package user

import (
	configModule "golang-example-project/config"
)

func MigrateTable() {
	configModule.MasterDB.AutoMigrate(&User{})
}

func FindByUsername(username string) *User {
	var user User
	if err := configModule.MasterDB.Where("username = ?", "lamnd").First(&user).Error; err != nil {
		return nil
	}
	return &user
}

func Create(insertData User) *User {
	if err := configModule.MasterDB.Omit("ID", "DeletedAt").Create(&insertData).Error; err != nil {
		return nil
	}
	return &insertData
}
