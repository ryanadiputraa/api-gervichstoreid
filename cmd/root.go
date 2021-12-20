package cmd

import (
	"log"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/spf13/viper"
)

func init() {
	loadConfig()
	configureHystrix()
}

func loadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".gervichstoreid")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("load config: %s", err.Error())
	}
}

func configureHystrix() {
	hystrix.ConfigureCommand("SimpleQuery", hystrix.CommandConfig{
		Timeout:               15000,
		MaxConcurrentRequests: 125,
		ErrorPercentThreshold: 25,
	})

	hystrix.ConfigureCommand("HistoricalExtractQuery", hystrix.CommandConfig{
		Timeout:               150000,
		MaxConcurrentRequests: 500,
		ErrorPercentThreshold: 25,
	})

	hystrix.ConfigureCommand("HighLevelQuery", hystrix.CommandConfig{
		Timeout:               5000000,
		MaxConcurrentRequests: 10000,
		ErrorPercentThreshold: 50,
	})
}

func Execute() {
	serveHTTP()
}
