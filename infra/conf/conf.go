package conf

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Port        string
	MongoUri    string
	MongoDbName string
}

var Env *Configuration

func LoadConfig() {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println(err)
	}
	Env = &Configuration{
		Port:        os.Getenv("PORT"),
		MongoUri:    os.Getenv("MONGO_URI"),
		MongoDbName: os.Getenv("MONGO_DB_NAME"),
	}
}
