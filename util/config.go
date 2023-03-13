package util

import "github.com/spf13/viper"

// Config stores all configurations of the application
type Config struct {
	DBDriver      string `mapStructure:"DB_DRIVER"`
	DBSource      string `mapStructure:"DB_SOURCE"`
	ServerAddress string `mapStructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
