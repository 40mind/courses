package models

type Config struct {
	Application		Application
	Server			Server
	DB				DB
}

type Application struct {
	Name			string
	Version			string
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
