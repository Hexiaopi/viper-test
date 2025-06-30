package config

type App struct {
	HTTP HTTP `mapstructure:"http" json:"http"`
	DB   DB   `mapstructure:"db" json:"db"`
}
