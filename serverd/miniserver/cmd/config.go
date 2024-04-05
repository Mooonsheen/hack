package main

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (c *Config) Read() {
	c.Host = "localhost:"
	c.Port = "8081"
}
