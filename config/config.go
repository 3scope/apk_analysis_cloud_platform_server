package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
	Type string
	Path string
}

func (c *Config) initConfig() error {
	viper.SetConfigName(c.Name)
	viper.SetConfigType(c.Type)
	viper.AddConfigPath(c.Path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// Check whether the file has changed.
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config File is Changed, Name: %s.\n", c.Name)
	})
}

func init() {
	c := Config{
		Name: "config",
		Type: "yaml",
		Path: "./config",
	}
	if err := c.initConfig(); err != nil {
		log.Panicln(err)
	}
	c.watchConfig()
}
