package seeders

import (
	"api/configs"
	"api/repositories"
	"api/types/structs/requests"
	"github.com/bxcodec/faker/v4"
)

func UserSeeder() {
	for i := 0; i < 100; i++ {
		user := requests.UserRequest{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: "password123",    
			Role:     "tenant",           
			Phone:    faker.Phonenumber(),
			Address:  faker.Word() + " Street",
		}

		_, _ = repositories.CreateUser(configs.DB, user)
	}
}
