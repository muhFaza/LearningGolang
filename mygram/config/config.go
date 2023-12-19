package config

import "fmt"

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		DBName:   "mygram",
	}
}

func (config *DBConfig) GetDBURL() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.DBName,
		config.Password,
	)
}
