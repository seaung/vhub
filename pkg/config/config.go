package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ServicesConfiguration struct {
	App       App       `mapstructure:"app" yaml:"app"`
	Databases Databases `mapstructure:databases" yaml:"databases"`
}

type App struct {
	Port string `mapstructure:"port" yaml:"port"`
}

type Databases struct {
	DbName  string `mapstructure:"db_name" yaml:"db_name"`
	DbUser  string `mapstructure:"db_user" yaml:"db_user"`
	DbPass  string `mapstructure:"db_pass" yaml:"db_pass"`
	DbPort  string `mapstructure:"db_port" yaml:"db_port"`
	Charset string `mapstructure:"charset" yaml:"charset"`
	Host    string `mapstructure:"host" yaml:"host"`
}

var SConfig = new(ServicesConfiguration)

func InitConfig() error {
	path, err := os.Getwd()
	if err != nil {
		panic("can't not find director")
	}
	viperConfig := viper.New()
	viperConfig.SetConfigType("yaml")
	viperConfig.AddConfigPath(path)
	viperConfig.SetConfigName("config.yml")
	viperConfig.SetConfigFile("config/config.yml")

	err = viperConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viperConfig.Unmarshal(SConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(SConfig.App.Port)
	return err
}
