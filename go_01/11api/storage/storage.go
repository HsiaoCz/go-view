package storage

import "go-cloud/go_01/11api/models"

type Storage interface {
	GetUserByID(int) *models.User
}
