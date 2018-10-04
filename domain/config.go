package domain

type Configuration struct {
	Engine   string
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

type Error struct {
	err error
	msg string
}