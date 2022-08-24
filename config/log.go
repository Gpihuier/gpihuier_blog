package config

type Zap struct {
	Level        string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Format       string `mapstructure:"format" json:"format" yaml:"format"`                         // 输出
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	RuntimeDir   string `mapstructure:"runtime_dir" json:"runtime_dir" yaml:"runtime_dir"`          // 缓存文件夹
	LogDir       string `mapstructure:"log_dir" json:"log_dir" yaml:"log_dir"`                      // 日志文件夹
	LogInConsole bool   `mapstructure:"log_in_console" json:"log_in_console" yaml:"log_in_console"` // 输出控制台
	EncodeLevel  string `mapstructure:"encode_level" json:"encode_level" yaml:"encode_level"`       // 编码器
}
