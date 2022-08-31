package util

/*
* @Author: XuanLT1
 */

import (
	"os"

	"github.com/spf13/viper"
)

// // Config stores all configuration of the application.
// // The values are read by viper from a config file or environment variable.
// type Config struct {
// 	Host     string `mapstructure:"postgres.host"`
// 	Port     string `mapstructure:"postgres.port"`
// 	Username string `mapstructure:"postgres.username"`
// 	Password string `mapstructure:"postgres.password"`
// 	Database string `mapstructure:"postgres.database"`
// 	SSLMode  string `mapstructure:"postgres.ssl_mode"`
// 	Timezone string `mapstructure:"postgres.time_zone"`
// }

// // LoadConfig reads configuration from file or environment variables.
// func LoadConfig(path string) (config Config, err error) {
// 	viper.AddConfigPath(path)
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yml")

// 	// Tell viper to automatically override values that it has read from config file with the values
// 	// of the corresponding environment variable if they exist.
// 	viper.AutomaticEnv()

// 	// Start reading values
// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return
// 	}

// 	err = viper.Unmarshal(&config)
// 	return

// }

func LoadConfig(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	return viper.MergeInConfig()
}
