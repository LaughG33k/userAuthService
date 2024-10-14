package config

type AppConfig struct {
	Addr         string   `json:"addr"`
	ReadTimeout  string   `json:"readTimeout"`
	WriteTimeout string   `json:"writeTimeout"`
	UserDb       DBConfig `json:"userDB"`
	SessionDb    DBConfig `json:"sessionDB"`
}
