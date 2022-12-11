package env

import (
	"errors"
	"fmt"
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

func NewConfig(configPath string) (IConfig, error) {
	envMap, err := godotenv.Read(configPath)

	if err != nil {
		return nil, fmt.Errorf("failed to load %v", configPath)
	}
	url := envMap["REPOSITORY_URL"]
	user := envMap["REPOSITORY_USER"]
	pass := envMap["REPOSITORY_PASS"]
	file := envMap["PASSWORD_FILE"]
	if url == "" {
		return nil, errors.New("missing env: REPOSITORY_URL")
	}
	if user == "" {
		return nil, errors.New("missing env: REPOSITORY_USER")
	}
	if pass == "" {
		return nil, errors.New("missing env: REPOSITORY_PASS")
	}
	if file == "" {
		return nil, errors.New("missing env: PASSWORD_FILE")
	}
	return &Config{
		RepositoryURL:  url,
		RepositoryUser: user,
		RepositoryPass: pass,
		PasswordFile:   file,
	}, nil
}
