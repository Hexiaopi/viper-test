package config

type HTTP struct {
	Host  string `mapstructure:"host" json:"host"`
	Port  int    `mapstructure:"port" json:"port"`
	Debug bool   `mapstructure:"debug" json:"debug"`
}
