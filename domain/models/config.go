package models

type Config struct {
	Application		Application
	Server			Server
	DB				DB
	Session         Session
	Email           Email
}

type Application struct {
	Name			string `json:"name"`
	Version			string `json:"version"`
}

type Server struct {
	Port			string
}

type Session struct {
	Key             string
}

type DB struct {
	User			string
	Password		string
	Name			string
	Driver			string
	Host			string
	Port			string
}

type Email struct {
	From            string
	Password        string
	Host            string
	Port            string
}
