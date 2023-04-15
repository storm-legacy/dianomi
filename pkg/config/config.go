package config

import (
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
}

func Get(targetEnvVar string) interface{} {
	return viper.Get(targetEnvVar)
}

func GetString(targetEnvVar string) string {
	return viper.GetString(targetEnvVar)
}

func GetBool(targetEnvVar string) bool {
	return viper.GetBool(targetEnvVar)
}

func GetInt(targetEnvVar string) int {
	return viper.GetInt(targetEnvVar)
}

func GetFloat(targetEnvVar string) float64 {
	return viper.GetFloat64(targetEnvVar)
}

func Set(targetEnvVar string, value string) {
	viper.Set(targetEnvVar, value)
}
