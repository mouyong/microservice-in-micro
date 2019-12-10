package config

type MysqlConfig interface {
	GetEnabled() bool
	GetUrl() string
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
}

type defaultMysqlConfig struct {
	Enabled bool `json:"enabled"`
	Url string `json:"url"`
	MaxIdleConnection int `json:"maxIdleConnection"`
	MaxOpenConnection int `json:"maxOpenConnection"`
}

func (c defaultMysqlConfig) GetEnabled() bool {
	return c.Enabled
}

func (c defaultMysqlConfig) GetUrl() string {
	return c.Url
}

func (c defaultMysqlConfig) GetMaxIdleConnection() int {
	return c.MaxIdleConnection
}

func (c defaultMysqlConfig) GetMaxOpenConnection() int {
	return c.MaxOpenConnection
}


