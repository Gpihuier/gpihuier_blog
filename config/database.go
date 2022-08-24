package config

import "fmt"

type Mysql struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Hostname string `mapstructure:"hostname" json:"hostname" yaml:"hostname"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	HostPort int    `mapstructure:"host_port" json:"host_port" yaml:"host_port"`
	Prefix   string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
}

func (m *Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Hostname,
		m.HostPort,
		m.Database,
	)
}
