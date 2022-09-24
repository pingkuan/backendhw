package config

import (
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

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

func Genbot() *linebot.Client {
	config, err := LoadConfig("./server")
	if err != nil {
		log.Fatal(err)
	}

	bot, err := linebot.New(config.ChannelSecret, config.ChannelToken)
	if err != nil {
		log.Fatal(err)
	}

	return bot
}

var Bot *linebot.Client = Genbot()
