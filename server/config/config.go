package config

import "github.com/spf13/viper"

type Config struct {
	MongoURI      string `mapstructure:"MONGO_URI"`
	ChannelSecret string `mapstructure:"CHANNEL_SECRET"`
	ChannelToken  string `mapstructure:"CHANNEL_ACCESS_TOKEN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("server")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
