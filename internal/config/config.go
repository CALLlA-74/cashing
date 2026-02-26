package config

type Config struct {
	HTTP struct {
		IP   string `yaml:"ip" env:"HTTP-IP" env-default:"localhost"`
		Port string `yaml:"port" env:"HTTP-PORT" env-default:"8080"`
	} `yaml:"http"`
}
