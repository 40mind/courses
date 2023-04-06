package models

type Config struct {
	Application		    Application
	Server			    Server
	DB				    DB
	Session             Session
	Email               Email
	YookassaProvider    Provider
	YookassaAuth        Auth
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

type Provider struct {
	Host            string
	Endpoint        map[string]Endpoint
}

type Endpoint struct {
	Path            string
	Method          string
}

type Auth struct {
	Login           string
	Password        string
}
