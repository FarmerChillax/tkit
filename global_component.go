package tkit

import "context"

type databaseConfig struct {
	Dsn               string `json:"dsn,omitempty"`
	Driver            string `json:"driver,omitempty"`
	Loc               string `json:"loc,omitempty"`
	ParseTime         bool   `json:"parse_time,omitempty"`
	Timeout           int64  `json:"timeout,omitempty"`
	MaxOpen           int    `json:"max_open,omitempty"`
	MaxIdle           int    `json:"max_idle,omitempty"`
	ConnMaxLifeSecond int    `json:"conn_max_life_second,omitempty"`
	Host              string `json:"host,omitempty"`
	UserName          string `json:"user_name,omitempty"`
	Password          string `json:"password,omitempty"`
	DBName            string `json:"db_name,omitempty"`
	Charset           string `json:"charset,omitempty"`
	MultiStatements   bool   `json:"multi_statements,omitempty"`
}

func registerDatabase(ctx context.Context) error {
	return nil
}

func registerOtel(ctx context.Context) error {
	return nil
}

func registerRedis(ctx context.Context) error {
	return nil
}
