package config

type Config struct {
	LogLevel      string `yaml:"LogLevel"`
	Cron          string `yaml:"Cron"`
	MaxConcurrent int    `yaml:"MaxConcurrent"`

	// Client
	Api api `yaml:"Api"`

	// Controller
	Mode      string `yaml:"Mode"`
	Address   string `yaml:"Address"`
	UdpxyHost string `yaml:"UdpxyHost"`
	Exclude   string `yaml:"Exclude"`
}

type api struct {
	Provider string            `yaml:"Provider"`
	ApiHost  string            `yaml:"ApiHost"`
	EPGPath  string            `yaml:"EPGPath"`
	Auth     map[string]string `yaml:"Auth"`
}
