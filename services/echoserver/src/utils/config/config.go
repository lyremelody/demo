package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var configPath = "/etc/service.yaml"

type ServiceInfo struct {
	Port     int    `yaml:"port"`
	Language string `yaml:"language"`
}

func GetServiceInfo() (*ServiceInfo, error) {
	serviceinfo := &ServiceInfo{}
	if f, err := os.Open(configPath); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(serviceinfo)
	}
	return serviceinfo, nil
}
