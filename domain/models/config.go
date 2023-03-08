package models

type Config struct {
	Application		Application
	Server			Server
	DB				DB
}

type Application struct {
	Name			string `json:"name"`
	Version			string `json:"version"`
}

type Server struct {
	Port			string
}

type DB struct {
	User			string
	Password		string
	Name			string
	Driver			string
	Host			string
	Port			string
}
