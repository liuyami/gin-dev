package config

type Jwt struct {
	Secret               string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Ttl                  int64  `mapstructure:"ttl" json:"ttl" yaml:"ttl"`                                                          // token 有效期（秒）
	BlacklistGracePeriod int64  `mapstructure:"blacklist_grace_period" json:"blacklist_grace_period" yaml:"blacklist_grace_period"` // token 有效期（秒）
}
