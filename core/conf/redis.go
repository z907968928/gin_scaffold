package conf

type RedisConf struct {
	List map[string]*RedisBaseConf `mapstructure:"list" yaml:"list" toml:"list" json:"list"`
}

type RedisBaseConf struct {
	ProxyList    []string `mapstructure:"proxy_list" yaml:"proxy_list" toml:"proxy_list" json:"proxy_list"`
	MaxActive    int      `mapstructure:"max_active" yaml:"max_active" toml:"max_active" json:"max_active"`
	MaxIdle      int      `mapstructure:"max_idle" yaml:"max_idle" toml:"max_idle" json:"max_idle"`
	DownGrade    bool     `mapstructure:"down_grade" yaml:"down_grade" toml:"down_grade" json:"down_grade"`
	Password     string   `mapstructure:"password" yaml:"password" toml:"password" json:"password"`
	Db           int      `mapstructure:"db" yaml:"db" toml:"db" json:"db"`
	ConnTimeout  int      `mapstructure:"conn_timeout" yaml:"conn_timeout" toml:"conn_timeout" json:"conn_timeout"`
	ReadTimeout  int      `mapstructure:"read_timeout" yaml:"read_timeout" toml:"read_timeout" json:"read_timeout"`
	WriteTimeout int      `mapstructure:"write_timeout" yaml:"write_timeout" toml:"write_timeout" json:"write_timeout"`
}
