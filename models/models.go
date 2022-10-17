package models

type HTTPServerConfig struct {
	Port string `yaml:"port"`
	Path string `yaml:"path"`
}

type Config struct {
	HTTPServerConfig `yaml:"http_server_config"`
	DBConfig         string `yaml:"db_config"`
}
