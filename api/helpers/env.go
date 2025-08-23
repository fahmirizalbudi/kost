package helpers

import "github.com/joho/godotenv"

func LoadENV() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}
