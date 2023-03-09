package api

type IUserService interface {
	GetUsername(int) string
}

type UserService struct{}

func (s UserService) GetUsername(userID int) string {
	if userID == 101 {
		return "bob"
	}
	return "alex"
}
