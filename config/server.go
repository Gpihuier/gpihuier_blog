package config

type Server struct {
	Addr int `mapstructure:"addr" json:"addr" yaml:"addr"`
}
