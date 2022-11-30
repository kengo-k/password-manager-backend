package env

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RepositoryURL  string
	RepositoryUser string
	RepositoryPass string
	PasswordFile   string
}

func GetConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic("failed to load .env")
	}
	url := os.Getenv("REPOSITORY_URL")
	user := os.Getenv("REPOSITORY_USER")
	pass := os.Getenv("REPOSITORY_PASS")
	file := os.Getenv("PASSWORD_FILE")
	if url == "" {
		panic("missing env: REPOSITORY_URL")
	}
	if user == "" {
		panic("missing env: REPOSITORY_USER")
	}
	if pass == "" {
		panic("missing env: REPOSITORY_PASS")
	}
	if file == "" {
		panic("missing env: PASSWORD_FILE")
	}
	return &Config{
		RepositoryURL:  url,
		RepositoryUser: user,
		RepositoryPass: pass,
		PasswordFile:   file,
	}
}
