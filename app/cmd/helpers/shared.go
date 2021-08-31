package helpers

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func MustHaveEnv(key string) string {
	env := viper.GetString(key)
	if env == "" {
		viper.SetConfigFile(".env")
		viper.SetConfigType("env")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err, "can't read .env file")
		}
		env = viper.GetString(key)
	}
	if env == "" {
		log.Fatal(fmt.Sprintf("%s is not well set", key))
	}
	return env
}