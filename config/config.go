package config

type Enter struct {
	Server Server `mapstructure:"server" json:"server" yaml:"server"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Jwt    Jwt    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
