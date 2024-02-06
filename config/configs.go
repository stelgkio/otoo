package configs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env    string
	Auth   AuthConfig
	DB     DBConfig
	Server ServerConfig
	Job    JobConfig
}

func ReadConfigs(path string) (Config, error) {
	var cfg Config
	f, err := os.Open(path)
	if err != nil {
		return cfg, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

type AuthConfig struct {
	Secret string
	Exp    int
}
type DBConfig struct {
	Driver string
	URI    string
}
type ServerConfig struct {
	Port int
	Host string
}

func (s *ServerConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type JobConfig struct {
	SampleJobPeriod int `yaml:"sampleJobPeriod"`
}
