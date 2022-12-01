//Package config gather configuration data from file
package config

import (
	"log"
	"os"
)

//Config holds configuration information
type Config struct {
	EmployeeQueue string `env:"EMPLOYEE_QUEUE"`
}

//New creates a config
func New() (*Config, error) {
	cfg := &Config{}
	err := cfg.readConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

//readConfig reads config.json
func (c *Config) readConfig() error {
	log.Println("Carregando configurações...")

	c.EmployeeQueue = os.Getenv("EMPLOYEE_QUEUE")

	return nil
}
