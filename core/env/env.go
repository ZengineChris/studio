package env

import "os"

func IsDevelopment() bool {
	env, ok := os.LookupEnv("APP_ENVIRONMENT")
	if !ok {
		return true
	}

	if env == "development" {
		return true
	}
	return false
}

func IsProduction() bool {
	env, ok := os.LookupEnv("APP_ENVIRONMENT")
	if !ok {
		return false
	}

	if env == "production" {
		return true
	}
	return false
}
