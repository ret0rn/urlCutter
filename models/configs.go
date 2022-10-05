package models

type ServiceConfig struct {
	Port string `json:"port"`
}

type DbConfig struct {
	Db_host     string `json:"db_host"`
	Db_port     string `json:"db_port"`
	Db_user     string `json:"db_user"`
	Db_password string `json:"db_password"`
	Db_name     string `json:"db_name"`
	Sslmode     string `json:"sslmode"`
}
