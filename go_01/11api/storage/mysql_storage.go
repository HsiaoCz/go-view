package storage

import (
	"go-cloud/go_01/11api/models"
	"go-cloud/go_01/11api/utils"
)

type Mysql_Storage struct{}

func NewMysqlStorage() *Mysql_Storage {
	return &Mysql_Storage{}
}

func (m *Mysql_Storage) GetUserByID(id int) *models.User {
	return &models.User{
		ID:       id,
		Username: "bob",
		Password: "hsaicz",
		Email:    "123399@qq.com",
		Number:   utils.GetNumber(),
	}
}
