package config

type LogConfig struct {
	Console ConsoleLogConfig `toml:"console"`
	File    FileLogConfig    `toml:"file"`
}

type ConsoleLogConfig struct {
	Enable bool   `toml:"enable"`
	Level  string `toml:"level"`
	Format string `toml:"format"`
}

type FileLogConfig struct {
	Enable bool   `toml:"enable"`
	Path   string `toml:"path"`
	Level  string `toml:"level"`
	Format string `toml:"format"`
}
