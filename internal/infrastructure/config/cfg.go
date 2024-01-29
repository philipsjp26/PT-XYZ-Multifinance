package config

import (
	"go_playground/internal/core/common/utils"
	"log"
	"sync"
)

type Config struct {
	App      App       `yaml:"app"`
	Database *Database `yaml:"database"`
}

type App struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type Database struct {
	Driver          string `yaml:"driver"`
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	Name            string `yaml:"name"`
	ConnMaxLifetime int    `yaml:"max_lifetime_conn"`
	MaxOpenConn     int    `yaml:"max_open_conn"`
	MaxIdleConn     int64  `yaml:"max_idle_conn"`
}

func Configuration() *Config {
	var (
		once sync.Once
		cfg  *Config
	)
	path := "conf/app.yaml"

	once.Do(func() {

		err := utils.ReadFromYAML(path, &cfg)
		if err != nil {
			log.Fatal(err)
		}
	})

	return cfg
}
