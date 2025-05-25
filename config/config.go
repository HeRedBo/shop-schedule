package config

import "time"

type Config struct {
	App           App           `mapstructure:"app" yaml:"app"`
	Database      Database      `mapstructure:"database" yaml:"database"`
	MongoDB       MongoDB       `mapstructure:"mongodb"`
	Redis         Redis         `mapstructure:"redis" yaml:"redis"`
	Elasticsearch Elasticsearch `mapstructure:"elasticsearch" yaml:"elasticsearch"`
	Zap           Zap           `mapstructure:"zap" yaml:"zap"`
	Api           Api           `mapstructure:"api" yaml:"api"`
}

type App struct {
	RunMode         string `mapstructure:"run_mode"  yaml:"run_mode"`
	RuntimeRootPath string `mapstructure:"runtime_root_path" yaml:"runtime_root_path"`
	LogSavePath     string `mapstructure:"log_save_path" yaml:"log_save_path"`
	LogSaveName     string `mapstructure:"log_save_name" yaml:"log_save_name"`
	LogFileExt      string `mapstructure:"log_file_ext" yaml:"log_file_ext"`
	TimeFormat      string `mapstructure:"time_format" yaml:"time_format"`
	Task            string `mapstructure:"task" yaml:"task"`
}

type Database struct {
	Type        string `mapstructure:"type" yaml:"type"`
	User        string `mapstructure:"user" yaml:"user"`
	Password    string `mapstructure:"password" yaml:"password"`
	Host        string `mapstructure:"host" yaml:"host"`
	Name        string `mapstructure:"name" yaml:"name"`
	TablePrefix string `mapstructure:"table-prefix" yaml:"table-prefix"`
}

type Redis struct {
	Host        string        `mapstructure:"host" yaml:"host"`
	Password    string        `mapstructure:"password" yaml:"password"`
	IdleTimeout time.Duration `mapstructure:"idle-timeout" yaml:"idle-timeout"`
}

type Elasticsearch struct {
	Hosts    []string `mapstructure:"hosts" yaml:"hosts"`
	Username string   `mapstructure:"username" yaml:"username"`
	Password string   `mapstructure:"password" yaml:"password"`
}

type Zap struct {
	LogFilePath     string `mapstructure:"log-filepath" yaml:"log-filepath"`
	LogInfoFileName string `mapstructure:"log-info-filename" yaml:"log-info-filename"`
	LogWarnFileName string `mapstructure:"log-warn-filename" yaml:"log-warn-filename"`
	LogFileExt      string `mapstructure:"log-fiile-ext" yaml:"log-fiile-ext"`
}

type MongoDB struct {
	DBname   string   `mapstructure:"dbname"`
	User     string   `mapstructure:"user"`
	Password string   `mapstructure:"password"`
	Host     []string `mapstructure:"host"`
}

type Api struct {
	OrderAK  string `mapstructure:"order-ak" yaml:"order-ak"`
	OrderSK  string `mapstructure:"order-sk" yaml:"order-sk"`
	ShopHost string `mapstructure:"shop-host" yaml:"shop-host"`
}
