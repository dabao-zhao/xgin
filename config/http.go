package config

type HttpConfig struct {
	Host                 string   `toml:"host"`
	Port                 int64    `toml:"port"`
	AccessLogFile        string   `toml:"accessLogFile"`
	CorsEnable           bool     `toml:"corsEnable"`
	CorsAllowCredentials bool     `toml:"corsAllowCredentials"`
	CorsAllowAllOrigins  bool     `toml:"corsAllowAllOrigins"`
	CorsAllowOrigins     []string `toml:"corsAllowOrigins"`
	CorsAllowMethods     []string `toml:"corsAllowMethods"`
	CorsAllowHeaders     []string `toml:"corsAllowHeaders"`
	CorsExposeHeaders    []string `toml:"corsExposeHeaders"`
	CorsMaxAge           int64    `toml:"corsMaxAge"`
}
