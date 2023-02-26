package config

type System struct {
	Env          string `mapstructure:"env" json:"env" yaml:"env"`
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
}
