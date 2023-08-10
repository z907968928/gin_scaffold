package conf

type MysqlConf struct {
	List map[string]*MysqlBaseConf `mapstructure:"list" yaml:"list" toml:"list" json:"list"`
}

type MysqlBaseConf struct {
	DriverName      string `mapstructure:"driver_name" yaml:"driver_name" toml:"driver_name" json:"driver_name"`
	Host            string `mapstructure:"host" yaml:"host" toml:"host" json:"host"`
	Port            string `mapstructure:"port" yaml:"port" toml:"port" json:"port"`
	User            string `mapstructure:"user" yaml:"user" toml:"user" json:"user"`
	Password        string `mapstructure:"password" yaml:"password" toml:"password" json:"password"`
	Database        string `mapstructure:"database" yaml:"database" toml:"database" json:"database"`
	Charset         string `mapstructure:"charset" yaml:"charset" toml:"charset" json:"charset"`
	AddTo           string `mapstructure:"add_to" yaml:"add_to" toml:"add_to" json:"add_to"`
	DataSourceName  string `mapstructure:"data_source_name" yaml:"data_source_name" toml:"data_source_name" json:"data_source_name"`
	MaxOpenConn     int    `mapstructure:"max_open_conn" yaml:"max_open_conn" toml:"max_open_conn" json:"max_open_conn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn" yaml:"max_idle_conn" toml:"max_idle_conn" json:"max_idle_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time" yaml:"max_conn_life_time" toml:"max_conn_life_time" json:"max_conn_life_time"`
}
