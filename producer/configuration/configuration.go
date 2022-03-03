package configuration

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

func Load() error {
	viper.SetConfigName("configuration")
	viper.SetConfigType("yml")
	viper.AddConfigPath("configuration")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	return err
}
