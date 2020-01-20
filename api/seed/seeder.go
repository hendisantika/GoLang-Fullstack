package seed

import (
	"github.com/victorsteven/fullstack/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}
