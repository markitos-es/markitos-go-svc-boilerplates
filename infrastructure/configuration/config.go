package configuration

import (
	"github.com/spf13/viper"
)

type BoilerplateConfiguration struct {
	DatabaseDsn       string `mapstructure:"DATABASE_DSN"`
	ServerAddress     string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
}

func LoadConfiguration(
	configFilesPath string) (config BoilerplateConfiguration, err error) {

	viper.AddConfigPath(configFilesPath)

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)

	return config, err
}
