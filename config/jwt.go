package config

type Jwt struct {
	SigningKey  string `mapstructure:"signing_key" json:"signing_key" yaml:"signing_key"`    // jwt签名
	ExpiresTime int64  `mapstructure:"expires_time" json:"expires_time" yaml:"expires_time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer_time" json:"buffer_time" yaml:"buffer_time"`    // 缓冲时间
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                   // 签发者
}
