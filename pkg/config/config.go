package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ServicesConfiguration struct {
	App       App       `mapstructure:"app" yaml:"app"`
	Databases Databases `mapstructure:"databases" yaml:"databases"`
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

// ValidateConfig validates the required configuration fields
func (c *ServicesConfiguration) ValidateConfig() error {
	if c.App.Port == "" {
		return fmt.Errorf("app port is required")
	}
	if c.Databases.DbName == "" {
		return fmt.Errorf("database name is required")
	}
	if c.Databases.DbUser == "" {
		return fmt.Errorf("database user is required")
	}
	return nil
}

// InitConfig initializes the configuration with support for:
// - Multiple config paths
// - Environment variables override
// - Configuration validation
// - Hot reload
func InitConfig() error {
	workDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %v", err)
	}

	viperConfig := viper.New()
	viperConfig.SetConfigType("yaml")
	
	// Support multiple config paths
	viperConfig.AddConfigPath(filepath.Join(workDir, "config"))
	viperConfig.AddConfigPath(workDir)
	viperConfig.SetConfigName("config")

	// Support environment variables override
	viperConfig.AutomaticEnv()
	viperConfig.SetEnvPrefix("VHUB")
	viperConfig.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viperConfig.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	if err := viperConfig.Unmarshal(SConfig); err != nil {
		return fmt.Errorf("failed to unmarshal config: %v", err)
	}

	// Validate configuration
	if err := SConfig.ValidateConfig(); err != nil {
		return fmt.Errorf("config validation failed: %v", err)
	}

	// Watch for config file changes
	viperConfig.WatchConfig()
	viperConfig.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
		if err := viperConfig.Unmarshal(SConfig); err != nil {
			fmt.Printf("Failed to reload config: %v\n", err)
		}
	})

	return nil
}
