package config

type AppConfig struct {
	Name    string       `mapstructure:"name"`
	Mode    string       `mapstructure:"mode"`
	Version string       `mapstructure:"version"`
	Port    int          `mapstructure:"port"`
	Log     *LogConfig   `mapstructure:"log"`
	Mysql   *MysqlConfig `mapstructure:"mysql"`
	Redis   *RedisConfig `mapstructure:"redis"`
	Es      *EsConfig    `mapstructure:"es"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DbName       string `mapstructure:"dbname"`
	MaxConn      int    `mapstructure:"max_conn"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poolsize"`
}
type EsConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
