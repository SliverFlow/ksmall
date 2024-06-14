package conf

import "fmt"

// Jaeger is a client for Jeager.
type Jaeger struct {
	Service string `mapstructure:"service" yaml:"service"`
	Host    string `mapstructure:"host" yaml:"host"`
	Port    int64  `mapstructure:"port" yaml:"port"`
	LogSpan bool   `mapstructure:"logSpan" yaml:"logSpan"`
}

func (j *Jaeger) Agent() string {
	return fmt.Sprintf("%s:%d", j.Host, j.Port)
}
