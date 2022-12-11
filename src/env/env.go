package env

import (
	"os"

	"github.com/joho/godotenv"
)

type IConfig interface {
	GetRepositoryURL() string
	GetRepositoryUser() string
	GetRepositoryPass() string
	GetPasswordFile() string
}

type Config struct {
	RepositoryURL  string
	RepositoryUser string
	RepositoryPass string
	PasswordFile   string
}

func (c *Config) GetRepositoryURL() string {
	return c.RepositoryURL
}

func (c *Config) GetRepositoryUser() string {
	return c.RepositoryUser
}

func (c *Config) GetRepositoryPass() string {
	return c.RepositoryPass
}

func (c *Config) GetPasswordFile() string {
	return c.PasswordFile
}

func NewConfig(configPath string) IConfig {
	err := godotenv.Load(configPath)
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
