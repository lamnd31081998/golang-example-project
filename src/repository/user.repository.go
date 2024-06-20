package repository

import (
	configModule "golang-example-project/config"
	structModule "golang-example-project/struct"
)

func MigrateUserTable() {
	configModule.MasterDB.AutoMigrate(&structModule.User{})
}

func FindUserByUsername(username string) *structModule.User {
	var user structModule.User
	if err := configModule.MasterDB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil
	}
	return &user
}

func FindUserById(id uint) *structModule.User {
	var user structModule.User
	if err := configModule.MasterDB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil
	}
	return &user
}

func CreateUser(insertData structModule.User) *structModule.User {
	if err := configModule.MasterDB.Omit("ID", "DeletedAt").Create(&insertData).Error; err != nil {
		return nil
	}
	return &insertData
}

func UpdateUserById(updateData structModule.User) *structModule.User {
	if err := configModule.MasterDB.Model(&updateData).Omit("ID", "CreatedAt", "UpdatedAt").Updates(updateData).Error; err != nil {
		return nil
	}
	return &updateData
}
