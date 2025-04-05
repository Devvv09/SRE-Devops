package config

import (
	"github.com/spf13/viper"
)

func LoadEnviromentVariable()(c Config,err error){

	viper.AddConfigPath("./config") //from env-variable
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig() //struct config
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&c)
	return
}
