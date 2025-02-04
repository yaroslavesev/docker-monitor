package config

import (
	"os"
)

//Config gets environmental variables

func GetDBHost() string {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}
	return host
}

func GetDBUser() string {
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = "postgres"
	}
	return user
}

func GetDBPassword() string {
	pass := os.Getenv("POSTGRES_PASSWORD")
	if pass == "" {
		pass = "postgres"
	}
	return pass
}

func GetDBName() string {
	dbname := os.Getenv("POSTGRES_DB")
	if dbname == "" {
		dbname = "backend_db"
	}
	return dbname
}
