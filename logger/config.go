package logger

type Config struct {
	JSON   bool   `yaml:"json"`
	Level  string `yaml:"level"`
	Sample bool   `yaml:"sample,omitempty"`
	// when sampling, the first N logs will be logged
	SampleInitial int `yaml:"sampleInitial,omitempty"`
	// when sampling, every Mth log will be logged
	SampleInterval int `yaml:"sampleInitial,omitempty"`
}
