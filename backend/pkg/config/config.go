package config

import (
	"fmt"
	"math"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	log.SetLevel(log.DebugLevel)

	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.WithField("msg", err.Error()).Warn("Could not read config file")
	}

	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("APP_PG_HOST"),
		viper.GetInt("APP_PG_PORT"),
		viper.GetString("APP_PG_USER"),
		viper.GetString("APP_PG_PASSWORD"),
		viper.GetString("APP_PG_DB"),
	)

	// Custom made variables
	Set("PG_CONNECTION_STRING", connectionString)
}

func GetString(targetEnvVar string, defaultValue ...string) string {
	if viper.IsSet(targetEnvVar) {
		return viper.GetString(targetEnvVar)
	} else {
		for _, value := range defaultValue {
			return value
		}
		log.WithField("field", targetEnvVar).Panic("Environment variable is not set")
		return ""
	}
}

func GetBool(targetEnvVar string, defaultValue ...bool) bool {
	if viper.IsSet(targetEnvVar) {
		return viper.GetBool(targetEnvVar)
	} else {
		for _, value := range defaultValue {
			return value
		}
		log.WithField("field", targetEnvVar).Panic("Environment variable is not set")
		return false
	}
}

func GetInt(targetEnvVar string, defaultValue ...int) int {
	if viper.IsSet(targetEnvVar) {
		return viper.GetInt(targetEnvVar)
	} else {
		for _, value := range defaultValue {
			return value
		}
		log.WithField("field", targetEnvVar).Panic("Environment variable is not set")
		return math.MinInt
	}
}

func GetFloat(targetEnvVar string, defaultValue ...float64) float64 {
	if viper.IsSet(targetEnvVar) {
		return viper.GetFloat64(targetEnvVar)
	} else {
		for _, value := range defaultValue {
			return value
		}
		log.WithField("field", targetEnvVar).Panic("Environment variable is not set")
		return math.SmallestNonzeroFloat64
	}
}

func Set(targetEnvVar string, value string) {
	viper.Set(targetEnvVar, value)
}
