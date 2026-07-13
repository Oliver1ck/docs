package config

type Logger struct {
	LogLevel   string `default:"debug"   envconfig:"LOGGER_LEVEL"`
	LogFormat  string `default:"console" envconfig:"LOGGER_FORMAT"`
	LogNoColor bool   `default:"false"   envconfig:"LOGGER_NOCOLOR"`
}
