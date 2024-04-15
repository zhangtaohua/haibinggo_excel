// Package user 存放用户 Model 相关逻辑
package user

import (
	"haibinggo/app/models"
	"haibinggo/pkg/database"
	"haibinggo/pkg/hash"
)

// User 用户模型
// User 用户模型
type User struct {
	models.BaseModel

	Name     string `json:"name"`
	NickName string `json:"nick_name"`

	City         string `json:"city"`
	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar"`

	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	if len(userModel.Password) < 32 {
		return _password == userModel.Password
	} else {
		return hash.BcryptCheck(_password, userModel.Password)
	}
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}

func (userModel *User) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&userModel)
	return result.RowsAffected
}
