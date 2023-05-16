package g

import (
	"log"
	"os"
	"sync"

	"github.com/toolkits/file"
	"gopkg.in/yaml.v2"
)

// EnableConfig configs which can be used
type EnableConfig struct {
	Collect   bool `yaml:"collect"`
	LogRotate bool `yaml:"log_rotate"`
	Witch     bool `yaml:"witch"`
}

// RabbitConfig ...
type RabbitConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// FalconConfig ...
type FalconConfig struct {
	API string `yaml:"api"`
}

// HTTPConfig ...
type HTTPConfig struct {
	ConnTimeout int `yaml:"conn_timeout"`
	RespTimeout int `yaml:"response_time