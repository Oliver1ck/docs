package config

import "time"

type JWT struct {
	Secret          string        `required:"true"  envconfig:"JWT_SECRET"`
	AccessTokenTTL  time.Duration `required:"false" envconfig:"JWT_ACCESS_TTL"  default:"15m"`
	RefreshTokenTTL time.Duration `required:"false" envconfig:"JWT_REFRESH_TTL" default:"168h"` // 7 дней
}
