package config 

import (
	"os"
	"fmt"
	"strconv"
)

type Config struct {
	PublicHost string
	Port string
	DBUser string
	DBPassword string
	DBAddress string
	DBName string
	JWTExpirationInSeconds int64
	JWTSecret string
}

var Envs = initConfig()

func initConfig() Config {
	
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port: getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "ENTER_MYSQL_PASSWORD_HERE"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "forum"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600 * 24 * 7),
		JWTSecret: getEnv("JWT_SECRET", "ENTER_JWT_SECRET_HERE"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}