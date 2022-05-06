package logx

type Config struct {
	Env         string `yaml:"env,omitempty"`
	Level       string `yaml:"level,omitempty"`
	ServiceName string `yaml:"service_name,omitempty"`
}
