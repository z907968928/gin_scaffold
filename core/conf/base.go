package conf

type BaseConf struct {
	Base    Base    `mapstructure:"base" yaml:"base" toml:"base" json:"base"`
	HTTP    HTTP    `mapstructure:"http" yaml:"http" toml:"http" json:"http"`
	Log     Log     `mapstructure:"log" yaml:"log" toml:"log" json:"log"`
	Swagger Swagger `mapstructure:"swagger" yaml:"swagger" toml:"swagger" json:"swagger"`
}

type Base struct {
	DebugMode    string `mapstructure:"debug_mode" yaml:"debug_mode" toml:"debug_mode" json:"debug_mode"`
	TimeLocation string `mapstructure:"time_location" yaml:"time_location" toml:"time_location" json:"time_location"`
}

type HTTP struct {
	Addr           string   `mapstructure:"addr" yaml:"addr" toml:"addr" json:"addr"`
	PidPath        string   `mapstructure:"pid_path" yaml:"pid_path" toml:"pid_path" json:"pid_path"`
	ReadTimeout    int      `mapstructure:"read_timeout" yaml:"read_timeout" toml:"read_timeout" json:"read_timeout"`
	WriteTimeout   int      `mapstructure:"write_timeout" yaml:"write_timeout" toml:"write_timeout" json:"write_timeout"`
	MaxHeaderBytes int      `mapstructure:"max_header_bytes" yaml:"max_header_bytes" toml:"max_header_bytes" json:"max_header_bytes"`
	AllowIP        []string `mapstructure:"allow_ip" yaml:"allow_ip" toml:"allow_ip" json:"allow_ip"`
}

// log config
type Log struct {
	Level         string        `mapstructure:"level" yaml:"level" toml:"level" json:"level"`
	Layout        string        `mapstructure:"layout" yaml:"layout" toml:"layout" json:"layout"`
	TakeUp        bool          `mapstructure:"take_up" yaml:"take_up" toml:"take_up" json:"take_up"`
	FileWriter    FileWriter    `mapstructure:"file_writer" yaml:"file_writer" toml:"file_writer" json:"file_writer"`
	ConsoleWriter ConsoleWriter `mapstructure:"console_writer" yaml:"console_writer" toml:"console_writer" json:"console_writer"`
}

type FileWriter struct {
	Open            bool   `mapstructure:"open" yaml:"open" toml:"on" json:"open"`
	LogPath         string `mapstructure:"log_path" yaml:"log_path" toml:"log_path" json:"log_path"`
	RotateLogPath   string `mapstructure:"rotate_log_path" yaml:"rotate_log_path" toml:"rotate_log_path"  json:"rotate_log_path"`
	WfLogPath       string `mapstructure:"wf_log_path" yaml:"wf_log_path" toml:"wf_log_path" json:"wf_log_path"`
	RotateWfLogPath string `mapstructure:"rotate_wf_log_path" yaml:"rotate_wf_log_path" toml:"rotate_wf_log_path" json:"rotate_wf_log_path"`
}

type ConsoleWriter struct {
	Open  bool `mapstructure:"open" yaml:"open" toml:"open" json:"open"`
	Color bool `mapstructure:"color" yaml:"color" toml:"color" json:"color"`
}

// swagger
type Swagger struct {
	Title    string `mapstructure:"title" yaml:"title" toml:"title" json:"title"`
	Desc     string `mapstructure:"desc" yaml:"desc" toml:"desc" json:"desc"`
	Host     string `mapstructure:"host" yaml:"host" toml:"host" json:"host"`
	BasePath string `mapstructure:"base_path" yaml:"base_path" toml:"base_path" json:"base_path"`
}
