package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	HTTP struct {
		IP   string `yaml:"ip" env:"HTTP-IP" env-default:"localhost"`
		Port string `yaml:"port" env:"HTTP-PORT" env-default:"8080"`
	} `yaml:"http"`
}

const (
	RootPath         string = "/"
	HtmlPath         string = "./html/page.html"
	FindChangingPath string = "/find-changing"
)

var (
	configPath string
	once       sync.Once
	instance   *Config
)

const (
	FlagConfigPathName = "config"
)

func GetConfig() *Config {
	once.Do(func() {
		flag.StringVar(
			&configPath,
			FlagConfigPathName,
			"configs/config.yaml",
			"this is app config file",
		)
		flag.Parse()

		log.Print("config init")

		instance = &Config{}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			helpText := "changing money"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})

	return instance
}
